Go Errors Detail
=================
<!--suppress ALL -->
<img align="right" src="gopher-debug.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.0.5-vividgreen.svg)](https://github.com/GabrielHCataldo/go-errors/releases/tag/v1.0.5)
[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielHCataldo/go-errors)](https://goreportcard.com/report/github.com/GabrielHCataldo/go-errors)
[![Coverage Status](https://coveralls.io/repos/GabrielHCataldo/go-errors/badge.svg?branch=main&service=github)](https://coveralls.io/github/GabrielHCataldo/go-errors?branch=main)
[![Open Source Helpers](https://www.codetriage.com/gabrielhcataldo/go-errors/badges/users.svg)](https://www.codetriage.com/gabrielhcataldo/go-errors)
[![GoDoc](https://godoc.org/github/GabrielHCataldo/go-errors?status.svg)](https://pkg.go.dev/github.com/GabrielHCataldo/go-errors/errors)
![License](https://img.shields.io/dub/l/vibe-d.svg)

[//]: # ([![build workflow]&#40;https://github.com/GabrielHCataldo/go-errors/actions/workflows/go.yml/badge.svg&#41;]&#40;https://github.com/GabrielHCataldo/go-errors/actions&#41;)

[//]: # ([![Source graph]&#40;https://sourcegraph.com/github.com/go-errors/errors/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/go-errors/errors?badge&#41;)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/GabrielHCataldo/go-errors/errors&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/GabrielHCataldo/go-errors&#41;)

The go-errors project came to make the return of errors, very common in Golang, clearer, thus facilitating the 
debugging of your applications.

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
For more details on the examples, visit [All examples link](https://github/GabrielHCataldo/go-errors/blob/main/_example/main)

### Simple example

```go
package main

import (
    "github.com/GabrielHCataldo/go-errors/errors"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    err := simple()
    logger.Error("simple result:", err)
}

func simple() error {
    return errors.New("error by message with any value", 2, true)
}
```

Output:

    [ERROR 2024/01/04 09:11:18] main.go:11: simple result: {"file":"/Users/gabrielcataldo/Innovfor/go-errors/_example/main.go","line":15,"message":"error by message with any value 2 true"}

How to contribute
------
Make a pull request, or if you find a bug, open it
an Issues.

License
-------
Distributed under MIT license, see the license file within the code for more details.