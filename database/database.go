package database

import (
    cfg "../config"
    log "../logger"

    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Setup() {
    log.InfoLogger.Println("Initializing database connection...")

    db, err := sql.Open("mysql", 
                        cfg.DatabaseSetting.User + ":" +
                        cfg.DatabaseSetting.Password + "@" +
                        "tcp(" + cfg.DatabaseSetting.Host + ")" +
                        "/" + cfg.DatabaseSetting.Name) 
                        // "username:password@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    } else {
        log.InfoLogger.Println("Connection to mysql established!")
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

}
