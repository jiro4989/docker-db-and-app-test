package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		user   = os.Getenv("DB_USERNAME")
		pass   = os.Getenv("DB_PASSWORD")
		host   = os.Getenv("DB_HOSTNAME")
		port   = os.Getenv("DB_PORT")
		dbname = os.Getenv("DB_DATABASE")
	)

	// dns user:password@tcp(host:port)/dbname
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	fmt.Println(dns)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// 今回はダミーでuserというテーブルを作成したので
	rows, err := db.Query("select * from users;")
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}
