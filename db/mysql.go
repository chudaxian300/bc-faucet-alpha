package db

import (
	"api/conf"
	"api/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init()  {
	var err error
	var dbConfig = conf.Conf.Db
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Db,
		dbConfig.Charset)

	Db,err = sql.Open(dbConfig.Dialects,url)

	if err != nil {
		fmt.Println(url)
		panic(err)
	}

	Db.SetMaxOpenConns(dbConfig.MaxOpen)
	Db.SetMaxIdleConns(dbConfig.MaxIdle)

	logger := utils.Log()
	logger.Info("mysql connect success")
}

