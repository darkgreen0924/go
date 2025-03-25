package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	//gorm.Model
	Id   int
	Name string
	Age  int
}

func (u User) TableName() string {
	return "user"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	user := User{Name: "lq"}
	db.Table("user").AutoMigrate(&user)
	//res := db.Table("user").Where(&user).First(&user)
	//select
	res := db.Find(&user)
	if res.Error == nil {
		log.Println(user)
	}

	users := []User{
		{
			Id:   1,
			Name: "lq",
			Age:  24,
		},
		{
			Id:   2,
			Name: "cyt",
			Age:  20,
		},
	}

	res = db.Create(&users)
	if res.Error == nil {
		log.Println("%s,插入数据数量：%d", time.Now(), res.RowsAffected)
	}

	u := User{Id: 3, Name: "lqqq"}
	res = db.Save(u)
	if res.Error == nil {
		log.Println("%s,更新数据数量：%d,%s", time.Now(), res.RowsAffected, u)
	}

	db.Delete(u)
	if res.Error == nil {
		log.Println("%s,更新数据数量：%d", time.Now(), res.RowsAffected)
	}

}
