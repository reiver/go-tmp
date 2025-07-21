# go-tmp

Package **tmp** implements an **expiring** **optional-type** with a **time-out**, for the Go programming language.

In other programming languages, an **optional-type** might be known as: a **option-type**, or a **maybe-type**.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-tmp

[![GoDoc](https://godoc.org/github.com/reiver/go-tmp?status.svg)](https://godoc.org/github.com/reiver/go-tmp)

## Example

Here is a simple example:

```golang
import "github.com/reiver/go-tmp"

// ...

var until time.Time = time.Now().Add(8 * time.Hour)

var temporal tmp.Temporal[string] = tmp.Temporary(value, until)

// ...

// If the temporal has expired then this will return the optional value 'nothing'.
// Else, it will return the optional value 'something' (containing the same internal value that was inside of the temporal).
optional := temporal.Optional()

val, found := optional.Get()
if found {
	// ...
} else {
	// ...

}

```

## Import

To import package **tmp** use `import` code like the follownig:
```
import "github.com/reiver/go-tmp"
```

## Installation

To install package **tmp** do the following:
```
GOPROXY=direct go get github.com/reiver/go-tmp
```

## Author

Package **tmp** was written by [Charles Iliya Krempeaux](http://reiver.link)
