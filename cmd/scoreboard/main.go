package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pkg "github.com/dapr-volleyball-demo/pkg"

	"github.com/dapr/go-sdk/service/common"
	dapr "github.com/dapr/go-sdk/service/http"
)

const stateStoreComponentName = "TODO(@GHC attendees): fill this field in"

var sub = &common.Subscription{
	PubsubName: "TODO(@GHC attendees): fill this field in",
	Topic:      "TODO(@GHC attendees): fill this field in",
	Route:      "/updatescore",
}

func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3002"
	}

	// Create the new server on appPort and add a topic listener
	s := dapr.NewService(":" + appPort)

	err := s.AddTopicEventHandler(sub, eventHandler)
	if err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	// handle incoming service requests
	if err := s.AddServiceInvocationHandler("/currentscore", getGameScoreboardHandler); err != nil {
		log.Fatalf("error adding invocation handler for scoreboard: %v", err)
	}

	// Start the server
	log.Printf("starting scoreboard service")
	err = s.Start()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listening: %v", err)
	}
}

// eventHandler receives data on the game topic and saves state on game point of 25 or higher for either team.
func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] Subscriber received data %v\n", currentTime, e.Data)

	fmt.Printf("TODO(@GHC attendees): create dapr client to be used in this event handler to save game event data\n")

	// Parse the incoming score message
	var game pkg.Game
	err = json.Unmarshal(e.RawData, &game)
	if err != nil {
		log.Fatalf("error unmarshalling into game %v", err)
	}

	// Save state into the state store if game point or higher (ie point 25 or higher)
	if game.FirstTeamScore >= 25 || game.SecondTeamScore >= 25 {
		fmt.Printf("TODO(@GHC attendees): save game event data and then uncomment the two lines below\n")

		// key := "game_" + strconv.Itoa(game.GameID)
		// fmt.Printf("[%s] Saved game score: %s\n", currentTime, string(e.RawData))
	}

	return false, nil
}

// curl -X POST http://localhost:3001/scoreboard -H "Content-Type: application/json" -d '{"id": 0}'
func getGameScoreboardHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"[%s] echo - ContentType:%s, Verb:%s, data: %s",
		currentTime, in.ContentType, in.Verb, string(in.Data),
	)

	var gameReq pkg.GameRequest
	err = json.Unmarshal(in.Data, &gameReq)
	if err != nil {
		log.Printf("error unmarshalling into gameReq")
		return nil, err
	}

	// Get the state from the state store using the game ID
	fmt.Printf("TODO(@GHC attendees): create dapr client to be used in this event handler to save game event data\n")

	fmt.Printf("TODO(@GHC attendees): get game event data and then uncomment the two lines below\n")

	// key := "game_" + strconv.Itoa(gameReq.GameID)
	// log.Printf("[%s] retrieved state for game: %s", currentTime, string(item.Value))

	out = &common.Content{
		Data:        []byte("TODO(@GHC attendees): fill in the item value here instead of this byte slice comment"),
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
