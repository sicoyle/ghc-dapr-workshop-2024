package pkg

type Game struct {
	GameID          int    `json:"gameID"`
	FirstTeamName   string `json:"firstTeamName"`
	SecondTeamName  string `json:"secondTeamName"`
	FirstTeamScore  int    `json:"firstTeamScore"`
	SecondTeamScore int    `json:"secondTeamScore"`
}

type GameRequest struct {
	GameID int `json:"gameID"`
}
