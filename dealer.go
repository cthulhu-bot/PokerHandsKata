package main

// Using rest framework "jas" for routing
import (
    "github.com/coocood/jas"
    "fmt"
    "net/http"
//    "os"
    "math/rand"
    "strconv"
    "time"
)

type Hello struct {}

type Hands struct {}

// Request: GET /hands
// Response: {"data":"data","error":null}
func (*Hands) Get (ctx *jas.Context) {
    // GET /hands/:hands

    myRand := random(1,52)
    ctx.Data = strconv.Itoa(myRand)
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    fmt.Println("listening...")

    router := jas.NewRouter(new(Hands))
    router.BasePath = "/"
    fmt.Println(router.HandledPaths(true))

    //output: GET /hands
    http.Handle(router.BasePath, router)

    // PRODUCTION: port detection added for Heroku's random port assignment
    //err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

    // TESTING: Hardcoded port assigned for testing
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        panic(err)
    }
}