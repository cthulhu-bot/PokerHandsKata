package main

import (
    "github.com/coocood/jas"
    "fmt"
    "net/http"
)

type Hello struct {}

type Deck struct {}

func (*Hello) Get (ctx *jas.Context) { // GET /v1/hello
    ctx.Data = "hello world"
    // response: {"data":"hello world","error":null}
}

func main() {
    router := jas.NewRouter(new(Hello))
    router.BasePath = "/v1/"
    fmt.Println(router.HandledPaths(true))
    //output: GET /v1/hello
    http.Handle(router.BasePath, router)
    http.ListenAndServe(":8080", nil)
}