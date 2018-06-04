package auth

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Auth Page \n"))

	user := UserFromContext(r.Context())
	log.Println("User: ", user)
}

func AttachHandlers(r *mux.Router) {
	r.HandleFunc("/login/", IndexHandler)
	r.HandleFunc("/logout/", IndexHandler)
	r.HandleFunc("/register/", IndexHandler)
	r.HandleFunc("/register/verify/", IndexHandler)
	r.HandleFunc("/reset-password/", IndexHandler)
	r.HandleFunc("/reset-password/verify/", IndexHandler)
	r.HandleFunc("/profile/", IndexHandler)

	r.Use(UserContextMiddleware)
}
