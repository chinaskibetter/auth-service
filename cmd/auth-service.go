package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	err := db.InitMongoClient()
	if err != nil {
		fmt.Errorf("%w: ошибка инициализации БД", err)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/get-tokens", handlers.GetTokensHandler).Methods("POST")
	r.HandleFunc("/refresh-tokens", handlers.RefreshTokensHandler).Methods("POST")
	http.Handle("/", r)
	_ = http.ListenAndServe(":8080", nil)
}
