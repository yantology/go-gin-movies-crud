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

type MovieService struct {
	movies []Movie
}

func (m *MovieService) getMovies(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, m.movies)
}

func (m *MovieService) getMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for _, item := range m.movies {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func (m *MovieService) createMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	m.movies = append(m.movies, movie)
	c.JSON(http.StatusCreated, movie)
}

func (m *MovieService) updateMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for index, item := range m.movies {
		if item.Id == id {
			m.movies = append(m.movies[:index], m.movies[index+1:]...)
			var movie Movie
			if err := c.ShouldBindJSON(&movie); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			movie.Id = id
			m.movies = append(m.movies, movie)
			c.JSON(http.StatusOK, movie)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func (m *MovieService) deleteMovie(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	for index, item := range m.movies {
		if item.Id == id {
			m.movies = append(m.movies[:index], m.movies[index+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Serve the index.html file at the root URL
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	movieService := &MovieService{
		movies: []Movie{
			{Id: "1", Isbn: "4598", Title: "The Lord of the Rings", Director: &Director{FirstName: "Jack", LastName: "Jhones"}},
			{Id: "2", Isbn: "1234", Title: "Inception", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}},
			{Id: "3", Isbn: "5678", Title: "The Matrix", Director: &Director{FirstName: "Lana", LastName: "Wachowski"}},
			{Id: "4", Isbn: "91011", Title: "Interstellar", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}},
			{Id: "5", Isbn: "121314", Title: "The Godfather", Director: &Director{FirstName: "Francis", LastName: "Coppola"}},
		},
	}

	router.GET("/movies", movieService.getMovies)
	router.GET("/movies/:id", movieService.getMovie)
	router.POST("/movies", movieService.createMovie)
	router.PUT("/movies/:id", movieService.updateMovie)
	router.DELETE("/movies/:id", movieService.deleteMovie)

	return router
}

func main() {
	router := setupRouter()

	fmt.Println("Starting server at port 3000")
	log.Fatal(router.Run(":3000"))
}
