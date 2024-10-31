package solutions

import (
	"dariomatias-dev/snap_code/cmd/database/models/solution"
	"dariomatias-dev/snap_code/cmd/utils"
	"database/sql"
	"fmt"
	"log"
)

func NewSolutionsQueries(dbcon *sql.DB) *SolutionsQueries {
	return &SolutionsQueries{dbcon: dbcon}
}

type SolutionsQueries struct {
	dbcon *sql.DB
}

func (sq SolutionsQueries) Create(
	createSolution solution.SolutionModel,
) error {
	queryPath := "cmd/database/queries/solutions/queries/createQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := sq.dbcon.Exec(query, createSolution.Key, createSolution.FileName)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (sq SolutionsQueries) GetByKey(
	solutionKey string,
) *solution.SolutionModel {
	queryPath := "cmd/database/queries/solutions/queries/getByKeyQuery.sql"
	query := utils.ReadFile(queryPath)

	var key string
	var fileName string

	rows, err := sq.dbcon.Query(query, solutionKey)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	rows.Scan(&key, &fileName)

	return &solution.SolutionModel{
		Key:      key,
		FileName: fileName,
	}
}

func (sq SolutionsQueries) GetAll() []solution.SolutionModel {
	queryPath := "cmd/database/queries/solutions/queries/getAllQuery.sql"
	query := utils.ReadFile(queryPath)

	rows, err := sq.dbcon.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	solutions := []solution.SolutionModel{}

	for rows.Next() {
		var key string
		var fileName string

		rows.Scan(&key, &fileName)

		solution := solution.SolutionModel{
			Key:      key,
			FileName: fileName,
		}

		solutions = append(solutions, solution)
	}

	return solutions
}

func (sq SolutionsQueries) UpdateByKey(
	key string,
	updateSolution solution.UpdateSolutionModel,
) error {
	queryPath := "cmd/database/queries/solutions/queries/updateByKeyQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := sq.dbcon.Exec(query, updateSolution.Key, updateSolution.FileName, key)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	return nil
}

func (sq SolutionsQueries) DeleteByKey(
	key string,
) error {
	queryPath := "cmd/database/queries/solutions/queries/deleteByKeyQuery.sql"
	query := utils.ReadFile(queryPath)

	_, err := sq.dbcon.Exec(query, key)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
