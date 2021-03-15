package config

import (
	"os"
	"time"
)

var DbString string

func init() {
	time.Sleep(5 * time.Second)
	println("Hello bridges db config!")

	DbString = os.Getenv("MYSQL_USER") + ":" +
		os.Getenv("MYSQL_PASSWORD") + "@tcp(" +
		os.Getenv("MYSQL_HOST") + ")/" +
		os.Getenv("MYSQL_DATABASE") + "?parseTime=true"

}

//var DbString = "root:rootroot@tcp(db:3306)/av-check?parseTime=true"

