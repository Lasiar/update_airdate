package model

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"kre_air_update/sys"
	"time"
)



const update = "UPDATE dbo.Lab_AirPub set DateStart = ?, DateEnd = ? WHERE ID = ?"
const getTime = "select  DateStart, DateEnd  from krasecology.dbo.Lab_AirPub  WHERE ID = 1 or ID = 17"

type database struct {
	*sql.DB
}

func (d *database) connect() (err error) {
	d.DB, err = sql.Open("mssql", sys.GetConfig().ConnStr)
	if err != nil {
		return err
	}
	return nil
}

func Update(dateStart, dateFinish time.Time, begin, end int) error {
	db := new(database)

	if err := db.connect(); err != nil {
		return fmt.Errorf("[db.connect] %v", err)
	}

	defer db.Close()

	stmt, err := db.Prepare(update)
	if err != nil {
		return fmt.Errorf("[db.Prepare] %v", err)
	}

	defer stmt.Close()

	for i := begin; i < end; i++ {
		_, err = stmt.Exec(dateStart, dateFinish, i)
		if err != nil {
			return fmt.Errorf("[db.Exec] %v", err)
		}
	}
	return nil
}

func Select() ([4]string, error) {
	var dates [4]string

	db := new(database)

	if err := db.connect(); err != nil {
		return dates, fmt.Errorf("[db.connect] %v", err)
	}

	defer db.Close()

	row, err := db.Query(getTime)
	if err != nil {
		return dates, fmt.Errorf("[db.Query] %v", err)
	}

	for i := 0; row.Next(); i+=2 {
		err := row.Scan(&dates[i], &dates[i+1])
		if err != nil {
			return dates, fmt.Errorf("[row Scan] %v", err)
		}
		fmt.Println(i)
	}

	for i, date := range dates {
		if len(date) < 10 {
			return dates, fmt.Errorf("[string processing] %v", "Из базы пришла пустая строка")
		}
		dates[i] = date[:10]
	}

	return dates, nil
}
