package basic

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
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

	var code int64 = 5
	actual := insert()
	if actual != code {
		t.Errorf("Insert Failed!")
	}
}
