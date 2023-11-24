package configs

import (
	"fmt"
	"github.com/dotenv-org/godotenvvault"
	"github.com/tidwall/buntdb"
	"gorm.io/driver/postgres"
	"log"
	"os"

	"gorm.io/gorm"
)

var (
	db *gorm.DB //database
	//eventDB *esdb.Client
	newDatabase *buntdb.DB
)

// init initializes Database connection
func init() {
	// Load .env file
	err := godotenvvault.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Keeps user credentials
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Building a connection string
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	dbStart := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbHost, dbName)
	log.Println(dbStart)

	// Opens the database connection
	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return
	}

	//esdbEndpoint := fmt.Sprintf("{%s}", os.Getenv("ESDB_ENDPOINT"))
	//settings, err := esdb.ParseConnectionString(esdbEndpoint)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//dbEv, err := esdb.NewClient(settings)
	//
	//eventDB = dbEv
	dbTxn, err := buntdb.Open("data.db")
	if err != nil {
		log.Println(err)
	}

	// Keeps the open database connection
	db = conn

	newDatabase = dbTxn

}

// GetDB Returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

func GetTxnDB() *buntdb.DB {
	return newDatabase
}
