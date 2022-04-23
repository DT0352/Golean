package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, error := sql.Open("mysql", "root:123456@tcp(120.53.241.149:3306)/gotest?charset=utf8")
	checkErr(error)
	fmt.Println(db)
	stmt, error := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(error)
	res, error := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(error)
	id, error := res.LastInsertId()
	checkErr(error)
	fmt.Println(id)
	stmt, error = db.Prepare("update userinfo set username=? where uid=? ")
	checkErr(error)
	res, error = stmt.Exec("astaxieupdate", id)
	affect, error := res.RowsAffected()
	checkErr(error)
	fmt.Println(affect)
	rows, error := db.Query("select * from userinfo")
	checkErr(error)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	stmt, error = db.Prepare("delete from userinfo")
	checkErr(error)
	res, error = stmt.Exec()
	affect, error = res.RowsAffected()
	checkErr(error)
	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
