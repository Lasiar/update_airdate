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

func (d *database) connect() error {
	tmp, err := sql.Open("mssql", sys.GetConfig().ConnStr)
	d.DB = tmp
	if err != nil {
		return fmt.Errorf("[ERROR] Connect to database: ", err)
	}
	return nil
}

func Update(dateStart, dateFinish time.Time, begin, end int) error {
	db := new(database)
	err := db.connect()
	defer db.Close()

	if err != nil {
		return fmt.Errorf("[ERROR] db.connect: ", err)
	}

	stmt, err := db.Prepare(update)
	if err != nil {
		return fmt.Errorf("[ERROR] db.Prepare: ", err)
	}
	defer stmt.Close()

	for i := begin; i < end; i++ {
		_, err = stmt.Exec(dateStart, dateFinish, i)
		if err != nil {
			return fmt.Errorf("[ERROR] db.Exec: ", err)
		}
	}
	return nil
}
