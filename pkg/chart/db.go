/*
Package chart is
*/
package chart

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	db *sql.DB
}

func NewConnect() Db {
	db := Db{}
	con, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/liquid")
	log.Println("Connected to mysql.")

	if err != nil {
		log.Fatal(err)
	}
	db.db = con
	return db
}

func (d *Db) Fetch() (ohlcs OHLCs) {
	rows, err := d.db.Query("SELECT * FROM minutes")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var ohlc OHLC
		err := rows.Scan(&ohlc.Timestamp, &ohlc.Open, &ohlc.High, &ohlc.Low, &ohlc.Close, &ohlc.Range)
		ohlc.Range = 60

		if err != nil {
			panic(err.Error())
		}
		ohlcs = append(ohlcs, ohlc)
	}
	return
}

func (d *Db) Insert(ohlc OHLC) (err error) {
	ins, err := d.db.Prepare("INSERT INTO minutes(timestamp,open,high,low,close,`range`,vbuy,vsell) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(ohlc.Timestamp, ohlc.Open, ohlc.High, ohlc.Low, ohlc.Close, ohlc.Range, ohlc.Vbuy, ohlc.Vsell)

	return nil
}

func (d *Db) Close() {
	d.db.Close()
}
