package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// func checkErr(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

type pginfo struct {
	username   string
	department string
	created    string
}

var (
	db *sql.DB
)

const (
	insertSQL = "insert into userinfo(username,department,created) values($1,$2,$3) returning uid"
)

func sqlHelpForSync(str string, c chan int) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatal(r)
		}
	}()
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	println("call", str)
	//连接数据库
	for i := 0; i < 25000; i++ {
		if _, err := stmt.Exec(str, "asdf", time.Now()); err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	c <- 1
}

func main() {
	type Bindtwo struct {
		Schema string
		Port   int
	}
	type Info struct {
		Driver         string
		ConnectString  string
		Env            string
		Bind           []Bindtwo
		SessionTimeout int
		UpdateURL      string
		NotMinify      bool
		DefaultURL     string
		AllowCORS      bool
	}

	type Infomation struct {
		Infolice []Info
	}

	var coninfo Info

	bys, err := ioutil.ReadFile("F:/MyGo/src/pgconnect/connect.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal([]byte(bys), &coninfo); err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open(coninfo.Driver, coninfo.ConnectString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err = db.Exec("truncate table userinfo"); err != nil {
		log.Fatal(err)
	}
	startime := time.Now()

	ch := make(chan int, 4)

	for i := 0; i < 4; i++ {
		go sqlHelpForSync("qwer", ch)
	}
	for i := 0; i < 4; i++ {
		<-ch
	}
	//<-lexit
	fmt.Println("last", time.Since(startime).String())
}
