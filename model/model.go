package model

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"fmt"
	"time"
	"air/sys"
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

func UpdatePb(dateStart time.Time, dateFinish time.Time) error {
	db := new(database)

	err := db.connect()
	defer db.Close()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Что-то пошло не так")
	}

	stmt, err := db.Prepare(update)
	if err != nil {
		log.Println("[ERROR] tx.stmt.Exec: ", err)
		return fmt.Errorf("Что-то пошло не так", err)
	}
	defer stmt.Close()

	for i := 16; i < 18; i++ {
		_, err = stmt.Exec(dateStart,dateFinish ,i)
		if err != nil {
			log.Println("[ERROR] tx.stmt.Exec: ", err)
			return fmt.Errorf("Что-то пошло не так")
		}
	}
	return nil
}


func UpdateAll(dateStart time.Time, dateFinish time.Time) error {
	db := new(database)

	err := db.connect()
	defer db.Close()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Что-то пошло не так")
	}

	stmt, err := db.Prepare(update)
	if err != nil {
		log.Println("[ERROR] tx.stmt.Exec: ", err)
		return fmt.Errorf("Что-то пошло не так", err)
	}
	defer stmt.Close()

	for i := 1; i < 16; i++ {
		_, err = stmt.Exec(dateStart,dateFinish ,i)
		if err != nil {
			log.Println("[ERROR] tx.stmt.Exec: ", err)
			return fmt.Errorf("Что-то пошло не так")
		}
	}
	return nil
}

