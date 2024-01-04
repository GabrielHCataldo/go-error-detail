Go Error Detail
=================
<!--suppress ALL -->
<img align="right" src="gopher-debug.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.0.0-vividgreen.svg)](https://github.com/GabrielHCataldo/go-error-detail/releases/tag/v1.0.5)
[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielHCataldo/go-error-detail)](https://goreportcard.com/report/github.com/GabrielHCataldo/go-error-detail)
[![Coverage Status](https://coveralls.io/repos/GabrielHCataldo/go-error-detail/badge.svg?branch=main&service=github)](https://coveralls.io/github/GabrielHCataldo/go-error-detail?branch=main)
[![Open Source Helpers](https://www.codetriage.com/gabrielhcataldo/go-error-detail/badges/users.svg)](https://www.codetriage.com/gabrielhcataldo/go-error-detail)
[![GoDoc](https://godoc.org/github/GabrielHCataldo/go-error-detail?status.svg)](https://pkg.go.dev/github.com/GabrielHCataldo/go-error-detail/errors)
![License](https://img.shields.io/dub/l/vibe-d.svg)

[//]: # ([![build workflow]&#40;https://github.com/GabrielHCataldo/go-error-detail/actions/workflows/go.yml/badge.svg&#41;]&#40;https://github.com/GabrielHCataldo/go-error-detail/actions&#41;)

[//]: # ([![Source graph]&#40;https://sourcegraph.com/github.com/go-error-detail/errors/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/go-error-detail/errors?badge&#41;)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/GabrielHCataldo/go-error-detail/errors&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/GabrielHCataldo/go-error-detail&#41;)

The go-error-detail project came to make the return of errors, very common in Golang, clearer, thus facilitating the 
debugging of your applications.

Installation
------------

Use go get.

	go get github.com/GabrielHCataldo/go-error-detail

Then import the go-error-detail package into your own code.

```go
import "github.com/GabrielHCataldo/go-error-detail/errors"
```

Usability and documentation
------------
**IMPORTANT**: Always check the documentation in the structures and functions fields.
For more details on the examples, visit [All examples link](https://github/GabrielHCataldo/go-error-detail/blob/main/_example/main)

### Simple example

```go
package main

import (
    "github.com/GabrielHCataldo/go-error-detail/errors"
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

    [ERROR 2024/01/04 09:11:18] main.go:11: simple result: {"file":"/Users/gabrielcataldo/Innovfor/go-error-detail/_example/main.go","line":15,"message":"error by message with any value 2 true"}

How to contribute
------
Make a pull request, or if you find a bug, open it
an Issues.

License
-------
Distributed under MIT license, see the license file within the code for more details.