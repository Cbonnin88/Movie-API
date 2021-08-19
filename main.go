package main

import (
	"fmt"
	"github.com/Cbonnin88/fiber-api/database"
	"github.com/Cbonnin88/fiber-api/movies"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {

	app := fiber.New() // opening an instance of Fiber
	initDatabase()
	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {

		}
	}(database.DBConn) // Defer until the main function has finished

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}


// Here I set up out routes
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/movie", movies.RentMovies)
	app.Get("/api/v1/movie/:movie_id", movies.RentOneMovie)
	app.Post("/api/v1/movie", movies.NewMovie)
	app.Delete("/api/v1/movie/:movie_id", movies.DeleteMovie)
}
// here we are initializing our database connection
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3","movies.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully open")

	database.DBConn.AutoMigrate(movies.Movie{}) // Creating our automigration
	fmt.Println("Database Migrated")
}
