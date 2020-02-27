package model

/*
Fixture represents info of a game in calendar
*/
type Fixture struct {
	Round     int    `bson:"round" json:"round"`
	Datetime  string `bson:"datetime" json:"datetime"`
	Venue     string `bson:"venue" json:"venue"`
	HomeTeam  string `bson:"home_team" json:"home_team"`
	AwayTeam  string `bson:"away_team" json:"away_team"`
	HomeScore int    `bson:"home_score" json:"home_score"`
	AwayScore int    `bson:"away_score" json:"away_score"`
	Played    bool   `bson:"played" json:"played"`
}
