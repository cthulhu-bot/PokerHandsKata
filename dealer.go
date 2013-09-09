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
//    "strings"
    "encoding/json"
    "log"
)

type Hands struct {}

type HandsId struct {}

type Deck struct {
    Cards []string
}

type Player struct {
    Player string
    Hand []string
}

type Response map[string]interface{}

// Request: GET /hands
// Response: {data:single_poker_hand,error:null}
func (*Hands) Get (ctx *jas.Context) {
    myRand := random(1,52)
    ctx.Data = strconv.Itoa(myRand)
}

// Request: GET /hands/number_of_poker_hands
// Response: {data:multiple_poker_hands,error:null}
func (*HandsId) Get (ctx *jas.Context) {
    var Players []*Player

    cards := []string{"2H"}
    user := &Player{Player: "Frank", Hand: cards}
    user2 := &Player{Player: "Fritz", Hand: cards}
    Players = append(Players, user)
    Players = append(Players, user2)

    b, err := json.Marshal(user)
    fmt.Println(string(b))

//    myRand := random(1,52)
//    d := new(Deck)
//    d.init()
//    var deck = strings.Join(d.Cards, ",")

//    numHands := ctx.Id

    if err != nil {
        log.Println(err)
        return
    }

    ctx.Data = Players
}

func (*Deck) dealAHand() []int {
    var hand []int
    return hand
}

func (d *Deck) dealACard() int {
    return 0
}

func handler(rw http.ResponseWriter, req *http.Request) {
//    numHands := req.URL.Path[len("/foo/"):]
    numHands := req.URL.Path
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, Response{"success":true, "message":numHands})
    return
}

func main() {
    fmt.Println("listening...")

    router := jas.NewRouter(new(Hands), new(HandsId))
    router.BasePath = "/"
    fmt.Println(router.HandledPaths(true))

    //output: GET /hands
    http.Handle(router.BasePath, router)
//    http.Handle(router.BasePath, handler)
//    http.HandleFunc("/foo/", handler)

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
        (*d).Cards = append(d.Cards, strconv.Itoa(i))
    }
}