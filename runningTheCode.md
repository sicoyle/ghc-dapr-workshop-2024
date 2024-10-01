# How to run the code

## Dapr Volleyball Services 

### Volleyball Game Simulator

Volleyball Game Simulator simulates a volleyball game scenario where there are 100 volleyball games.
It randomly adds a point to one of two teams during the game until one team wins by 2.
Game point in volleyball is set to 25, but there is no cap in our simulation.
As the game continues, it sends score updates onto the `game` topic of the `gamepubsub` pubsub.

```
cd cmd/game-sim \

dapr run \
--app-id game-sim \
--app-protocol http \
--dapr-http-port 3500 \
--resources-path ../../resources -- go run .
```

No app port
Dapr port: 3500

To verify the command ran successfully, you'll see something similar to the following:
```shell
dapr run \
--app-id game-sim \
--app-protocol http \
--dapr-http-port 3500 \
--resources-path ../../resources -- go run .
ℹ️  Starting Dapr with id game-sim. HTTP Port: 3500. gRPC Port: 53853
ℹ️  Checking if Dapr sidecar is listening on HTTP port 3500
Flag --dapr-http-max-request-size has been deprecated, use '--max-body-size 4Mi'
Flag --dapr-http-read-buffer-size has been deprecated, use '--read-buffer-size 4Ki'
INFO[0000] Starting Dapr Runtime -- version 1.14.4 -- commit 583960dc90120616124b60ad2b7820fc0b3edf44  app_id=game-sim instance=Cassandras-MacBook-Pro.local scope=dapr.runtime type=log ver=1.14.4
...
INFO[0000] Scheduler stream connected                    app_id=game-sim instance=Cassandras-MacBook-Pro.local scope=dapr.runtime.scheduler type=log ver=1.14.4
INFO[0000] Placement tables updated, version: 0          app_id=game-sim instance=Cassandras-MacBook-Pro.local scope=dapr.runtime.actors.placement type=log ver=1.14.4
ℹ️  Checking if Dapr sidecar is listening on GRPC port 53853
ℹ️  Dapr sidecar is up and running.
ℹ️  Updating metadata for appPID: 6961
ℹ️  Updating metadata for app command: go run .
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == dapr client initializing for: 127.0.0.1:53853
== APP == [2024-10-01 13:49:42] Published data: types.Game{ID:0, Round:1, Team1Name:"team0", Team2Name:"team1", Team1Score:0, Team2Score:1}
== APP == [2024-10-01 13:49:48] Published data: types.Game{ID:0, Round:2, Team1Name:"team0", Team2Name:"team1", Team1Score:0, Team2Score:2}
== APP == [2024-10-01 13:49:54] Published data: types.Game{ID:0, Round:3, Team1Name:"team0", Team2Name:"team1", Team1Score:0, Team2Score:3}
== APP == [2024-10-01 13:50:00] Published data: types.Game{ID:0, Round:4, Team1Name:"team0", Team2Name:"team1", Team1Score:1, Team2Score:3}
...
```

This means the dapr enabled app, with the app-id of `game-sim` is running successfully and publishing data to our message broker, redis.

Optionally: [See inside the Redis message broker by following these steps](./checkRedis.md#how-to-check-the-redis-message-broker-while-running-the-game-simulator)

### Scoreboard API

Scoreboard API Service is a Dapr service that saves volleyball game state,
and provides an API to retrieve game scores using Dapr topic event and service invocation handlers.
It listens to incoming game score update events on the `gamepubsub` pubsub `game` topic,
and any game score that is game point (25) or higher it will save to the statestore.
Specific game score may be found using this API when provided a game ID.

```
cd cmd/scoreboard \

dapr run \
  --app-port 3001 \
  --app-id scoreboard \
  --app-protocol http \
  --dapr-http-port 3501 \
  --resources-path=../../resources -- go run .
```

App port 3001
Dapr port: 3501


### Game Service

Game Service is a Dapr service that provides an interface for the web UI to interact with the system.
It has a `scoreboard` endpoint that invokes service invocation on the `scoreboard` service to retrieve game score for a specific game ID to display on the web UI.

```
cd cmd/game-service \

dapr run \
--app-id game-service \
--app-port 3002 \
--app-protocol http \
--dapr-http-port 3502 \
--resources-path ../../resources -- go run .
```

App Port: 3002
Dapr port: 3502

## Access the User Interface

The Web User Interface (UI) displays volleyball game score information.

```
cd web-ui/

npm install
npm start
```

UI can be reached at: http://localhost:3000/