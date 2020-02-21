/*
Package persistence provides methods to get information from MongoDB
*/
package persistence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../model"
)

/*
Persistence - Data structure to handle database access
*/
type Persistence struct {
	URI             string
	client          *mongo.Client
	database        *mongo.Database
	boardCollection *mongo.Collection
}

const database = "tinamar"
const boardCollection = "league_board"

/*
Connect initiates connection to MongoDB instance
*/
func (p *Persistence) Connect() error {
	client, clientErr := mongo.NewClient(options.Client().ApplyURI(p.URI))

	if clientErr != nil {
		return clientErr
	}

	p.client = client

	ctx, cancal := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancal()
	connectErr := client.Connect(ctx)

	if connectErr != nil {
		return connectErr
	}

	p.database = client.Database(database)
	p.boardCollection = p.database.Collection(boardCollection)

	return nil
}

/*
GetLeaderBoard returns map array with league table info from database
*/
func (p *Persistence) GetLeaderBoard() ([]model.Team, error) {
	resultSet := make([]model.Team, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"pos": 1})
	cur, findErr := p.boardCollection.Find(ctx, bson.D{}, findOptions)

	if findErr != nil {
		return resultSet, findErr
	}

	for cur.Next(context.TODO()) {
		var res model.Team
		decErr := cur.Decode(&res)
		if decErr != nil {
			return resultSet, decErr
		}

		resultSet = append(resultSet, res)
	}

	return resultSet, nil
}
