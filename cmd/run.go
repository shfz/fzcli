/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/shfz/fzcli/run"
	"github.com/shfz/fzcli/ui"
	"github.com/shfz/fzcli/util"
	"github.com/spf13/cobra"
)

var (
	target   string
	number   int
	parallel int64
	output   string
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !util.FileExists(target) {
			log.Fatal("[-] No target file found : ", target)
		}
		ui.Init()
		if err := run.Run(target, parallel, number, output); err != nil {
			log.Fatal("[-] Failed to run : ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.Flags().StringVarP(&target, "target", "t", "", "target script file (required)")
	if err := runCmd.MarkFlagRequired("target"); err != nil {
		panic(err)
	}

	runCmd.Flags().Int64VarP(&parallel, "parallel", "p", 1, "parallel")
	runCmd.Flags().IntVarP(&number, "number", "n", 1, "number")
	runCmd.Flags().StringVarP(&output, "output", "o", "", "output directory")
}
