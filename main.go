package main

import (
	"demo-shop-manager/handlers"
	"demo-shop-manager/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ben-fiw/go-database-bundle"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

// Define a struct for your API response
type Response struct {
	Message string `json:"message"`
}

// Define your API handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Create a sample response
	response := Response{Message: "Hello, World!"}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Encode the response as JSON and send it
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup the database
	databaseDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	database.InitConnection(databaseDsn, "mysql")
	dbConnection, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	// Run migrations
	migrator := database.GetMigrator()
	migrator.AddMigrationLoader(database.NewJsonMigrationLoader("migrations", ".migration.json"))
	migrations, err := migrator.LoadMigrations()
	if err != nil {
		log.Fatal(err)
	}
	err = migrations.Migrate(dbConnection)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize model stores
	models.InitAvailableVersionModelStore(dbConnection)
	models.InitDemoInstanceModelStore(dbConnection)

	// Get the 'APP_PORT'
	port := os.Getenv("APP_PORT")

	// Create a new router instance
	router := mux.NewRouter()

	// Load and register the handlers
	handlers.GetRegistry().RegisterHandlers(router)

	// Start the HTTP server
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
