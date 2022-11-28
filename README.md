# Introduction

`godo` stands for Go Digital Ocean, it is a library that makes working with
serverless functions much easier. It introduces structs that abstract the
interface between Digital Ocean's serverless infrastructure platform and Go.

The conventional way to declare a serverless function is like so:

```go
func Main(args map[string]interface{}) map[string]interface{} {
    // ...
    return r
}
```

The library introduces `godo.Request` and `godo.ResponseMap` which change the
function like so:

```go
func Main(args Request) ResponseMap {
    r := NewResponseJSON()
    r.Body = args.GetBody()
    r.StatusCode = http.StatusOK
    return r.ToMap()
}
```

Currently, the interface offered in this module is simple, however, it makes
working with serverless functions in Digital Ocean much easier and cuts down
development time.

## Go Package

The package can be found at the [Go packages website](https://pkg.go.dev/github.com/Yiannis128/godo).
