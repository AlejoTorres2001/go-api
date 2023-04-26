package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func buildPath(subPath string,identifierPath string) string{
	return fmt.Sprintf("%v/{%v}",subPath,identifierPath)
}


func ValidateHandlers(handlers map[string] func(w http.ResponseWriter, r *http.Request)) bool {
	isValid := true
	methods := []string{"CGET","POST","GET","PUT","DELETE"}

	for _,method := range methods {
		if _, ok := handlers[method]; !ok {
			handlers[method] = func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotImplemented)
			}
			isValid = false
		}
	}
	return isValid
}


func RegisterHandlers(handlers map[string] func(w http.ResponseWriter, r *http.Request) , subPath string,identifierPath string, r *mux.Router) {

	var fullPath string = buildPath(subPath,identifierPath)

	if (ValidateHandlers(handlers) == false) {
		log.Println("Handlers are not valid")
	}
	r.HandleFunc(subPath, handlers["CGET"]).Methods("GET")
	r.HandleFunc(subPath, handlers["POST"]).Methods("POST")

	r.HandleFunc(fullPath, handlers["GET"]).Methods("GET")
	r.HandleFunc(fullPath, handlers["PUT"]).Methods("PUT")
	r.HandleFunc(fullPath, handlers["DELETE"]).Methods("DELETE")
}