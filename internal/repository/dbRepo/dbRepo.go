package repository

import (
	"database/sql"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/repository"
)

type PostgresDbRepo struct {
	AppConfig *app.Config
	SqlDB     *sql.DB
}

func GetPostgresDbRepo(appConfig *app.Config, sqlDB *sql.DB) repository.DataBaseRepoInterface {
	return &PostgresDbRepo{
		AppConfig: appConfig,
		SqlDB:     sqlDB,
	}
}
