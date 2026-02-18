package handlers

import (
	"encoding/json"
	"fmt"
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
		return
	}

	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		w.Write([]byte(`{"error": "json error!"}`))
	}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user modules.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte(`{"error":"json error!"}`))
		return
	}

	id := useCaseForUser.NewUser(user)
	if id == -1 {
		w.Write([]byte(`{"error":"not found or error!"}`))
		return
	}

	response := fmt.Sprintf(`{"new user with id":"%d"}`, id)
	w.Write([]byte(response))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	deletedUser := useCaseForUser.DeleteUser(id)
	if deletedUser == nil {
		w.Write([]byte(`{"error":"no user with such id"}`))
		return
	}

	err = json.NewEncoder(w).Encode(deletedUser)
	if err != nil {
		w.Write([]byte(`{"error":"json error!"}`))
		return
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	user := useCaseForUser.GetUserByid(id)
	if user == nil {
		w.Write([]byte(`{"error":"user not found or error"}`))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.Write([]byte(`{"error":"json error!"}`))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	var userToUpdate modules.User
	err = json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		w.Write([]byte(`{"error":"json error!"}`))
		return
	}

	updatedUser := useCaseForUser.UpdateUser(id, userToUpdate)
	if updatedUser == nil {
		w.Write([]byte(`{"error":"not found!"}`))
		return
	}

	err = json.NewEncoder(w).Encode(&updatedUser)
	if err != nil {
		w.Write([]byte(`{"error":"json error!"}`))
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"OK"}`))
}
