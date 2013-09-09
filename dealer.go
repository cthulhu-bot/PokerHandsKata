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

type Hands struct {}

type HandsId struct {}

// Request: GET /hands
// Response: {data:single_poker_hand,error:null}
func (*Hands) Get (ctx *jas.Context) {
    myRand := random(1,52)
    ctx.Data = strconv.Itoa(myRand)
}

// Request: GET /hands/number_of_poker_hands
// Response: {data:multiple_poker_hands,error:null}
func (*HandsId) Get (ctx *jas.Context) {
    myRand := random(1,52)
    numHands = strconv.Atoi(ctx.id)

    ctx.Data = strconv.Itoa(myRand)
}

func main() {
    fmt.Println("listening...")

    router := jas.NewRouter(new(Hands), new(HandsId))
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

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

