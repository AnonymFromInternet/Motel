package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DataBaseConnectionPool struct {
	SQL *sql.DB
}

const maxOpenDBConnections = 15
const maxIdleDBConnections = 10
const maxDBConnectionLifeTime = 8 * time.Minute

func GetDataBaseConnectionPool(dataSourceName string) (*DataBaseConnectionPool, error) {
	sqlDBConnPool, err := getSqlDBConnPool(dataSourceName)
	if err != nil {
		panic(err)
	}

	sqlDBConnPool.SetMaxOpenConns(maxOpenDBConnections)
	sqlDBConnPool.SetConnMaxIdleTime(maxIdleDBConnections)
	sqlDBConnPool.SetConnMaxLifetime(maxDBConnectionLifeTime)

	var dataBaseConnectionPool DataBaseConnectionPool
	dataBaseConnectionPool.SQL = sqlDBConnPool

	err = testDB(sqlDBConnPool)
	if err != nil {
		return nil, err
	}

	return &dataBaseConnectionPool, nil

}

func testDB(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

func getSqlDBConnPool(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		return nil, err
	}

	fmt.Println("db :", db)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("cannot close a database connection")
		}
	}(db)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
