package main;

import (
	"database/sql"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
);

const (
	userName = "root"
	password = "123456"
	ip = "127.0.0.1"
	port = "3306"
	dbName = "song"
);

var DB * sql.DB;

func main()  {
	var connect bool = InitDb();
	if(connect == true){
		println("success");
		println(InsertUser());
	}else{
		println("error");
	}
}

func InitDb() bool {
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "");
	DB, _ = sql.Open("mysql", path);
	DB.SetConnMaxLifetime(100);
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil{
		return false;
	}else{
		return true;
	}
}


func InsertUser()  (int64) {
	tx, err := DB.Begin()
	if(err != nil){
		fmt.Println("tx fail");
		return 0;
	}

	stmt, err := tx.Prepare(`INSERT s_user (u_name,u_tel) values (?,?)`);
	if(err != nil){
		fmt.Println("stmt fail");
		return 0;
	}

	re, err := stmt.Exec("略略略", "13933863958");
	if(err != nil){
		fmt.Println("re fail");
		return 0;
	}

	last_id, err := re.LastInsertId()
	if(err != nil){
		fmt.Println("last_id fail");
		return 0;
	}

	return last_id;
}