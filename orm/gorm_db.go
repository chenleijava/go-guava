package orm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math"
	"runtime"
	"sync"
	"time"
)

var db *gorm.DB
var dbOnce sync.Once

//get gorm datasource
//username:
//password:
//host:
//dbName:
func GetDataSource(userName, password, host, dbName string) *gorm.DB {
	if db == nil {
		dbOnce.Do(func() {
			//https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
			//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
			dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName,
				password, host, dbName)
			_db, err := gorm.Open(mysql.New(
				mysql.Config{
					DSN:                       dsn,   // data source name
					DefaultStringSize:         256,   // default size for string fields
					DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
					DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
					DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
					SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
				}), &gorm.Config{})
			if err != nil {
				log.Fatalf("init gorm db err:%s", err.Error())
			}
			db = _db
			sqlDb, err := db.DB()
			if sqlDb != nil {
				//https://gorm.io/docs/connecting_to_the_database.html
				//https://github.com/brettwooldridge/HikariCP
				maxOpenConns := runtime.NumCPU()*2 + 1
				maxOpenConns = int(math.Max(float64(maxOpenConns), 30))
				//
				sqlDb.SetMaxIdleConns(maxOpenConns / 2)    //>=15
				sqlDb.SetMaxOpenConns(maxOpenConns)        // >=30
				sqlDb.SetConnMaxIdleTime(time.Minute * 10) // idle num less than max open cons 10 minutes
				sqlDb.SetConnMaxLifetime(time.Minute * 30) // 30 minutes  Expired connections may be closed lazily before reuse.
			}
		})
	}
	return db
}
