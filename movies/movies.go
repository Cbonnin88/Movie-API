package movies

import(
	"github.com/Cbonnin88/fiber-api/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// Creating our Movie struct
type Movie struct {
	gorm.Model
	Title  string  `json:"title"`
	Year   int		`json:"year"`
	Director  string  `json:"director"`
	Rating	int		`json:"rating"`
}

// Here we create out movie functions for our API
func RentMovies(c *fiber.Ctx) error {
	db := database.DBConn // adding our database connection
	var movies []Movie
	db.Find(&movies)
	return c.JSON(movies)
}

func RentOneMovie(c *fiber.Ctx) error {
	movieID := c.Params("movie_id") // adding a parameter
	db := database.DBConn
	var movies Movie
	db.Find(&movies, movieID)
	return c.JSON(movies)
}

func NewMovie(c *fiber.Ctx) error {
	db := database.DBConn
	movies := new(Movie)
	if err := c.BodyParser(movies); err != nil {
		c.Status(503).SendString("Unable to find the movie")
	}
	db.Create(&movies) // Inserting a movie into our movie database
	return c.JSON(movies)
}

func DeleteMovie(c *fiber.Ctx) error {
	movieID :=  c.Params("movie_id")
	db := database.DBConn

	var movies Movie
	db.First(&movies, movieID)
	if movies.Title == "" {
		err := c.Status(500).SendString("No book found with title")
		if err != nil {
			return err
		}
	}
	db.Delete(&movies)
	return c.SendString("Movie successfully removed")
}


