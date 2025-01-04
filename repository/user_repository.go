package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) *UserRepository {
	return &UserRepository{
		connection: connection,
	}
}

func (pr *UserRepository) CreateUser(user model.User) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO users" +
		"(name, email)" +
		" VALUES ($1, $2) RETURNING id_user")

	if err != nil {
		fmt.Println(err)
		fmt.Print(query)
		return 0, err
	}

	err = query.QueryRow(user.Name, user.Email).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil

}
