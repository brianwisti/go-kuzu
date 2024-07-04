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
	databases := []kuzu.Database{*db}
	conn := kuzu.ConnectionInit(databases)
	connections := []kuzu.Connection{*conn}
	log.Info("Database connected")
	queryString := "match (p:Page) return count(p)"
	log.Info("Querying database with: ", queryString)
	queryResult := kuzu.ConnectionQuery(connections, queryString)
	results := []kuzu.QueryResult{*queryResult}

	if kuzu.QueryResultHasNext(results) {
		log.Info("Results found")
		row := kuzu.QueryResultGetNext(results)
		rows := []kuzu.FlatTuple{*row}
		value := kuzu.FlatTupleGetValue(rows, 0)
		intValue := kuzu.ValueGetInt64([]kuzu.Value{*value})
		log.Info("Value:", intValue)
	} else {
		log.Info("No results found")
	}

	kuzu.ConnectionDestroy(connections)
	log.Info("Database disconnected")
	kuzu.DatabaseDestroy(databases)
	log.Info("Database memory deallocated")
}
