package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/inihikam/web-sti-api/config"
)  

var DB *sql.DB  

func InitDB() error {  
    cfg := config.ConnectDB()  
    dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"  
    var err error  
    DB, err = sql.Open("mysql", dsn)  
    if err != nil {  
        return err  
    }  

    err = DB.Ping()  
    if err != nil {  
        return err  
    }  

    fmt.Println("Database connection established")  
    return nil  
}