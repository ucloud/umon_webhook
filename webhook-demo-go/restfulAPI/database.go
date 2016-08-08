package restfulAPI

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"utils"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	singleMutex sync.Mutex
)

type MysqlDsn struct {
	User   string
	Passwd string
	Host   string
	Port   uint16
	DbName string
}

const DSN_Format = "%s:%s@tcp(%s:%d)/%s?charset=utf8"

func initOptions() *MysqlDsn {
	cfg := utils.GetGlobalConf()
	mysqldsn := &MysqlDsn{
		Host: "localhost",
		Port: 3306,
	}

	if cfg != nil {
		mysqldsn.User = cfg.GetString("mysql-user")
		mysqldsn.Passwd = cfg.GetString("mysql-passwd")
		mysqldsn.DbName = cfg.GetString("mysql-db")
		mysqldsn.Host = cfg.GetString("mysql-host")

		port := cfg.GetInt("mysql-port")
		if port > 0 {
			mysqldsn.Port = uint16(port)
		}
	}

	return mysqldsn
}

func NewMySQLDB() *sql.DB {
	opt := initOptions()
	dsn := fmt.Sprintf(DSN_Format, opt.User, opt.Passwd, opt.Host, opt.Port, opt.DbName)
	newdb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("With DSN : ", dsn, ", Open mysql failed : ", err.Error())
		newdb.Close()
		return nil
	}

	err = newdb.Ping()
	if err != nil {
		log.Println("With DSN : ", dsn, ", Ping mysql failed : ", err.Error())
		newdb.Close()
		return nil
	}

	return newdb
}

func GetDB() *sql.DB {
	singleMutex.Lock()
	defer singleMutex.Unlock()
	if db == nil {
		db = NewMySQLDB()
	}

	return db
}
