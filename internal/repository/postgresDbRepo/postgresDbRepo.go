package repository

import (
	"database/sql"
	"github.com/AnonymFromInternet/Motel/internal/app"
	"github.com/AnonymFromInternet/Motel/internal/repository"
)

type postgresDBRepo struct {
	AppConfig *app.Config
	SqlDB     *sql.DB
}

func GetPostgresDBRepo(appConfig *app.Config, sqlDB *sql.DB) repository.DataBaseRepoInterface {
	return &postgresDBRepo{
		AppConfig: appConfig,
		SqlDB:     sqlDB,
	}
}
