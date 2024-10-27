package users

import (
	models "dariomatias-dev/snap_code/cmd/database/models/user"
	"dariomatias-dev/snap_code/cmd/utils"
	"database/sql"
	"log"
)

func NewUsersQueries(dbcon *sql.DB) *UsersQueries {
	return &UsersQueries{dbcon: dbcon}
}

type UsersQueries struct {
	dbcon *sql.DB
}

func (uq *UsersQueries) Create(
	createUser models.CreateUserModel,
) {
	queryPath := "cmd/database/queries/users/usersQueries/createQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := uq.dbcon.Exec(query, createUser.Username)
	if err != nil {
		log.Fatalln(err)
	}
}

func (uq *UsersQueries) Count() int {
	queryPath := "cmd/database/queries/users/usersQueries/countQuery.sql"
	query := utils.ReadFile(queryPath)

	var count int

	err := uq.dbcon.QueryRow(query).Scan(&count)
	if err != nil {
		log.Fatalln(err)
	}

	return count
}

func (uq *UsersQueries) GetAll() []models.UserModel {
	queryPath := "cmd/database/queries/users/usersQueries/getAllQuery.sql"
	query := utils.ReadFile(queryPath)

	response, err := uq.dbcon.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Close()

	users := []models.UserModel{}

	for response.Next() {
		var username string

		err := response.Scan(&username)
		if err != nil {
			log.Fatalln(err)
		}

		user := models.UserModel{
			Username: username,
		}

		users = append(users, user)
	}

	return users
}

func (uq *UsersQueries) UpdateByUsername(
	username string,
	updateUser models.UpdateUserModel,
) {
	queryPath := "cmd/database/queries/users/usersQueries/updateByUsernameQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := uq.dbcon.Exec(query, updateUser.Username, username)
	if err != nil {
		log.Fatalln(err)
	}
}

func (uq *UsersQueries) DeleteByUsername(
	username string,
) {
	queryPath := "cmd/database/queries/users/usersQueries/deleteByUsernameQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := uq.dbcon.Exec(query, username)
	if err != nil {
		log.Fatalln(err)
	}
}
