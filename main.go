package main

import (
	"os"

	"github.com/lpernett/godotenv"
	log "github.com/sirupsen/logrus"

	"go-kuzu/kuzu"
)

func main() {
	log.Info("Initializing...")
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("Loading .env file:", dotEnvErr)
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH is not set in .env file or environment variables")
	}

	log.Info("Connecting to database at", dbPath)
	systemConfig := kuzu.DefaultSystemConfig()
	db := kuzu.DatabaseInit(dbPath, systemConfig)
	conn := kuzu.ConnectionInit(db)
	log.Info("Database connected")
	queryString := "match (p:Page) return count(p)"
	log.Info("Querying database with: ", queryString)
	queryResult := kuzu.ConnectionQuery(conn, queryString)

	if kuzu.QueryResultHasNext(queryResult) {
		log.Info("Results found")
		row := kuzu.QueryResultGetNext(queryResult)
		value := kuzu.FlatTupleGetValue(row, 0)
		intValue := kuzu.ValueGetInt64(value)
		log.Info("Value:", intValue)
	} else {
		log.Info("No results found")
	}

	kuzu.ConnectionDestroy(conn)
	log.Info("Database disconnected")
	kuzu.DatabaseDestroy(db)
	log.Info("Database memory deallocated")
}
