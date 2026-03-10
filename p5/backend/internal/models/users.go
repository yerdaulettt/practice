package models

import "time"

type User struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Gender    string    `json:"gender" db:"gender"`
	BirthDate time.Time `json:"birthDate" db:"birth_date"`
}

type PaginatedResponse struct {
	Data       []User `json:"data"`
	TotalCount int    `json:"totalCount"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
}
