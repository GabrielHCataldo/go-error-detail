Go Errors Detail
=================
<!--suppress ALL -->
<img align="right" src="gopher-debug.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.1.7-vividgreen.svg)](https://github.com/GabrielHCataldo/go-errors/releases/tag/v1.1.7)
[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielHCataldo/go-errors)](https://goreportcard.com/report/github.com/GabrielHCataldo/go-errors)
[![Coverage Status](https://coveralls.io/repos/GabrielHCataldo/go-errors/badge.svg?branch=main&service=github)](https://coveralls.io/github/GabrielHCataldo/go-errors?branch=main)
[![Open Source Helpers](https://www.codetriage.com/gabrielhcataldo/go-errors/badges/users.svg)](https://www.codetriage.com/gabrielhcataldo/go-errors)
[![GoDoc](https://godoc.org/github/GabrielHCataldo/go-errors?status.svg)](https://pkg.go.dev/github.com/GabrielHCataldo/go-errors/errors)
![License](https://img.shields.io/dub/l/vibe-d.svg)

[//]: # ([![build workflow]&#40;https://github.com/GabrielHCataldo/go-errors/actions/workflows/go.yml/badge.svg&#41;]&#40;https://github.com/GabrielHCataldo/go-errors/actions&#41;)

[//]: # ([![Source graph]&#40;https://sourcegraph.com/github.com/go-errors/errors/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/go-errors/errors?badge&#41;)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/GabrielHCataldo/go-errors/errors&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/GabrielHCataldo/go-errors&#41;)

The go-errors project came to clarify the return of errors, which are very common in Golang, thus facilitating the
debugging your projects.

Installation
------------

Use go get.

	go get github.com/GabrielHCataldo/go-errors

Then import the go-errors package into your own code.

```go
import "github.com/GabrielHCataldo/go-errors/errors"
```

Usability and documentation
------------
**IMPORTANT**: Always check the documentation in the structures and functions fields.
For more details on the examples, visit [All examples link](https://github/GabrielHCataldo/go-errors/blob/main/_example/main).

### Simple example

```go
package main

import (
    "github.com/GabrielHCataldo/go-errors/errors"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    err := simple()
    logger.Info("simple err:", err)
    logger.Info("simple err msg:", errors.Details(err).GetMessage())
    logger.Info("simple err file:", errors.Details(err).GetFile())
    logger.Info("simple err line:", errors.Details(err).GetLine())
    logger.Info("simple err func:", errors.Details(err).GetFuncName())
    errors.Details(err).PrintCause()
    errors.Details(err).PrintStackTrace()
}

func simple() error {
    return errors.New("error by message with any value", 2, true)
}
```

Output:

    [INFO 2024/01/26 10:16:38] _example/main.go:12: simple err: [CAUSE]: (_example/main.go:25) simple: error by message with any value 2 true [STACK]:
    goroutine 1 [running]:
    runtime/debug.Stack()
        /Users/gabrielcataldo/go/go1.21.3/src/runtime/debug/stack.go:24 +0x64
    github.com/GabrielHCataldo/go-errors/errors.New({0x1400039fd18?, 0x0?, 0x1400039fc38?})
        /Users/gabrielcataldo/Innovfor/GabrielHCataldo/go-errors/errors/errors.go:31 +0xe0
    main.simple(...)
        /Users/gabrielcataldo/Innovfor/GabrielHCataldo/go-errors/_example/main.go:25
    main.main()
        /Users/gabrielcataldo/Innovfor/GabrielHCataldo/go-errors/_example/main.go:11 +0x88
    [INFO 2024/01/26 10:16:38] _example/main.go:12: simple err msg: error by message with any value 2 true
    [INFO 2024/01/26 10:16:38] _example/main.go:13: simple err file: _example/main.go
    [INFO 2024/01/26 10:16:38] _example/main.go:14: simple err line: 25
    [INFO 2024/01/26 10:16:38] _example/main.go:15: simple err func: simple
    [ERROR 2024/01/26 10:16:38] _example/main.go:16: (_example/main.go:25) simple: error by message with any value 2 true
    [ERROR 2024/01/26 10:16:38] _example/main.go:17: goroutine 1 [running]:

How to contribute
------
Make a pull request, or if you find a bug, open it
an Issues.

License
-------
Distributed under MIT license, see the license file within the code for more details.