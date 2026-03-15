package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"p5/internal/models"
	"p5/internal/usecase"
)

var userUsecase = usecase.NewUserUsecase(initDB())

func GetPaginatedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var filter models.UserFilter

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		w.Write([]byte(`{"error":"page is not a number!"}`))
		return
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		w.Write([]byte(`{"error":"page_size is not a number!"}`))
		return
	}

	orderOption := r.URL.Query().Get("order_by")
	switch orderOption {
	case "id":
		filter.OrderBy = "id"
	case "name":
		filter.OrderBy = "name"
	case "email":
		filter.OrderBy = "email"
	case "gender":
		filter.OrderBy = "gender"
	case "birth_date":
		filter.OrderBy = "birth_date"
	}

	switch r.URL.Query().Get("sort") {
	case "desc":
		filter.Sort = "desc"
	default:
		filter.Sort = "asc"
	}
	filter.Name = r.URL.Query().Get("name")
	filter.Email = r.URL.Query().Get("email")
	filter.Gender = r.URL.Query().Get("gender")
	filter.BirthDateMoreThan, _ = time.Parse("2006-01-02", r.URL.Query().Get("bd_more_than"))
	filter.BrithDateLessThan, _ = time.Parse("2006-01-02", r.URL.Query().Get("bd_less_than"))

	users, err := userUsecase.GetPaginatedUsers(&filter, page, pageSize)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func GetCommonFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id1, err := strconv.Atoi(r.URL.Query().Get("id1"))
	if err != nil {
		w.Write([]byte(`{"error":"id1 not specified"}`))
		return
	}
	id2, err := strconv.Atoi(r.URL.Query().Get("id2"))
	if err != nil {
		w.Write([]byte(`{"error":"id2 not specified"}`))
		return
	}

	commonFriends, err := userUsecase.GetCommonFriends(id1, id2)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(`{"error":"some error occured"}`))
		return
	}

	err = json.NewEncoder(w).Encode(&commonFriends)
	if err != nil {
		w.Write([]byte(`{"error":"json error"}`))
	}
}
