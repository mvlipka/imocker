# imocker
[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/mvlipk/imocker)
[![Go Report](https://goreportcard.com/badge/github.com/mvlipka/imocker)](https://goreportcard.com/report/github.com/mvlipka/imocker)
[![Go Coverage](https://github.com/mvlipka/imocker/wiki/coverage.svg)](https://raw.githack.com/wiki/mvlipka/imocker/coverage.html)

## Overview
imocker is a tool designed to generate mock structs that implement any interface in order to effectively write unit tests.

imocker generates actual Go code that can be used in unit tests, no using strings to look up methods, no chaining function calls in order to add expectations, and no ability to diverge from the expected interface implementations.

## Installation
### Go Install
`go install github.com/mvlipka/imocker@latest`

## Usage
```
imocker generate ./...
imocker generate ./testdata
```

# Examples
A small example can be seen in the `testdata` folder.  
[testdata/mock_thinger.go](testdata/mock_thinger.go) was generated by using `imocker generate ./...`

# Current Limitations
* `MyMethod(multiple, vars bool) (error)`
  * Methods with multiple parameters defined are unsupported
* `MyMethod(bool, bool) (error)`
  * Methods with no parameter names defined are unsupported