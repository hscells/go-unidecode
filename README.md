go-unidecode
==============

[![Go Report Card](https://goreportcard.com/badge/github.com/dan-locke/go-unidecode)](https://goreportcard.com/report/github.com/dan-locke/go-unidecode)
[![GoDoc](https://godoc.org/github.com/dan-locke/go-unidecode?status.svg)](https://godoc.org/github.com/mozillazg/go-unidecode)

A modified version of mozillag/go-unidecode for byte slice output. ASCII transliterations of Unicode text.


Installation
------------

```
go get -u github.com/dan-locke/go-unidecode
```

Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/dan-locke/go-unidecode


Usage
------

```go
package main

import (
	"fmt"
	"github.com/dan-locke/go-unidecode"
)

func main() {
	s := "abc"
	fmt.Printf("%s\n", unidecode.Unidecode(s))
	// Output: abc

	s = "北京"
	fmt.Printf("%s\n", unidecode.Unidecode(s))
	// Output: Bei Jing

	s = "kožušček"
	fmt.Printf("%s\n", unidecode.Unidecode(s))
	// Output: kozuscek
}
```
