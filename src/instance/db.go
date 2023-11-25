package instance

import (
	"database/sql"
	"fmt"

	"github.com/rlapz/clean_arch_template/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	dbp := config.Db
	param := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbp.User, dbp.Password, dbp.Host, dbp.Port, dbp.Name)

	db, err := sql.Open("mysql", param)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(dbp.ConnectionLifetimeMax)
	db.SetMaxOpenConns(dbp.ConnectionsOpenMax)
	db.SetMaxIdleConns(dbp.ConnectionsIdleMax)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
