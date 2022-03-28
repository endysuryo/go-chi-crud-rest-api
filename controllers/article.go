package controllers

import (
	"encoding/json"
	"net/http"

	"go-chi-crud-rest-api/interfaces"

	"github.com/go-chi/chi/v5"
)

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	data := []interfaces.Article{
		{
			Id:      "1",
			Title:   "Article 1",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		},
	}

	response := response{
		StatusCode: 200,
		Message:    "Success Get Article",
		Data:       data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var body interfaces.Article
	json.NewDecoder(r.Body).Decode(&body)

	response := response{
		StatusCode: 201,
		Message:    "Success Create Article",
		Data:       body,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var params = chi.URLParam(r, "id")
	var body interfaces.Article

	var Response response
	json.NewDecoder(r.Body).Decode(&body)

	w.Header().Set("Content-Type", "application/json")

	if params != body.Id {
		Response = response{
			StatusCode: 400,
			Message:    "Id not found",
			Data:       nil,
		}

		w.WriteHeader(http.StatusNotFound)
	} else {
		Response = response{
			StatusCode: 200,
			Message:    "Success Update Article",
			Data:       body,
		}
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(Response)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")

	response := response{
		StatusCode: 200,
		Message:    "Success Delete Article",
		Data:       params,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
