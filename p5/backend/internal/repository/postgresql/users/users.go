package users

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"p5/internal/models"
	"p5/internal/repository/postgresql"
)

type Repository struct {
	db *postgresql.Dialect
}

func NewUserRepository(db *postgresql.Dialect) *Repository {
	return &Repository{db: db}
}

func (r *Repository) filterUsers(filter *models.UserFilter, limit int, offset int) (*sql.Rows, error) {
	var filterParamValues []any
	query := `SELECT id, name, email, gender, birth_date FROM users WHERE 1=1`

	if filter.Name != "" {
		filterParamValues = append(filterParamValues, filter.Name)
		query += ` AND name LIKE '%'||$` + strconv.Itoa(len(filterParamValues)) + `||'%'`
	}

	if filter.Email != "" {
		filterParamValues = append(filterParamValues, filter.Email)
		query += ` AND email like '%'||$` + strconv.Itoa(len(filterParamValues)) + `||'%'`
	}

	if filter.Gender != "" {
		filterParamValues = append(filterParamValues, filter.Gender)
		query += fmt.Sprintf(` AND gender = $%d`, len(filterParamValues))
	}

	if !filter.BirthDateMoreThan.IsZero() {
		filterParamValues = append(filterParamValues, filter.BirthDateMoreThan)
		query += ` AND birth_date >= $` + strconv.Itoa(len(filterParamValues))
	}

	if !filter.BrithDateLessThan.IsZero() {
		filterParamValues = append(filterParamValues, filter.BrithDateLessThan)
		query += ` AND birth_date <= $` + strconv.Itoa(len(filterParamValues))
	}

	if filter.OrderBy != "" {
		query += ` ORDER BY ` + filter.OrderBy

		if filter.Sort == "desc" {
			query += ` desc`
		} else {
			query += ` asc`
		}
	}

	filterParamValues = append(filterParamValues, limit, offset)
	query += ` LIMIT $` + strconv.Itoa(len(filterParamValues)-1) + ` OFFSET $` + strconv.Itoa(len(filterParamValues))

	log.Println(query, filterParamValues)
	rows, err := r.db.DB.Query(query, filterParamValues...)
	return rows, err
}

func (r *Repository) GetPaginatedUsers(filters *models.UserFilter, page int, pageSize int) (models.PaginatedResponse, error) {
	var users []models.User
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM users`
	err := r.db.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	filteredUserRows, err := r.filterUsers(filters, pageSize, offset)
	if err != nil {
		return models.PaginatedResponse{}, err
	}
	defer filteredUserRows.Close()

	for filteredUserRows.Next() {
		var u models.User
		if err := filteredUserRows.Scan(&u.Id, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return models.PaginatedResponse{}, err
		}

		users = append(users, u)
	}

	return models.PaginatedResponse{
		Data:       users,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

func (r *Repository) GetCommonFriends(id1 int, id2 int) ([]models.User, error) {
	var commonFriends []models.User

	query := `SELECT id, name, email, gender, birth_date FROM
	(SELECT friend_id AS id FROM user_friends WHERE user_id = $1) JOIN
	(SELECT friend_id AS id FROM user_friends WHERE user_id = $2) USING(id) JOIN  users USING(id)`

	rows, err := r.db.DB.Query(query, id1, id2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
			return nil, err
		}

		commonFriends = append(commonFriends, u)
	}

	return commonFriends, nil
}
