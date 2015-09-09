## go-binary

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-binary)
[![Build Status](https://travis-ci.org/ferhatelmas/go-binary.png?branch=master)](https://travis-ci.org/ferhatelmas/go-binary)

> Check if a filepath is a binary file.

### Install

```
go get github.com/ferhatelmas/go-binary
```

### Usage

```go
import "github.com/ferhatelmas/go-binary"

binary.Is("src/unicorn.exe")
//=> true

binary.Is("src/unicorn.txt")
//=> false
```

### Related
[gobinext](https://github.com/shamsher31/gobinext)

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
