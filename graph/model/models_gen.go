// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateRecipeInput struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}

type Recipe struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}