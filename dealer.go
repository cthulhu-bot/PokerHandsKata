package main

// Using rest framework "jas" for routing
import (
    "github.com/coocood/jas"
    "fmt"
    "net/http"
    "os"
    "math/rand"
    "strconv"
    "time"
//    "strings"
    "encoding/json"
    "log"
//    "io/ioutil"
)

type Game struct {
    Id int
}

type GameId struct {}

type Hands struct {}

type HandsId struct {}

type Users struct {
    Id int
}

type UserId struct {}

type Deal struct {}

type DealId struct {}

type Deck struct {
    AvailCards [52]int
    Cards []string
}

type DeckId struct {
}

type Player struct {
    Name string
    Hand []string
}

// Deal 2 hands
func (d *Deal) Get (ctx *jas.Context) {
    fmt.Println("Deal.Get")
    deck := new(Deck)
    deck.init()
    var Players []*Player
    cards := dealAHand(deck)
    user := &Player{Name: "Hans", Hand: cards}
    cards2 := dealAHand(deck)
    user2 := &Player{Name: "Franz", Hand: cards2}
    Players = append(Players, user)
    Players = append(Players, user2)
    ctx.Data = Players
}

// Deal n hands
func (d *DealId) Get (ctx *jas.Context) {
    numHands := ctx.Id
    var Players []*Player
    cards := []string{"2H"}

    user := &Player{Name: "Frank", Hand: cards}
    for i := 0; i < int(numHands); i++ {
        user := &Player{Name: "Frank", Hand: cards}
        Players = append(Players, user)
    }
//    user2 := &Player{Player: "Fritz", Hand: cards}
//    Players = append(Players, user2)

    b, err := json.Marshal(user)
    fmt.Println(string(b))

//    myRand := random(1,52)
//    d := new(Deck)
//    d.init()
//    var deck = strings.Join(d.Cards, ",")


    if err != nil {
        log.Println(err)
        return
    }

    ctx.Data = Players
}

func (d *Deck) init() {
    for i := 0; i < 52; i++ {
        (*d).AvailCards[i] = 1
    }
    cardArray := []string{"AD","2D","3D","4D","5D","6D","7D","8D","9D","10D","JD","QD","KD","AS","2S","3S","4S","5S","6S","7S","8S","9S","10S","JS","QS","KS","AH","2H","3H","4H","5H","6H","7H","8H","9H","10H","JH","QH","KH","AC","2C","3C","4C","5C","6C","7C","8C","9C","10C","JC","QC","KC"}
    (*d).Cards = cardArray
}

func dealAHand(d *Deck) []string {
    var hand []string
    for i := 0; i < 5; i++ {
        hand = append(hand,dealACard(d))
    }
    fmt.Println(hand)
    return hand
}

func dealACard(d *Deck) string {
    cardIndex := random(1,52)
    for (*d).AvailCards[cardIndex] == 0 {
        cardIndex = random(1,52)
    }
    (*d).AvailCards[cardIndex] = 0
    fmt.Println((*d).AvailCards)
    return (*d).Cards[cardIndex]
}

// Request: GET /game
// Response: {data:single_poker_hand,error:null}
func (g *Game) Get (ctx *jas.Context) {
    myRand := random(1,52)
    ctx.Data = strconv.Itoa(myRand)
}

// Request: GET /hands/number_of_poker_hands
// Response: {data:multiple_poker_hands,error:null}
func (g *GameId) Get (ctx *jas.Context) {
    var Players []*Player
    numHands := ctx.Id

    cards := []string{"2H"}

    user := &Player{Name: "Frank", Hand: cards}
    for i := 0; i < int(numHands); i++ {
        user := &Player{Name: "Frank", Hand: cards}
        Players = append(Players, user)
    }
//    user2 := &Player{Player: "Fritz", Hand: cards}
//    Players = append(Players, user2)

    b, err := json.Marshal(user)
    fmt.Println(string(b))

//    myRand := random(1,52)
//    d := new(Deck)
//    d.init()
//    var deck = strings.Join(d.Cards, ",")


    if err != nil {
        log.Println(err)
        return
    }

    ctx.Data = Players
}

func main() {
    rand.Seed(time.Now().Unix())
    fmt.Println("listening...")

    router := jas.NewRouter(new(Deal), new(DealId))
    router.BasePath = "/"
    fmt.Println(router.HandledPaths(true))
    http.Handle(router.BasePath, router)


    // PRODUCTION: port detection added for Heroku's random port assignment
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

    // TESTING: Hardcoded port assigned for testing
    //err := http.ListenAndServe(":8080", nil)

    if err != nil {panic(err)}
}

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func exists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { 
        fmt.Println("File exists!")
        return true, nil
    }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}