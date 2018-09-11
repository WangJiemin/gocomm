package gocomm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type DBConfig struct {
	IpEndPoint  string
	UserName    string
	Pwd         string
	MaxConn     int
	MaxIdleConn int
}

func DBConnect(config DBConfig,dbName string) *sql.DB{
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.UserName, config.Pwd, config.IpEndPoint,dbName)
	fmt.Printf("conn:%s\n", connStr)
	db, _ := sql.Open("mysql", connStr)
	fmt.Printf("conn:%s%v\n", "opened",db)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetMaxOpenConns(config.MaxConn)
	db.Ping()
	return db
}

func DBQuery(db *sql.DB,sqlString string, args ...interface{}) {



	fmt.Printf("query db:%s\n",sqlString)
	rows, _ := db.Query(sqlString)

	defer rows.Close()
	//	quick.CheckError{err}
	columns, _ := rows.Columns()
	fmt.Printf("select column:%v\n", columns)
}

func DBUpdate(db *sql.DB,sqlString string,  args ...interface{}) int32{
	return 0
}

func DBDelete(db *sql.DB,sqlString string,  args ...interface{}) int32 {
	return 0
}

func DBInsert(db *sql.DB,sqlString string,  args ...interface{}) int32{
	return 0
}





