package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"p3/internal/usecase"
	"p3/pkg/modules"
)

var useCaseForUser = usecase.NewUserUseCase(dbStart())

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := useCaseForUser.GetUsers()
	if users == nil {
		w.Write([]byte(`{"error":"something went wrong!"}`))
	}

	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		w.Write([]byte(`{"error": "json error!"}`))
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	user := useCaseForUser.GetUserbyid(id)
	var nilUser modules.User
	if user == nilUser {
		w.Write([]byte(`{"error":"..."}`))
		return
	}

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		w.Write([]byte(`{"error":"json error"}`))
	}
}
