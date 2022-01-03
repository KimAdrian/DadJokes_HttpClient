package model

import "fmt"

type dadJokeResponse []struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func (joke dadJokeResponse) TextOutput() string {
	p := fmt.Sprintf(
		"id: %s\njoke: %s\nstatus: %d\n",
		joke[0].Id, joke[0].Joke, joke[0].Status)
	return p
}
