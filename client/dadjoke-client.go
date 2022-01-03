package client

import (
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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

}
