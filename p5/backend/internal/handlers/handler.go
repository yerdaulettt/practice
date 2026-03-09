package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"p5/internal/models"
	"p5/internal/repository"
	"p5/internal/repository/postgresql"
	"p5/internal/usecase"
)

func initDB() *repository.Repositories {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	dbconfig := &models.PostgresConfiguration{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbname,
		SSLMode:  ssl,
	}

	db := postgresql.NewDialect(dbconfig)
	repos := repository.NewRepositories(db)

	return repos
}

var userUsecase = usecase.NewUserUsecase(initDB())

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.Write([]byte(`{"error":"not a number! Enter number please"}`))
		return
	}

	user, err := userUsecase.GetUserByID(id)
	if err != nil {
		// w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		w.Write([]byte(`{"error":"no such user with that id"}`))
		return
	}

	if user == nil {
		w.Write([]byte(`{"error":"nil user"}`))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	}
}

func GetPaginatedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	page := r.URL.Query().Get("page")
	pagesize := r.URL.Query().Get("pagesize")

	pagei, err := strconv.Atoi(page)
	if err != nil {
		w.Write([]byte(`{"error":"not a number page!"}`))
		return
	}

	pagesizei, err := strconv.Atoi(pagesize)
	if err != nil {
		w.Write([]byte(`{"error":"not a number pagesize"}`))
		return
	}

	response, err := userUsecase.GetPaginatedUsers(pagei, pagesizei)
	if err != nil {
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	}
}
