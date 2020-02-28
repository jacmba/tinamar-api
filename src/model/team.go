package model

/*
Team - Data structure to represent teams info in league table
*/
type Team struct {
	ID      int    `bson:"id" json:"id"`
	Name    string `bson:"name" json:"name"`
	Pos     int    `bson:"pos" json:"pos"`
	Points  int    `bson:"points" json:"points"`
	Played  int    `bson:"played" json:"played"`
	Won     int    `bson:"won" json:"won"`
	Draw    int    `bson:"draw" json:"draw"`
	Lost    int    `bson:"lost" json:"lost"`
	Favour  int    `bson:"favour" json:"favour"`
	Against int    `bson:"against" json:"against"`
}
