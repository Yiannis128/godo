// Package godo - Stands for Go Digital Ocean, it is a library that makes
// working with serverless functions much easier. The 2 main structs offered
// here are godo.Request and godo.Response. The way they are meant to be used is
// that they should replace the Main method arguments and return type like so:
//
//	package main
//
//	import do "github.com/Yiannis128/godo"
//
//	func Main(args do.Request) do.ResponseMap {
//	    r := do.NewResponseJSON()
//	    r.Body = args.GetBody()
//	    r.StatusCode = http.StatusOK
//	    return r.ToMap()
//	}
//
// Currently, the interface offered in this module is simple, however, it makes
// working with serverless functions in Digital Ocean much easier.
package godo
