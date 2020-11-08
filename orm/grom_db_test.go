package orm

import (
	"encoding/json"
	"github.com/chenleijava/go-guava/orm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"testing"
	"time"
)

//test grom starter
func TestGORMStart(t *testing.T) {
	dataSource := GetDataSource("root", "123456", "127.0.0.1:3306", "cwhd")
	var user = &model.User{}
	dataSource.Where("uid=?", "202001").First(user)
	log.Printf("nickName:%s", user.NickName)

	//raw sql ,scan
	var users []*model.User
	dataSource.Raw("select uid,bind_time from user order by bind_time desc limit 1 ").Scan(&users)
	show(&users)
	users = users[:0]

	// Use GORM API build SQL
	//https://gorm.io/docs/sql_builder.html
	var bindTime time.Time
	row := db.Table(user.TableName()).Select("bind_time").Where("uid=?", "202001").Row()
	err := row.Scan(&bindTime)
	if err != nil {

	}

	//Create Record
	_now := time.Now()
	_user := model.User{
		UId:         12,
		AvatarIcon:  "///",
		NickName:    "test",
		Sex:         0,
		Mobile:      "1738302",
		WxID:        "1738302",
		QqID:        "1738302",
		RoleType:    0,
		UserStatus:  0,
		RegTime:     _now,
		OUId:        0,
		BindTime:    _now,
		Birthday:    _now,
		OpenTime:    _now,
		ChannelName: "test",
		Balance:     0,
		PetCount:    0,
		DeviceID:    "1738302",
	}

	db.Delete(&_user)

	//result := dataSource.Create(&_user)
	//err = result.Error
	//if err != nil {
	//	log.Fatalf("insert err:%s", err.Error())
	//}
	//rowsAffected := result.RowsAffected
	//log.Printf("insert value ok , rows affected :%d", rowsAffected)

	//Batch Insert
	//To efficiently insert large number of records, pass a slice to the Create method. GORM will generate a single SQL statement to
	//insert all the data and backfill primary key values, hook methods will be invoked too.
	var tmpUsers []*model.User
	_user2 := _user //copy new value
	_user2.UId = 19
	//db.Delete(&_user2)
	tmpUsers = append(tmpUsers, &_user, &_user2)

	result := dataSource.Create(&tmpUsers) //pass the slice pointer
	err = result.Error
	if err != nil {
		log.Printf("insert err:%s", err.Error()) // err , which uid =19 inert err
		//batch insert err
		//try Upsert/ on conflict!
		result = db.Clauses(clause.OnConflict{
			//匿名字段
			Columns: []clause.Column{
				{
					Name: "uid",
				},
				{
					Name: "avatar_icon",
				},
			},
			DoUpdates: clause.AssignmentColumns([]string{"avatar_icon"}), // on conflict . update avatar_icon
		}).Create(&tmpUsers)
	}
	rowsAffected := result.RowsAffected
	log.Printf("insert value ok , rows affected :%d", rowsAffected)

	//query
	//https://gorm.io/docs/query.html

	//update
	//https://gorm.io/docs/update.html

	//gorm delete
	//https://gorm.io/docs/delete.html

	//base query
	dataSource.Select("uid", "bind_time").Order("bind_time desc").Limit(1).Find(&users)
	log.Printf("........")
	show(&users)

}

func show(users *[]*model.User) {
	if users != nil {
		for _, user := range *users {
			marshalData, _ := json.Marshal(user)
			log.Printf("%s", string(marshalData))
		}
	}
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func sqliteQuick() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	_ = db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, "id=1")            // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
