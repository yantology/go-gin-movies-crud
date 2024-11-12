package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, item := range movies {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func createMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	c.JSON(http.StatusCreated, movie)
}

func updateMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for index, item := range movies {
		if item.Id == id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			if err := c.ShouldBindJSON(&movie); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			movie.Id = id
			movies = append(movies, movie)
			c.JSON(http.StatusOK, movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func deleteMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for index, item := range movies {
		if item.Id == id {
			movies = append(movies[:index], movies[index+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Serve the index.html file at the root URL
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	movies = append(movies, Movie{
		Id:       "1",
		Isbn:     "4598",
		Title:    "The Lord of the Rings",
		Director: &Director{FirstName: "Jack", LastName: "Jhones"},
	})

	movies = append(movies, Movie{
		Id:       "2",
		Isbn:     "1234",
		Title:    "Inception",
		Director: &Director{FirstName: "Christopher", LastName: "Nolan"},
	})

	movies = append(movies, Movie{
		Id:       "3",
		Isbn:     "5678",
		Title:    "The Matrix",
		Director: &Director{FirstName: "Lana", LastName: "Wachowski"},
	})

	movies = append(movies, Movie{
		Id:       "4",
		Isbn:     "91011",
		Title:    "Interstellar",
		Director: &Director{FirstName: "Christopher", LastName: "Nolan"},
	})

	movies = append(movies, Movie{
		Id:       "5",
		Isbn:     "121314",
		Title:    "The Godfather",
		Director: &Director{FirstName: "Francis", LastName: "Coppola"},
	})

	router.GET("/movies", getMovies)
	router.GET("/movies/:id", getMovie)
	router.POST("/movies", createMovie)
	router.PUT("/movies/:id", updateMovie)
	router.DELETE("/movies/:id", deleteMovie)

	fmt.Println("Starting server at port 3000")
	log.Fatal(router.Run(":3000"))
}
