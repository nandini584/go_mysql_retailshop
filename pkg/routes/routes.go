package routes
import (
	"fmt"
"github.com/gorilla/mux"
"github.com/nandini584/go_mysql_retailshop/pkg/controllers")

var RegisterBookStoreRoutes = func(router *mux.Router){
    router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{bookid}",controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookid}",controllers.UpdateBookDetailsByID).Methods("PUT")
	router.HandleFunc("/book/{bookid}",controllers.DeleteBookByID).Methods("DELETE")
}
