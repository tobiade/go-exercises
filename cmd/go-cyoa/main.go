package main

import (
	"fmt"
	"net/http"

	"html/template"

	"github.com/tobiade/go-exercises/cyoa/adventure"
	"github.com/tobiade/go-exercises/cyoa/book"
)

func main() {
	b := getBook()
	t := template.Must(template.New("cyoa").Parse(adventure.GetTemplate()))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		storyArc := r.URL.Path[1:]
		fmt.Printf("Story arc requested: %s\n", storyArc)
		s, err := b.Get(book.StoryArc(storyArc))
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, s)
	})
	http.ListenAndServe(":80", nil) //visit localhost/intro to begin your adventure
}

func getBook() book.Book {
	s := adventure.Get()
	return book.Make(s)
}
