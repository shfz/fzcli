# fzcli

## Features

- Easy to customize fuzzing for web applications
- Scripting fuzzing scenario in JavaScript / TypeScript
- Third-party packages can be used in scenario script

## Install

### Binary

Download binary from [Releases](https://github.com/shfz/fzcli/releases) page.

#### Linux (amd64)

```
$ curl -Lo fzcli.tar.gz https://github.com/shfz/fzcli/releases/download/v0.0.1/fzcli_0.0.1_linux_amd64.tar.gz
$ tar -zxvf fzcli.tar.gz
$ sudo mv fzcli /usr/local/bin/
$ sudo chmod +x /usr/local/bin/fzcli
```

## Usage

Please refer to [shfz/fzlib-node](https://github.com/shfz/fzlib-node) for creating scenarios.

```
fzcli run -t scenario.js -o /tmp/fzlog -p 10 -n 100
```

### Options

- -t : Scenario script file
- -o : Log output location
- -p : Number of parallel executions of fuzzing
- -n : Number of total executions of fuzzing
