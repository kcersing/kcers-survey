package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 测试Add函数
func main() {
	fmt.Println("开始")
	dbConn, err := sql.Open("mysql", "read_only:XFL_666-ncs-cnzd-zcygss-ssjss@tcp(120.0.0.1:3306)/kcers?charset=utf8")

	if err != nil {
		fmt.Println(err)
	}

	type User struct {
		Name   string `json:"name,omitempty"`
		Mobile string `json:"mobile,omitempty"`
	}
	//type Details struct {
	//	Sex      sql.NullInt64  `json:"sex,omitempty"`
	//	Birthday sql.NullString `json:"birthday,omitempty"`
	//	IdCard   sql.NullString `json:"id_Card,omitempty"`
	//}

	query, err := dbConn.Query("SELECT name,mobile FROM `kcers`.`saas_merchant_customer` WHERE `merchant_id` = '6' LIMIT 0,1000")
	if err != nil {
		fmt.Println(err)
	}
	var u User
	for query.Next() {
		err := query.Scan(&u.Mobile, &u.Name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("u: %v\n", u)

		//r, err := dbConn.Exec("insert into user_tbl (name,mobile) values(?,?)", u.Name, u.Mobile)
		//if err != nil {
		//	fmt.Printf("err: %v\n", err)
		//} else {
		//	i, _ := r.LastInsertId()
		//	fmt.Printf("i: %v\n", i)
		//}

	}

	fmt.Println("结束")
}
