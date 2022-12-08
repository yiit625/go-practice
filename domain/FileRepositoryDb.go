package domain

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lithammer/shortuuid"
	"strconv"
)

type FileRepositoryDb struct {
	client *sqlx.DB
}

func (f FileRepositoryDb) SaveImage(file File) (*File, *errs.AppError) {
	sqlInsert := "INSERT INTO file (id, name, path) values (?, ?, ?)"

	result, err := f.client.Exec(sqlInsert, shortuuid.New().String(), file.Name, file.Path)
	if err != nil {
		logger.Error("Error while creating new image: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new image: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	file.Id = strconv.FormatInt(id, 10)
	return &file, nil
}

func NewFileRepositoryDb(dbClient *sqlx.DB) FileRepositoryDb {
	return FileRepositoryDb{dbClient}
}
