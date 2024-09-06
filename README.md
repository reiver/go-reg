# go-reg

Package **reg** implements a thread-safe registry, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-reg

[![GoDoc](https://godoc.org/github.com/reiver/go-reg?status.svg)](https://godoc.org/github.com/reiver/go-reg)

## Examples

Here is an example:

```golang
import "github.com/reiver/go-reg"

// ...

var registrar reg.Registry[MyCustomType]

// ...

var value MyCustomType // = ...

reg.Set(name, value)

// ...

value := reg.Get(name))
```

## Import

To import package **reg** use `import` code like the follownig:
```
import "github.com/reiver/go-reg"
```

## Installation

To install package **reg** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-reg
```

## Author

Package **reg** was written by [Charles Iliya Krempeaux](http://reiver.link)
