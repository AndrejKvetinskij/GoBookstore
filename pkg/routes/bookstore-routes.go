package routes

import (
	c "github.com/andrejkvetinskij/gobookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", c.CreateBook).Methods("POST")
	router.HandleFunc("/book/", c.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", c.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", c.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", c.DeleteBook).Methods("DELETE")
}
