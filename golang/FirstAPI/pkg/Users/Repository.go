package Users

import (
	"database/sql"
)

type Repository struct {
	Context *sql.DB
}

func GetRepository(context *sql.DB) *Repository {
	return &Repository{
		Context: context,
    }
}

func (repository Repository) GetAll() *[]UserModel {
	var users []UserModel

	rows, err := repository.Context.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user UserModel
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return &users
}