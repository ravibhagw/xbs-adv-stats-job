package models

type Team struct {
	ClubDivision        string      `json:"clubDivision"`
	CNhlOnlineGameType  string      `json:"cNhlOnlineGameType"`
	Garaw               string      `json:"garaw"`
	Gfraw               string      `json:"gfraw"`
	Losses              string      `json:"losses"`
	MemberString        string      `json:"memberString"`
	OpponentClubId      string      `json:"opponentClubId"`
	OpponentScore       string      `json:"opponentScore"`
	OpponentTeamArtAbbr string      `json:"opponentTeamArtAbbr"`
	Passa               string      `json:"passa"`
	Passc               string      `json:"passc"`
	Ppg                 string      `json:"ppg"`
	Ppo                 string      `json:"ppo"`
	Result              string      `json:"result"`
	Score               string      `json:"score"`
	ScoreString         string      `json:"scoreString"`
	Shots               string      `json:"shots"`
	TeamArtAbbr         string      `json:"teamArtAbbr"`
	TeamSide            string      `json:"teamSide"`
	Toa                 string      `json:"toa"`
	WinnerByDnf         string      `json:"winnerByDnf"`
	WinnerByGoalieDnf   string      `json:"winnerByGoalieDnf"`
	TeamDetails         TeamDetails `json:"details"`
	Goals               string      `json:"goals"`
	GoalsAgainst        string      `json:"goalsAgainst"`
}