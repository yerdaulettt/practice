package users

import (
	"log"
	"p5/internal/models"
	"p5/internal/repository/postgresql"
)

type Repository struct {
	db *postgresql.Dialect
}

func NewUserRepository(db *postgresql.Dialect) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id = $1"
	err := r.db.DB.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email, &user.Gender, &user.BirthDate)
	if err != nil {
		log.Println("error", err)
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetPaginatedUsers(page int, pageSize int) (models.PaginatedResponse, error) {
	var users []models.User
	offset := (page - 1) * pageSize

	var totalCount int
	countQuery := `SELECT COUNT(*) FROM users`
	err := r.db.DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	query := `SELECT id, name, email, gender, birth_date FROM users ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.db.DB.Query(query, pageSize, offset)
	if err != nil {
		return models.PaginatedResponse{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Gender, &u.BirthDate); err != nil {
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
