package routes
import (
	"fmt"
"github.com/gorilla/mux"
"github.com/nandini584/go_mysql_retailshop/pkg/controllers")
func routes (){
	r:= mux.NewRouter()
	r.HandleFunc("/book",GetBooks)


}