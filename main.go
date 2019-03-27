package main

import (
	"blog/models"
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"github.com/martini-contrib/render"
)

var posts map[string]*models.Post

func indexHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	rnd.HTML(200, "index", posts)
}

func writeHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {
	rnd.HTML(200, "write", nil)
}

func editHandler(rnd render.Render, w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	post, found := posts[id]

	if !found {
		http.NotFound(w, r)
	}

	rnd.HTML(200, "write", post)
}

func savePostHandler(rnd render.Render, w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")
	title := r.FormValue("title")
	text := r.FormValue("text")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Text = text
	} else {
		id = GenerateId()

		post := models.NewPost(id, title, text)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func deleteHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")

	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

//func getPostHandler(w http.ResponseWriter, r *http.Request)  {
//	post_id := r.URL.Query()["id"]
//
//	fmt.Println(posts[post_id]);
//}

func main() {
	fmt.Println("Hello");

	posts = make(map[string]*models.Post, 0)

	m :=  martini.Classic();
	m.Use(render.Renderer(render.Options{
		Directory: "templates", // Specify what path to load the templates from.
		Layout: "layout", // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset: "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true, // Output human readable JSON
		IndentXML: true, // Output human readable XML
		HTMLContentType: "text/html", // Output XHTML content type instead of default "text/html"
	}))

	m.Get("/", indexHandler)
	m.Get("/write", writeHandler)
	m.Post("/savePost", savePostHandler)
	m.Get("/edit", editHandler)
	m.Get("/delete", deleteHandler)

	m.Run()
}
