package fastsql

import (
	"github.com/chenleijava/go-guava"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestDB_Close(t *testing.T) {
	dbh, err := Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dsp_db_setting?charset=utf8&loc=Asia%2FShanghai&parseTime=true", 2)
	if err != nil {
		log.Fatalln(err)
	}
	var j = 1
	//batch insert test
	begin := guava.GetTimeStampTimeMillis()
	query := "INSERT INTO dsp_db_setting.app_info (app_name, app_sigin, app_packagename, app_size, app_status, app_version_name, app_version_number, app_update_role_name,app_update_time) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	for j <= 3 {
		dbh.BatchInsert(query, "aw123!@#:"+strconv.Itoa(j), "xxas", "com.xx.google.co", 12, 0, "1.0", 1, "石头哥", time.Now().UTC())
		j++
	}
	dbh.LastBatchInsert(query)
	//Mock seconds call ,the stmtQuery not be del
	dbh.BatchInsert(query, "aw123!@#:"+strconv.Itoa(j), "xxas", "com.xx.google.co", 12, 0, "1.0", 1, "石头哥", time.Now().UTC())

	//last xx
	dbh.LastBatchInsert(query)

	log.Printf("batch insert cost time:%d ms", guava.GetTimeStampTimeMillis()-begin)
}
