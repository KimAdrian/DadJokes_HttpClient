package model

import "fmt"

type DadJokeResponse struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func (joke DadJokeResponse) TextOutput() string {
	p := fmt.Sprintf(
		"id: %s\njoke: %s\nstatus: %d\n",
		joke.Id, joke.Joke, joke.Status)
	return p
}
