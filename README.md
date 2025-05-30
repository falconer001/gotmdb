# Go TMDB Wrapper

A Go client library for The Movie Database (TMDb) API v3 and v4(not complete). This library provides type-safe access to TMDb API endpoints with proper Go types and idiomatic interfaces.

## Features

- Type-safe access to TMDb API endpoints
- Support for Movies and TV Shows endpoints
- Support for Discover endpoint and more
- Strongly typed responses and requests
- Simple and straightforward API design

## Installation

```bash
go get github.com/falconer001/gotmdb
```

## Quick Start

First, get your TMDb API key from [The Movie Database](https://www.themoviedb.org/settings/api).

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/falconer001/gotmdb"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the TMDb client
	tmdb, err := gotmdb.New(gotmdb.Config{
		APIKey:      os.Getenv("TMDB_API_KEY"),
		BearerToken: os.Getenv("TMDB_BEARER_TOKEN"),
	})
	if err != nil {
		log.Fatalf("Error creating TMDb client: %v", err)
	}

	// Get movie details
	movie, err := tmdb.Movies.GetDetails(585511).Exec()
	if err != nil {
		log.Fatalf("Error getting movie: %v", err)
	}

  // Get tv show details
	tvShow, err := tmdb.TV.GetDetails(1399).Exec()
	if err != nil {
		log.Fatalf("Error getting tv show: %v", err)
	}

  // Search for a movie or tv show
  searchResults, err := tmdb.Search.GetMulti("luck").Exec()
  if err != nil {
    log.Fatalf("Error searching: %v", err)
  }
	

	// Access typed movie properties
	fmt.Printf("Title: %s\n", movie.Title)
	fmt.Printf("Overview: %s\n", movie.Overview)
	fmt.Printf("Release Date: %s\n", movie.ReleaseDate)
}
```

## Available Endpoints

### Movies and TV Shows

- Get Details (with AppendToResponse support for both movies and TV shows)
- Get Similar
- Get Recommendations
- Get Popular
- Get Top Rated
- Get Upcoming
- Get Now Playing
- Get Credits
- Get Images
- Get Videos
- Get Aggregate Credits (TV Shows only)
- Get On The Air (TV Shows only)
- Get Airing Today (TV Shows only)
- Rate Movie and TV Show
- Get Account States (Movies and TV Shows)

### Search

- Get Multi (Search for movies and TV shows)
- Get Companies
- Get Collections
- Get Keywords
- Get People
....and more

## Type System

All TMDb API responses are mapped to strongly-typed Go structs in the `github.com/falconer001/gotmdb/types` package. Here are some commonly used types:

### Movie Types

```go

import "github.com/falconer001/gotmdb/types"

var movie types.MovieDetails

// MovieDetails contains detailed information about a movie
type MovieDetails struct {
    ID               int32         `json:"id"`
    Title            string        `json:"title"`
    OriginalTitle    string        `json:"original_title"`
    Overview         string        `json:"overview"`
    PosterPath       string        `json:"poster_path"`
    BackdropPath     string        `json:"backdrop_path"`
    ReleaseDate      string        `json:"release_date"`
    Runtime          int           `json:"runtime"`
    VoteAverage      float32       `json:"vote_average"`
    VoteCount        int           `json:"vote_count"`
    Genres           []Genre       `json:"genres"`
    ProductionCos    []Company     `json:"production_companies"`
    ProductionCos    []Country     `json:"production_countries"`
    SpokenLanguages  []Language    `json:"spoken_languages"`
    Status           string        `json:"status"`
    Tagline          string        `json:"tagline"`
    // ... and many more fields
}
```

### TV Show Types

```go
import "github.com/falconer001/gotmdb/types"

var tvShow types.TVDetails

// TVDetails contains detailed information about a TV show
type TVDetails struct {
    ID               int32         `json:"id"`
    Name             string        `json:"name"`
    OriginalName     string        `json:"original_name"`
    Overview         string        `json:"overview"`
    PosterPath       string        `json:"poster_path"`
    BackdropPath     string        `json:"backdrop_path"`
    FirstAirDate     string        `json:"first_air_date"`
    LastAirDate      string        `json:"last_air_date"`
    NumberOfSeasons  int           `json:"number_of_seasons"`
    NumberOfEpisodes int           `json:"number_of_episodes"`
    VoteAverage      float32       `json:"vote_average"`
    VoteCount        int           `json:"vote_count"`
    EpisodeRunTime   []int         `json:"episode_run_time"`
    Genres           []Genre       `json:"genres"`
    CreatedBy        []Creator     `json:"created_by"`
    Networks         []Network     `json:"networks"`
    ProductionCos    []Company     `json:"production_companies"`
    Seasons          []Season      `json:"seasons"`
    // ... and many more fields
}
```

## Error Handling

All API errors are returned as Go errors. You can check for specific error types:

```go
_, err := tmdb.Movies.GetDetails(585511).Exec()
if err != nil {
    // Check for specific error types
    if strings.Contains(err.Error(), "not found") {
        log.Println("Movie not found")
    } else if strings.Contains(err.Error(), "Invalid API key") {
        log.Fatal("Invalid API key")
    } else {
        log.Fatalf("API error: %v", err)
    }
}
```

## Appending Responses

Some endpoints support appending additional data to the response:

```go
// Get movie details with credits and images
movie, err := tmdb.Movies.GetDetails(585511).
    AppendToResponse("credits,images"). // Append additional data to the response (allowed values are other movie endpoints)
    Exec()
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT
