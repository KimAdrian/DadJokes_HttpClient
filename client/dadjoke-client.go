package client

import (
	"encoding/json"
	"github.com/KimAdrian/DadJokes_HttpClient/model"
	"io"
	"log"
	"net/http"
)

func FetchJoke() (string, error) {

	url := "https://icanhazdadjoke.com/"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "https://github.com/KimAdrian/DadJokes_HttpClient")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	var jokeResponse model.DadJokeResponse

	//Decode the data
	if err := json.NewDecoder(res.Body).Decode(&jokeResponse); err != nil {
		log.Fatal("An error occurred: ", err)
	}

	return jokeResponse.TextOutput(), err
}
