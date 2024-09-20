package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/colebowl/coles-stream/internal/db"
	"github.com/colebowl/coles-stream/internal/models"
	"github.com/colebowl/coles-stream/templates"
	"github.com/gorilla/mux"
)

func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templates.PostForm(nil).Render(r.Context(), w)
		return
	}

	// Handle POST request
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	post := &models.Post{
		Type:          models.PostType(r.FormValue("type")),
		Description:   r.FormValue("description"),
		PublishStatus: r.FormValue("publish_status"),
		Visibility:    models.Visibility(r.FormValue("visibility")),
	}

	// Handle thoughts
	thoughts := r.Form["thoughts[]"]
	for _, thought := range thoughts {
		post.Thoughts = append(post.Thoughts, models.Thought{Content: thought})
	}

	// Handle tags
	tags := strings.Split(r.FormValue("tags"), ",")
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			post.Tags = append(post.Tags, models.Tag{Name: tag})
		}
	}

	if err := db.CreatePost(post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := db.GetPostByID(uint(id))
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		templates.PostForm(post).Render(r.Context(), w)
		return
	}

	// Handle POST request (similar to NewPostHandler)
	// ... (implement update logic)
}

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetLatestPosts(20)
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	templates.Stream(posts).Render(r.Context(), w)
}
