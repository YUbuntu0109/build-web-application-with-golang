package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {

	// connect
	db, err := sql.Open("mysql", "root:GoodTime@/Go_Example?charset=utf8")
	checkErr(err)

	// insert data
	statement, err := db.Prepare("INSERT INTO userinfo SET username =?,department=?,created=?")
	checkErr(err)
	result, err := statement.Exec("GoogTech", "Google", "1998-01-09")
	checkErr(err)
	affect, err := result.RowsAffected()
	checkErr(err)
	fmt.Println("insert affect : ", affect)

	// select data
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	fmt.Println("select data from the table of userinfo successfully")
	for rows.Next() {
		var uid string
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid : ", uid)
		fmt.Println("username : ", username)
		fmt.Println("department :", department)
		fmt.Println("created : ", created)
	}

	// update data
	statement, err = db.Prepare("UPDATE userinfo SET username = ? WHERE uid = ?")
	checkErr(err)
	result, err = statement.Exec("GoogLead", 1)
	checkErr(err)
	affect, err = result.RowsAffected()
	checkErr(err)
	fmt.Println("update affect : ", affect)

	// delete data
	statement, err = db.Prepare("DELETE FROM userinfo WHERE uid = ?")
	checkErr(err)
	result, err = statement.Exec(2)
	checkErr(err)
	affect, err = result.RowsAffected()
	checkErr(err)
	fmt.Println("delete affect : ", affect)
	db.Close()
}

// check error
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
