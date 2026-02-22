package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"p3/internal/usecase"
	"p3/pkg/modules"
)

var useCaseForUser = usecase.NewUserUseCase(dbStart(), initRedis())

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := useCaseForUser.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	if users == nil {
		w.Write([]byte(`{"message":"empty!"}`))
		return
	}

	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "json error!"}`))
	}
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user modules.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"json error"}`))
		return
	}

	if user.Name == "" || user.Age == 0 || user.Hobby == "" || user.Profession == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"please, fill all the data"}`))
		return
	}

	id, err := useCaseForUser.NewUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := fmt.Sprintf(`{"new user with id":"%d"}`, id)
	w.Write([]byte(response))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	deletedUser, err := useCaseForUser.DeleteUser(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"no user with this id to delete"}`))
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			return
		}
	}

	err = json.NewEncoder(w).Encode(deletedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"json error!"}`))
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	user, err := useCaseForUser.GetUserByid(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"user not found!"}`))
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			return
		}
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"json error!"}`))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"not a number!"}`))
		return
	}

	var userToUpdate modules.User
	err = json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"json error!"}`))
		return
	}

	updatedUser, err := useCaseForUser.UpdateUser(id, userToUpdate)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"user not found"}`))
			return
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			return
		}
	}

	err = json.NewEncoder(w).Encode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"json error!"}`))
	}
}

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"OK"}`))
}
