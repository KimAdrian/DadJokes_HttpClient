package main

import (
	"fmt"
	"github.com/KimAdrian/DadJokes_HttpClient/client"
	"log"
)

func main() {

	joke, err := client.FetchJoke()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(joke)
}
