package instance

import (
	"database/sql"
	"fmt"

	"github.com/rlapz/clean_arch_template/src/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	dbp := config.Db
	args := ""

	if len(config.Db.Args) > 0 {
		args = fmt.Sprint("?", config.Db.Args)
	}

	param := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", dbp.User, dbp.Password, dbp.Host,
		dbp.Port, dbp.Name, args)

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
