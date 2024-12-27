package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// gorm 默认会将表名添加复数: User->users, 使用TableName()方法设置表名

/*
	gorm会将字段转变为蛇形，比如UserName -> user_name
	Keyword->keyword
*/

type User struct {
	Id      int    // Id默认被gorm认为是主键  `gorm:"column:aid,primaryKey"`  数据库里面字段名是aid,是主键
	Keyword string `gorm:"column:keywords"`
	City    string // city
	Uid     int    `gorm:"column:uid"`
	Degree  string
	Gender  string
}

func (User) TableName() string {
	return "user"
}

func read(client *gorm.DB, city string) *User {
	var user User
	// client.Select("id,city").Where("city=?", city).Find(&user)
	//if len(user) > 0 {
	//	return &user[0]
	//}
	//return nil
	user.Id = 1707919 // 隐含的where条件
	err := client.Select("id,city").Where("city=?", city).Limit(1).Take(&user).Error
	checkError(err)
	return &user

}

func main() {
	dataSourceName := "root:ll1995ll@tcp(192.168.164.132:3306)/test?charset=utf8&parseTime=True"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkError(err)
	// 查询
	//user := read(client, "上海")
	//if user != nil {
	//	fmt.Printf("%v\n", *user)
	//}

	// 增
	//user := User{
	//	Id:      5858,
	//	Keyword: "golang",
	//	City:    "天津",
	//	Uid:     5858,
	//	Degree:  "de",
	//	Gender:  "f",
	//}
	//client.Create(user)

	// 改
	//client.Model(User{}).Where("id=?", 5858).Update("city", "郑州")
	//client.Model(User{}).Where("id=?", 5858).Updates(
	//	map[string]interface{}{"city": "郑州", "gender": "男"})

	// 删
	client.Where("id=?", 5858).Delete(&User{})
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
