/*
Package config - provides configuration values
*/
package config

import "os"

/*
MongoURL - connection string to MongoDB
*/
var MongoURL string = getVar("MONGO_URL", "mongodb://localhost:27017")

/*
ServerPort - port where the HTTP service listens to
*/
var ServerPort string = getVar("SERVER_PORT", "3000")

func getVar(key, defValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defValue
	}
	return value
}
