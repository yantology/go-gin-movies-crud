# go-gin-movies-crud

## Introduction

This project is a simple CRUD (Create, Read, Update, Delete) application for managing movies using the Go programming language and the Gin web framework. It provides a RESTful API to perform operations on a list of movies.

## Project Structure

```plaintext
go-gin-movies-crud/ 
├── go.mod
|── go.sum 
├── main.go 
├── static/ 
│ └── index.html
  └── README.md
```

## Libraries Used

- [Gin](https://github.com/gin-gonic/gin): A web framework written in Go.
- [exp/rand](https://pkg.go.dev/golang.org/x/exp/rand): A package for random number generation.

## Setup Instructions

### Development

1. Clone the repository:

    ```sh
    git clone https://github.com/yantology/go-gin-movies-crud.git
    cd go-gin-movies-crud
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Run the application:

    ```sh
    go run main.go
    ```

4. Open your browser and navigate to `http://localhost:3000`.

### Production

1. Build the application:

    ```sh
    go build -o go-gin-movies-crud main.go
    ```

2. Run the executable:

    ```sh
    ./go-gin-movies-crud
    ```

### Docker

1. Create dockerfile(Setup look dockerfile in this project)

2. Build the Docker image:

    ```sh
    docker build -t go-gin-movies-crud .
    ```

3. Run the Docker container:

    ```sh
    docker run -p 3000:3000 go-gin-movies-crud
    ```

## Running Tests

1. Run the tests:

    ```sh
    go test ./...
    ```

2. Check test coverage:

    ```sh
    go test -cover ./...
    ```

3. To generate a detailed coverage report, use the following command:

    ```sh
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
    ```

    or

    ```sh
    go test -coverprofile coverage.out ./...
    go tool cover -html coverage.out
    ```

## API Endpoints

### Get All Movies

- **URL:** `/movies`
- **Method:** `GET`
- **Description:** Retrieve a list of all movies.
- **Example Response:**

    ```json
    [
        {
            "id": "1",
            "isbn": "4598",
            "title": "The Lord of the Rings",
            "director": {
                "firstname": "Jack",
                "lastname": "Jhones"
            }
        },
        // ...other movies
    ]
    ```

### Get a Single Movie

- **URL:** `/movies/:id`
- **Method:** `GET`
- **Description:** Retrieve a single movie by its ID.
- **Example Response:**

    ```json
    {
        "id": "1",
        "isbn": "4598",
        "title": "The Lord of the Rings",
        "director": {
            "firstname": "Jack",
            "lastname": "Jhones"
        }
    }
    ```

### Create a Movie

- **URL:** `/movies`
- **Method:** `POST`
- **Description:** Create a new movie.
- **Example Request:**

    ```json
    {
        "isbn": "1234",
        "title": "Inception",
        "director": {
            "firstname": "Christopher",
            "lastname": "Nolan"
        }
    }
    ```

- **Example Response:**

    ```json
    {
        "id": "6",
        "isbn": "1234",
        "title": "Inception",
        "director": {
            "firstname": "Christopher",
            "lastname": "Nolan"
        }
    }
    ```

### Update a Movie

- **URL:** `/movies/:id`
- **Method:** `PUT`
- **Description:** Update an existing movie by its ID.
- **Example Request:**

    ```json
    {
        "isbn": "5678",
        "title": "The Matrix Reloaded",
        "director": {
            "firstname": "Lana",
            "lastname": "Wachowski"
        }
    }
    ```

- **Example Response:**

    ```json
    {
        "id": "3",
        "isbn": "5678",
        "title": "The Matrix Reloaded",
        "director": {
            "firstname": "Lana",
            "lastname": "Wachowski"
        }
    }
    ```

### Delete a Movie

- **URL:** `/movies/:id`
- **Method:** `DELETE`
- **Description:** Delete a movie by its ID.
- **Example Response:** `204 No Content`

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.
