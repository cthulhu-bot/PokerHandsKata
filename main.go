package main

import (
    "github.com/coocood/jas"
    "fmt"
    "net/http"
    "os"
//    "html"
//    "log"
)

type Hello struct {}

type Deck struct {}

func (*Hello) Get (ctx *jas.Context) { // GET /v1/hello
    ctx.Data = "hello world"
    // response: {"data":"hello world","error":null}
}

func main() {
    fmt.Println("listening...")
//    http.Handler("/foo", fooHandler)
//    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
//        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
//    })

//    log.Fatal(http.ListenAndServe(":8080", nil))

    router := jas.NewRouter(new(Hello))
    router.BasePath = "/v1/"
    fmt.Println(router.HandledPaths(true))
    //output: GET /v1/hello
    http.Handle(router.BasePath, router)
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
        panic(err)
    }
}