package main

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Pginfo struct {
	Uid        int
	Username   string
	Department string
	Created    string
}

var count int = 0

func main() {
	//open postgresSql
	newtime := time.Now()
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=stbcon sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() //open and close sql

	res, err := db.Query("select*from userinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	fp, err := os.OpenFile("F:/MyGo/src/dboutinfile/json.txt", os.O_CREATE, 0755) // open your file
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	ffp, err := os.OpenFile("F:/MyGo/src/dboutinfile/json.gob", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer ffp.Close()

	for res.Next() {
		var outinfomation Pginfo //get one infomation
		if err = res.Scan(&outinfomation.Uid, &outinfomation.Username, &outinfomation.Department, &outinfomation.Created); err != nil {
			log.Fatal(err)
		}
		jsoninfo := json.NewEncoder(fp)
		jsoninfo.Encode(&outinfomation) //struct transform json into file

		code := gob.NewEncoder(ffp)
		if err = code.Encode(outinfomation); err != nil {
			log.Fatal(err)
		}

		count++
		fmt.Println("stophere", count)
	}
	fmt.Println("the time use", time.Since(newtime).String())
}
