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
    "strings"
)

type Hands struct {}

type HandsId struct {}

type Deck struct {
    deck []string
}

// Request: GET /hands
// Response: {data:single_poker_hand,error:null}
func (*Hands) Get (ctx *jas.Context) {
    myRand := random(1,52)
    ctx.Data = strconv.Itoa(myRand)
}

// Request: GET /hands/number_of_poker_hands
// Response: {data:multiple_poker_hands,error:null}
func (*HandsId) Get (ctx *jas.Context) {
//    myRand := random(1,52)
    d := new(Deck)
    d.init()
    var deck = strings.Join(d.deck, ",")

//    numHands := ctx.Id

    ctx.Data = deck
}

func (*Deck) dealAHand() []int {
    var hand []int
    return hand
}

func (d *Deck) dealACard() int {
    return 0
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

func (d *Deck) init() {
    for i := 1; i <= 52; i++ {
        (*d).deck = append(d.deck, strconv.Itoa(i))
    }
}