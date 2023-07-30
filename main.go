package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "time"
) 

type Song struct {
  Name string
  Artist string
}

func main() {
    fmt.Println("hello world")

    h1 := func (w http.ResponseWriter, r *http.Request) {
      tmpl := template.Must(template.ParseFiles("index.html"))
      songs := map[string][]Song{
        "Songs": {
          {Name: "Runaway", Artist: "Kanye West"},
          {Name: "Wish You Were Here", Artist: "Pink Floyd"},
          {Name: "ALL CAPS", Artist: "MF DOOM"},
        },
      }
      tmpl.Execute(w, songs)
    } 

    // handler function #2 - returns the template block with the newly added film, as an HTMX response
	  h2 := func(w http.ResponseWriter, r *http.Request) {
		  time.Sleep(1 * time.Second)
		  name := r.PostFormValue("name")
		  artist := r.PostFormValue("artist")
		  tmpl := template.Must(template.ParseFiles("index.html"))
		  tmpl.ExecuteTemplate(w, "song-list-element", Song{Name: name, Artist: artist})
	  }
    http.HandleFunc("/", h1)
    http.HandleFunc("/add-song/", h2)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
