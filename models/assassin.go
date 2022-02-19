package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
)

/*
	Assassin: User 中嵌套一个 Student, students 表的外键和 users 表的主键同名，均为 username
	https://gorm.io/zh_CN/docs/has_one.html
*/

type AssassinUser struct {
	Username string           `gorm:"primaryKey;comment:用户名"`
	Password string           `gorm:"not null;comment:用户"`
	Student  *AssassinStudent `gorm:"foreignKey:Username;references:Username"`
	// 期望: foreign key students.username references users(username)
}

type AssassinStudent struct {
	StuId    string `gorm:"primaryKey;comment:学号"`
	Name     string `gorm:"not null;comment:姓名"`
	Username string `gorm:"not null;unique;comment:用户名"`
}

func TestAssassin(db *gorm.DB) bool {
	log.Println("Migrating model Assassin......")
	err := db.AutoMigrate(&AssassinUser{}, &AssassinStudent{})
	if err != nil {
		log.Println("Assassin migration failed: ", err)
		return false
	}
	query := db.Debug().Create(&AssassinUser{
		Username: "assassin",
		Password: "@55@551n",
		Student: &AssassinStudent{
			StuId:    "12345678",
			Name:     "刺客",
			Username: "assassin",
		},
	})
	if query.Error != nil {
		log.Println("Assassin create error: ", query.Error)
		return false
	}
	if query.RowsAffected == 0 {
		log.Println("Assassin did not finish create.")
		return false
	}
	var result *AssassinUser
	db.Debug().Where(&AssassinUser{Username: "assassin"}).Find(&result)
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
	log.Println("Assassin test finished.")
	return true
}
