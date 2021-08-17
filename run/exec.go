package run

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/shfz/fzcli/model"
	"github.com/shfz/fzcli/ui"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type Output struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
	Seed    uint64 `json:"seed"`
	Http    Http   `json:"http"`
}

type Http struct {
	Status uint64 `json:"status"`
	Url    string `json:"url"`
	Method string `json:"method"`
}

var r = model.Result{}

func Run(target string, parallel int64, number int, logDir string) (err error) {
	r.Total = uint64(number)
	path := filepath.Join(logDir, time.Now().Format("20060102150405")+".log")
	start := time.Now()
	if err := ExecParallel(target, parallel, number, path); err != nil {
		return err
	}
	end := time.Now()
	ui.Close()
	fmt.Printf("[+] Time : %f seconds\n", (end.Sub(start)).Seconds())
	fmt.Println("[+] Log : " + path)
	fmt.Println("[+] Success : " + strconv.FormatUint(r.Success, 10) + " (" + strconv.Itoa(int(100*r.Success/r.Total)) + "%)")
	fmt.Println("[+] Failure : " + strconv.FormatUint(r.Failure, 10) + " (" + strconv.Itoa(int(100*r.Failure/r.Total)) + "%)")
	return nil
}

func ExecParallel(target string, parallel int64, number int, path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// https://zenn.dev/aikizoku/articles/golang-goroutine
	ctx := context.Background()
	eg := errgroup.Group{}
	mutex := sync.Mutex{}
	sem := semaphore.NewWeighted(parallel)
	for i := 0; i < number; i++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			return err
		}
		eg.Go(func() error {
			out, err := Exec(target)
			if err != nil {
				return err
			}

			var res Output
			if err := json.Unmarshal([]byte(out), &res); err != nil {
				return err
			}

			mutex.Lock()
			// update result
			if res.Code == 0 {
				r.Success += 1
			}
			if res.Code == 1 {
				r.Failure += 1
				r.Message = append(r.Message, time.Now().Format("15:04:05")+" "+res.Http.Method+" "+strconv.FormatUint(res.Http.Status, 10)+" "+res.Http.Url+" "+res.Message)
			}
			ui.Update(r)
			// write log
			_, err = file.WriteString(out)
			if err != nil {
				return err
			}
			mutex.Unlock()

			sem.Release(1)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func Exec(target string) (output string, err error) {
	rand.Seed(time.Now().UnixNano())
	seed := strconv.FormatUint(rand.Uint64(), 10)
	out, err := ExecCommand(seed, target)
	if err != nil {
		return "", err
	}
	return out, nil
}

func ExecCommand(seed string, target string) (output string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	out, err := exec.CommandContext(ctx, "node", target, seed).Output()

	if ctx.Err() == context.DeadlineExceeded {
		return "", errors.New("ExecCommand Timeout")
	}

	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
