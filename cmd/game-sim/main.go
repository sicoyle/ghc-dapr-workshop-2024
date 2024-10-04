package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pkg "github.com/dapr-volleyball-demo/pkg"
	dapr "github.com/dapr/go-sdk/client"
)

const (
	maxPoints     = 25
	minPointsDiff = 2
	gameCount     = 100

	pubsubComponentName = "gamepubsub"
	pubsubTopic         = "game"
)

func main() {
	// Create a new client for Dapr using the SDK
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Publish events using Dapr pubsub
	// simulate 100 games to play
	for i := 0; i < gameCount; i++ {
		var game pkg.Game
		game.GameID = i
		game.FirstTeamName = "team" + strconv.Itoa(i)
		game.SecondTeamName = "team" + strconv.Itoa(i+1)
		for {
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			if game.FirstTeamScore >= maxPoints && game.FirstTeamScore-game.SecondTeamScore >= minPointsDiff {
				log.Printf("[%s] team 1 wins: %+v", currentTime, game)
				break
			}

			if game.SecondTeamScore >= maxPoints && game.SecondTeamScore-game.FirstTeamScore >= minPointsDiff {
				log.Printf("[%s] team 2 wins: %+v", currentTime, game)
				break
			}

			// Create a new random source with a seed based on the current time
			source := rand.NewSource(time.Now().UnixNano())
			r := rand.New(source)
			if r.Intn(2) == 0 {
				game.FirstTeamScore++
			} else {
				game.SecondTeamScore++
			}

			fmt.Printf("TODO(@GHC attendees): publish game event data and then uncomment line below\n")

			// fmt.Printf("[%s] Published data: %#v\n", currentTime, game)

			time.Sleep(2 * time.Second)
		}
	}

	// Note: the following is added so the container keeps running for the demo.
	stop := make(chan struct{})
	<-stop // block the main goroutine from exiting
}
