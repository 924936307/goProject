package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//mysql的示例
//参考项 ：https://blog.csdn.net/naiwenw/article/details/79281220
type user struct {
	id   int64
	name string
	age  int
}

//数据库配置
const (
	userName = "root"
	password = "lNsk8mIXTOrpmJ8n"
	ip       = "127.0.0.1"
	port     = "3306"
	DBName   = "go_test"
)

//DB数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", DBName, "?charset=utf8"}, "")
	fmt.Println("path :", path)
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func QueryRow(id int64) {
	var u user
	queryStr := "select id,name,age from user where id = ?"
	err := DB.QueryRow(queryStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v \n", err)
	}
	fmt.Printf("id:%d, name:%s, age:%d \n", u.id, u.name, u.age)
}

func QueryRows() {
	sqlStr := "select id,name,age from user where id > ?"
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed ,err:%v \n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v \n", err)
			return
		}
		fmt.Printf("id:%d,name:%s,age:%d \n", u.id, u.name, u.age)
	}
}

func InsertRow(name string, age int) {
	sqlStr := "insert into user(name,age) values(?,?)"
	exec, err := DB.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed,err:%v \n", err)
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsert id failed,err::%v \n", err)
		return
	}
	fmt.Printf("insert success,the id is :%d \n", id)
}

func UpdateRow(id int64, age int) {
	sqlStr := "update user set age = ? where id = ?"
	exec, err := DB.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed,err:%v \n", err)
		return
	}
	affected, err := exec.RowsAffected() //影响的行数
	if err != nil {
		fmt.Printf("get the rowsAffected failed,err:%v \n", err)
		return
	}
	fmt.Printf("update success,affected rows :%d \n", affected)
}

func DeleteRow(id int64) {
	sqlStr := "delete from user where id = ?"
	exec, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v \n", err)
		return
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffectedRows failed,err:%v \n", err)
		return
	}
	fmt.Printf("delete success,affected rows :%d \n", affected)
}

func main() {
	InitDB()
	InsertRow("bobo", 18)
	InsertRow("dog", 1)
	InsertRow("cat", 2)
	InsertRow("orange", 4)
	InsertRow("bnnana", 6)
	InsertRow("sheep", 9)
	fmt.Println("---------------查询id =1 的行------------------")
	QueryRow(1)
	fmt.Println("--------------查询所有的记录----------------")
	QueryRows()
	fmt.Println("-------------修改id = 1记录--------------------------")
	UpdateRow(1, 20)
	fmt.Println("-----------------刪除id =1的记录----------------------")
	DeleteRow(1)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
