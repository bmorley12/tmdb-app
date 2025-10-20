package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// Stores the root JSON data
type Root struct {
	Results					 []Movie	 `json:"results"`
	Dates							 Dates	 `json:"dates"`
}

// Stores movie details provided from API
type Movie struct {
	Adult            	bool     `json:"adult"`
	BackdropPath     	string   `json:"backdrop_path"`
	GenreIDs          []int    `json:"genre_ids"`
	ID               	int      `json:"id"`
	OriginalLanguage 	string   `json:"original_language"`
	OriginalTitle    	string   `json:"original_title"`
	Overview         	string   `json:"overview"`
	Popularity       	float64  `json:"popularity"`
	PosterPath       	string   `json:"poster_path"`
	ReleaseDate      	string   `json:"release_date"`
	Title            	string   `json:"title"`
	Video            	bool     `json:"video"`
	VoteAverage      	float64  `json:"vote_average"`
	VoteCount        	int      `json:"vote_count"`
}

// Stores date ranges for Now Playing and Upcoming movies
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
	key, err := os.ReadFile(".env")
	check(err)

	// Build API request
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+string(key))

	// API Call
	response, err := http.DefaultClient.Do(request)
	check(err)
	defer response.Body.Close()

	// Parse JSON data and store it in predfined data structures
	var data Root
	check(json.NewDecoder(response.Body).Decode(&data))

	printPretty(*dataType, data)

}

// Quick error chcking function
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Builds api end point based on cli argument
func endPointBuilder(dataType string) string{
	apiURL := "https://api.themoviedb.org/3/movie/"
	postfix := "?language=en-US&page=1"


	switch dataType{
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

// Formatting print for each data type
func printPretty(dataType string, data Root){
		switch dataType{
		case "playing":
			printNowPlaying(data)
		case "popular":
			printPopular(data)
		case "top":
			printTop(data)
		case "upcoming":
			printUpcoming(data)
	}
}

// Prints now playing movies
func printNowPlaying(data Root){
	
	fmt.Println("      Now Playing      ")
	fmt.Printf("%v - %v\n", data.Dates.Minimum, data.Dates.Maximum)
	fmt.Println("=======================")

	for _, m := range data.Results {
		fmt.Println("- ", m.Title)
	}

	fmt.Println("=======================")
}

// Prints upcoming movies
func printUpcoming(data Root){
	
	fmt.Println("        Upcoming       ")
	fmt.Printf("%v - %v\n", data.Dates.Minimum, data.Dates.Maximum)
	fmt.Println("=======================")

	for _, m := range data.Results {
		fmt.Println("- ", m.Title)
	}

	fmt.Println("=======================")
}

// Prints popular movies
func printPopular(data Root){
	fmt.Println("    Popular Movies     ")
	fmt.Println("=======================")

	n := 1
	for _, m := range data.Results {
		fmt.Printf("%02d - %v\n", n, m.Title)
		n++
	}

	fmt.Println("=======================")
}

// Prints top movies
func printTop(data Root){
	fmt.Println("       Top Movies      ")
	fmt.Println("=======================")

	n := 1
	for _, m := range data.Results {
		fmt.Printf("%02d - %v\n", n, m.Title)
		n++
	}

	fmt.Println("=======================")
}