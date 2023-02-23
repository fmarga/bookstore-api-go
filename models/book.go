package models

type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicationYear"`
	Pages           int    `json:"pages"`
	Publisher       string `json:"publisher"`
}

var Books []Book
