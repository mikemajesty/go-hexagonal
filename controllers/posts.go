package controllers

import (
	"encoding/json"
	"golang/projects/go-hexagonal/entities"
	"golang/projects/go-hexagonal/services"
	"net/http"
)

var postService services.IPostService = services.Create()

func GetPost(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}
	result, err := json.Marshal(posts)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entities.Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	eer := postService.Validate(&post)
	if eer != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error validating the post"}`))
		return
	}

	postService.Create(&post)

	result, err := json.Marshal(post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the post"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
