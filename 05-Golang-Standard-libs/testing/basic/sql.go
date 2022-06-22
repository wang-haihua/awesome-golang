package basic

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
func main() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	defer db.Close() // defer 关键字用于资源释放
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(0)  // 最大空闲连接数
	db.SetMaxOpenConns(30) // 最大打开连接数
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connected mysql Successfully")

	//insert()
	//update()
	//delete()
	query()
}
*/

func insert() int64 {
	tx, err := db.Begin() // 开启一个事务
	if err != nil {
		return -1
	}
	defer tx.Rollback() // 开启事务成功 终止回滚

	var sql_stat string = "INSERT INTO user (username,sex,email) VALUES (?,?,?)"
	stmt, err := tx.Prepare(sql_stat) // 准备SQL语句
	if err != nil {
		return -1
	}

	res, err := stmt.Exec("Rainy", "female", "rainy@tencent.com") // 执行不返回查询结果的SQL语句
	if err != nil {
		return -1
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		return -1
	}

	id, err := res.LastInsertId() // 返回插入成功后数据生成的自增id
	fmt.Println("Inserted row_id: ", id)
	defer stmt.Close() // 关闭 statement 并释放资源

	return id
}

func update() {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	var sql_stat string = "UPDATE user SET username=? where user_id=?"
	stmt, err := tx.Prepare(sql_stat)
	if err != nil {
		return
	}

	_, err = stmt.Exec("Zenowang", 1) // 声明新的变量才使用 :=
	if err != nil {
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}

	fmt.Println("Update Successfully")

	defer stmt.Close()
	return
}

func delete() {
	tx, e := db.Begin()
	if e != nil {
		return
	}
	defer tx.Rollback()

	var sql_stat string = "DELETE FROM user where user_id=?"
	stmt, e := tx.Prepare(sql_stat)
	if e != nil {
		return
	}

	_, e = stmt.Exec(3)
	if e != nil {
		return
	}
	e = tx.Commit()
	if e != nil {
		return
	}

	fmt.Println("Delete Successfully")

	defer stmt.Close()
	return
}

func query() {
	var sql_stat string = "SELECT * FROM user WHERE sex='male'"
	rows, err := db.Query(sql_stat) // 执行 Select 等查询语句 返回查询记录
	defer rows.Close()              // 关闭 rows 防止重复枚举
	if err != nil {
		return
	}

	fmt.Println("id name sex email")
	for rows.Next() { // next 迭代器遍历记录
		var id int
		var name, sex, email string
		rows.Scan(&id, &name, &sex, &email) // 扫描将当前行记录中的属性值复制到DEST指向的值中
		fmt.Println(id, name, sex, email)
	}

	if err := rows.Err(); err != nil {
		return
	}
	return
}

type User struct {
	Id    int
	Name  string
	sex   string
	email string
}

func queryByUserId(userId int) *User {

	row := db.QueryRow("SELECT * FROM user WHERE user_id=?", userId) // 执行 Select 等查询语句 返回查询记录

	if err := row.Err(); err != nil {
		return nil
	}

	var id int
	var name, sex, email string
	row.Scan(&id, &name, &sex, &email) // 扫描将当前行记录中的属性值复制到DEST指向的值中

	return &User{id, name, sex, email}
}
