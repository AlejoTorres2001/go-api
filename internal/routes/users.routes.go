package routes
import "net/http"

var UsersSubPath string = "/users"
var UsersIdentifierPath string = "id"

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post User"))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}

var UsersHandlers = map[string]func(w http.ResponseWriter, r *http.Request){
	"CGET": getUsersHandler,
	"POST": postUserHandler,
	"GET": getUserHandler,
	"PUT": updateUserHandler,
	"DELETE": deleteUserHandler,
}

