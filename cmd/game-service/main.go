package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dapr/go-sdk/client"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/go-chi/chi/v5"
	"github.com/sicoyle/ghc-dapr-workshop-2024/pkg"
)

var (
	// Note: for now leaving this workaround bc kept getting err without this setup:
	// "error invoking rpc error: code = Canceled desc = grpc: the client connection is closing"
	daprClient, daprClientClose = newDaprClient()
)

func newDaprClient() (client.Client, func()) {
	daprClient, err := client.NewClient()
	if err != nil {
		log.Fatalf("failed to create dapr client: %v", err)
	}
	return daprClient, func() {
		defer daprClient.Close()
	}
}

func main() {
	defer daprClientClose()
	router := chi.NewRouter()

	// curl -X GET http://localhost:3002/score/7
	router.HandleFunc("/score/{gameID}", scoreboardHandler)
	srv := daprd.NewServiceWithMux(":3002", router)

	// Start the Dapr service
	log.Printf("starting service game-service")
	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		log.Printf("error: %v", err)
	}
}

func scoreboardHandler(w http.ResponseWriter, r *http.Request) {
	gameID := chi.URLParam(r, "gameID")
	id, err := strconv.Atoi(gameID)
	if err != nil {
		log.Fatalf("error converting %v", err)
	}

	gameReq := pkg.GameRequest{
		GameID: id,
	}
	b, err := json.Marshal(gameReq)
	if err != nil {
		log.Fatalf("error unmarshalling into game %v", err.Error())
	}

	content := &client.DataContent{
		Data:        b,
		ContentType: "application/json",
	}

	log.Println("TODO(@GHC attendees): invoke scoreboard app at currentscore endpoint")

	// process the response
	log.Println(string(resp))
	w.Header().Set("Access-Control-Allow-Origin", "*") // add this line to set the CORS header
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(string(resp))
}
