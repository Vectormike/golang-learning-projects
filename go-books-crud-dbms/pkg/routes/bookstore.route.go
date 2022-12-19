package routes

import (
	"github.com/Vectormike/go-books-crud-dbms/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/bookstore", controllers.CreateBookStore).Methods("POST")
	router.HandleFunc("/bookstore", controllers.GetBookStores).Methods("GET")
	router.HandleFunc("/bookstore/{bookId}", controllers.GetBookStore).Methods("GET")
	router.HandleFunc("/bookstore/{bookId}", controllers.UpdateBookStore).Methods("PUT")
	router.HandleFunc("/bookstore/{bookId}", controllers.DeleteBookStore).Methods("DELETE")
}

// Path: go-midprojects/go-books-crud-dbms/pkg/routes/book.route.go
