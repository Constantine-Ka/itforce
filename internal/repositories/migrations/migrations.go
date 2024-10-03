package migrations

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type DB struct {
	*sqlx.DB
}

func CreateTable(db *sqlx.DB) (DB, error) {
	var dbOut DB
	for _, table := range []string{"acton_history", "dict_action", "users"} {
		filename := fmt.Sprintf("./migration/%s.up.sql", table)
		content, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Файл %s не найден \n", filename)
			return dbOut, err
		}
		_, err = db.Exec(string(content))
		if err != nil {
			fmt.Println(string(content))
			log.Println(err)

			return dbOut, err
		}
	}
	dbOut.DB = db
	return dbOut, nil
}
