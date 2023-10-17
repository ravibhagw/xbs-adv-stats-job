package models

type Match struct {
	MatchId   string                       `json:"matchId"`
	Timestamp int                          `json:"timestamp"`
	Teams     map[string]Team              `json:"clubs"`
	Players   map[string]map[string]Player `json:"players"`
}
