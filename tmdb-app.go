package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

type Movie struct {
	Adult            	bool      `json:"adult"`
	BackdropPath     	string    `json:"backdrop_path"`
	GenreIDs          []int     `json:"genre_ids"`
	ID               	int       `json:"id"`
	OriginalLanguage 	string    `json:"original_language"`
	OriginalTitle    	string    `json:"original_title"`
	Overview         	string    `json:"overview"`
	Popularity       	float64   `json:"popularity"`
	PosterPath       	string    `json:"poster_path"`
	ReleaseDate      	string    `json:"release_date"`
	Title            	string    `json:"title"`
	Video            	bool      `json:"video"`
	VoteAverage      	float64   `json:"vote_average"`
	VoteCount        	int       `json:"vote_count"`
}

type Dates struct{
	Maximum						string		`json:"maximum"`
	Minimum						string		`json:"minimum"`					
}


func main(){
	// Setup and parse Flags
	dataType := flag.String("type", "playing", "The type of Movie data you would like to get")
	flag.Parse()

	// Build api url based on Flags
	url := endPointBuilder(*dataType)

	// Read API access token from file
	key, err := os.ReadFile("key.txt")
	check(err)

	// Build API request
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+string(key))

	// API Call
	response, err := http.DefaultClient.Do(request)
	check(err)
	defer response.Body.Close()

		// Root JSON structure (only need "results")
	var root struct {
		Results []Movie `json:"results"`
		Dates Dates `json:"dates"`
	}

	check(json.NewDecoder(response.Body).Decode(&root))

		// Print results
	for _, m := range root.Results {
		fmt.Println(m.Title, "-", m.ReleaseDate)
	}

}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func endPointBuilder(data string) string{
	apiURL := "https://api.themoviedb.org/3/movie/"
	postfix := "?language=en-US&page=1"


	switch data{
		case "playing":
			return apiURL + "now_playing" + postfix
		case "popular":
			return apiURL + "popular" + postfix
		case "top":
			return apiURL + "top_rated" + postfix
		case "upcoming":
			return apiURL + "upcoming" + postfix
		default:
			return ""
	}
}