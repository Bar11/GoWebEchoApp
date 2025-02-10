package models

import (
	"GoWebEchoApp/config/setting"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type UserInfo struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func InitDB(e *echo.Echo) {
	// 设置数据库连接信息
	dsn := setting.Settings.Username + ":" + setting.Settings.Password + "@tcp(" + setting.Settings.Host + ":" + setting.Settings.Port + ")/" + setting.Settings.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	// 替换username、password、dbname为你的MySQL用户名、密码和数据库名

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// 验证数据库连接是否成功
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}

	//// 执行查询
	//rows, err := db.Query("SELECT id, name FROM users")
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	//defer rows.Close()
	//
	//// 处理查询结果
	//for rows.Next() {
	//	var id int
	//	var name string
	//	if err := rows.Scan(&id, &name); err != nil {
	//		e.Logger.Fatal(err)
	//	}
	//	fmt.Printf("ID: %d, Name: %s\n", id, name)
	//}
	//
	//// 检查是否出现查询错误
	//if err := rows.Err(); err != nil {
	//	e.Logger.Fatal(err)
	//}
	//
	//// 插入数据
	//stmt, err := db.Prepare("INSERT INTO users (name, age) VALUES (?, ?)")
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	//defer stmt.Close()
	//
	//result, err := stmt.Exec("John Doe", 30)
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	//
	//// 获取插入的ID
	//lastInsertId, err := result.LastInsertId()
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	//fmt.Printf("Last Insert ID: %d\n", lastInsertId)
	//
	//// 受影响的行数
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	e.Logger.Fatal(err)
	//}
	//fmt.Printf("Rows Affected: %d\n", rowsAffected)
}
