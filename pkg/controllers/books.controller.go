package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shubham/bookstore/pkg/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	err = json.Unmarshal(body, &book)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	models.DB.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	models.DB.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	id := r.URL.Query().Get("id")
	models.DB.First(&book, id)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	id := r.URL.Query().Get("id")
	models.DB.First(&book, id)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)

		return
	}
	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)

		return
	}
	models.DB.Save(&book)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := models.Book{}
	id := r.URL.Query().Get("id")
	models.DB.First(&book, id)
	models.DB.Delete(&book)
	json.NewEncoder(w).Encode(book)
	fmt.Fprintf(w, "Book with id %s deleted successfully", id)

}
