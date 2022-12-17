package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Vectormike/go-books-crud-dbms/pkg/models"
	"github.com/Vectormike/go-books-crud-dbms/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	response, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
	}
	book, err := models.GetBook(id)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b, err := book.CreateBook(book)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		fmt.Println(err)
	}
	book, err := models.DeleteBook(id)
	if err != nil {
		fmt.Println(err)
	}
	response, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
