package main

import (
	"fmt"
	"html/template"
	"net/http"

	"./models"
)

var posts map[string]*models.Post

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", posts)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
	id := GenerateID()
	title := r.FormValue("title")
	content := r.FormValue("content")

	post := models.NewPost(id, title, content)
	posts[post.GetID()] = post

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("Listening in port :3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}
