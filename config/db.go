package config

import(
"fmt"
"database/sql"
_ "github.com/go-sql-driver/mysql"
"time"

)

const(
	DB_HOST="localhost"
	DB_USER="root"
	DB_PASS="pass"
	DB_NAME="blog"
	DB_PORT="3306"
)
var DB *sql.DB
func InitDB() {
	DB, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
if err!=nil{
fmt.Println("error connecting database",err.Error())
}
	DB.SetConnMaxLifetime(2*time.Second)
	DB.SetMaxOpenConns(80)
	DB.SetMaxIdleConns(20)
defer DB.Close()
}