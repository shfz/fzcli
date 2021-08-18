# fzcli

Scenario-based fuzzing test execution tool

![fzcli](https://user-images.githubusercontent.com/12624257/129859210-f29a0d7f-7e1e-42a6-a855-0e83d05b4cde.gif)

## Features

- **Easy to customize** fuzzing test for web applications
- Scripting fuzzing scenario **in JavaScript / TypeScript**
- **Third-party packages** can be used in scenario script

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

This tool runs a scenario that calls http requests for the web application, with automatically embeds the fuzz in the request parameter (`username`, `password`, etc).

Please refer to [shfz/fzlib-node](https://github.com/shfz/fzlib-node) for how to script scenarios.

```
fzcli run -t scenario.js -o /tmp/fzlog -p 10 -n 100
```

### Options

- -t : Scenario script file
- -o : Log output location
- -p : Number of parallel executions of fuzzing
- -n : Number of total executions of fuzzing
