package users

import (
	models "dariomatias-dev/snap_code/cmd/database/models/user"
	"dariomatias-dev/snap_code/cmd/utils"
	"database/sql"
	"log"
)

func NewUsers(dbcon *sql.DB) *UsersQueries {
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

	_, err := uq.dbcon.Exec(query, createUser.UserName)
	if err != nil {
		log.Fatalln(err)
	}
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
		var userName string

		err := response.Scan(&userName)
		if err != nil {
			log.Fatalln(err)
		}

		user := models.UserModel{
			UserName: userName,
		}

		users = append(users, user)
	}

	return users
}

func (uq *UsersQueries) UpdateOneByUserName(
	userName string,
	updateUser models.UpdateUserModel,
) {
	queryPath := "cmd/database/queries/users/usersQueries/updateOneByUserNameQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := uq.dbcon.Exec(query, userName, updateUser.UserName)
	if err != nil {
		log.Fatalln(err)
	}
}

func (uq *UsersQueries) DeleteOneByUserName(
	userName string,
) {
	queryPath := "cmd/database/queries/users/usersQueries/deleteOneByUserNameQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := uq.dbcon.Exec(query, userName)
	if err != nil {
		log.Fatalln(err)
	}
}
