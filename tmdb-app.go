package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"net/http"
)


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

	
	body, err := io.ReadAll(response.Body)
	check(err)

	fmt.Println(string(body))
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