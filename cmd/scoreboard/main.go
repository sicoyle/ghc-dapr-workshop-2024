package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	dapr "github.com/dapr/go-sdk/service/http"
	"github.com/sicoyle/ghc-dapr-workshop-2024/pkg"
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
	log.Printf("[%s] Subscriber received data %v\n", currentTime, e.Data)

	daprClient, err := client.NewClient()
	if err != nil {
		return false, fmt.Errorf("err creating dapr client: %v", err)
	}

	// Parse the incoming score message
	var game pkg.Game
	err = json.Unmarshal(e.RawData, &game)
	if err != nil {
		return false, fmt.Errorf("error unmarshalling into game %v", err)
	}

	// Save state into the state store if game point or higher (ie point 25 or higher)
	if game.FirstTeamScore >= 25 || game.SecondTeamScore >= 25 {
		log.Println("TODO(@GHC attendees): save game event data and then uncomment the two lines below")

		// key := "game_" + strconv.Itoa(game.GameID)
		// log.Printf("[%s] Saved game score: %s\n", currentTime, string(e.RawData))
	}

	return false, nil
}

// curl -X POST http://localhost:3001/currentscore -H "Content-Type: application/json" -d '{"id": 0}'
func getGameScoreboardHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"[%s] echo - ContentType:%s, Verb:%s, data: %s\n",
		currentTime, in.ContentType, in.Verb, string(in.Data),
	)

	var gameReq pkg.GameRequest
	err = json.Unmarshal(in.Data, &gameReq)
	if err != nil {
		log.Printf("error unmarshalling into gameReq")
		return nil, err
	}

	// Get the state from the state store using the game ID
	daprClient, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	key := "game_" + strconv.Itoa(gameReq.GameID)
	item, err := daprClient.GetState(context.Background(), stateStoreComponentName, key, nil)
	if err != nil {
		log.Printf("error getting state for id %d: %v\n", &gameReq.GameID, err)
		return nil, err
	}
	log.Printf("[%s] retrieved state for game: %s\n", currentTime, string(item.Value))

	out = &common.Content{
		Data:        item.Value,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
