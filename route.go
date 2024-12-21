package main

import (
	"encoding/json"
	"golang/projects/go-hexagonal/entities"
	repository "golang/projects/go-hexagonal/repositories"
	"net/http"
)

var PostRepository repository.IPostRepository = repository.Create()

func getPost(response http.ResponseWriter, _ *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := PostRepository.FindAllPosts()

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

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entities.Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	posts, err := PostRepository.FindAllPosts()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}

	var id uint = uint(len(posts)) + 1

	post.ID = id

	PostRepository.SavePost(&post)

	result, err := json.Marshal(post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the post"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}
