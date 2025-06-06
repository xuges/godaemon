# godaemon

## Overview

Simple API to start daemon process in Go program.

Normaly in Go, child process will exit when parent process exitd, sometimes you maybe want start a indepent, detached daemon process, not a child process, so you need this library.


Supported OS:
- Windows (developing)
- Linux (developing)

## Usage

Just like `exec.Command`:

Windows:

```go
package main

import "github.com/xuges/godaemon"

func main() {
    name := "notepad.exe"
    args := []string{ "test.txt"}
    attr := &godaemon.ProcAttr {
        Dir: "C:/",
        Env: []string{ "TEST=daemon" },
    }

    err := godaemon.StartDaemonProcess(name, args, attr)
    if err != nil {
        panic(err)
    }
}
```

Linux:

```go
package main

import "github.com/xuges/godaemon"

func main() {
    name := "gedit"
    args := []string{ "test.txt"}
    attr := &godaemon.ProcAttr {
        Dir: "/tmp",
        Env: []string{ "TEST=daemon" },
    }

    err := godaemon.StartDaemonProcess(name, args, attr)
    if err != nil {
        panic(err)
    }
}
```