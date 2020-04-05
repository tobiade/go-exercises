package book

import (
	"encoding/json"
	"errors"
)

//StoryArc model
type StoryArc string

//Option model
type Option struct {
	Text string
	Arc  StoryArc
}

//Story model
type Story struct {
	Title   string
	Content []string `json:"story"`
	Options []Option
}

//Book model
type Book struct {
	Contents map[StoryArc]Story
}

//Get story for a given story arc
func (b *Book) Get(s StoryArc) (*Story, error) {
	if v, ok := b.Contents[s]; ok {
		return &v, nil
	}
	return nil, errors.New("story not found")

}

//Make our book!
func Make(s string) Book {
	b := make(map[StoryArc]Story)
	json.Unmarshal([]byte(s), &b)
	return Book{Contents: b}
}
