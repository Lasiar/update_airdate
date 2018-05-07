package model

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"kre_air_update/sys"
	"time"
)

type database struct {
	*sql.DB
}

const update = "UPDATE ke_bak.dbo.Lab_AirPub set DateStart = ?, DateEnd = ?	 WHERE ID = ?"

func (d *database) connect() (err error) {
	d.DB, err = sql.Open("mssql", sys.GetConfig().ConnStr)
	if err != nil {
		return err
	}
	return nil
}

func Update(dateStart, dateFinish time.Time, begin, end int) error {
	db := new(database)
	err := db.connect()
	defer db.Close()

	if err != nil {
		return fmt.Errorf("db.connect: %v", err)
	}

	stmt, err := db.Prepare(update)
	if err != nil {
		return fmt.Errorf("db.Prepare: %v", err)
	}
	defer stmt.Close()

	for i := begin; i < end; i++ {
		_, err = stmt.Exec(dateStart, dateFinish, i)
		if err != nil {
			return fmt.Errorf("db.Exec: %v", err)
		}
	}
	return nil
}
