package config

import (
	"log"
	"os"
	"strings"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func InitDatabase() *gocql.Session {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading ENV")
	}

	host := os.Getenv("CASSANDRA_HOST")
	keyspace := os.Getenv("CASSANDRA_KEYSPACE")
	consistencyStr := os.Getenv("CASSANDRA_CONSISTENCY")

	cluster := gocql.NewCluster(strings.Split(host, ",")...)
	cluster.Keyspace = keyspace

	switch strings.ToUpper(consistencyStr) {
	case "ONE":
		cluster.Consistency = gocql.One
	case "QUORUM":
		cluster.Consistency = gocql.Quorum
	case "ALL":
		cluster.Consistency = gocql.All
	default:
		cluster.Consistency = gocql.Quorum
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}

	return session
}
