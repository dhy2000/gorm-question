package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

/*
	Caster: 与 Assassin 类似的 has one 联系, 但是 students 表的外键字段与 users 的主键字段不相同
*/

type CasterUser struct {
	Username string         `gorm:"primaryKey;comment:用户名"`
	Password string         `gorm:"not null;comment:密码"`
	Student  *CasterStudent `gorm:"foreignKey:AuthName;references:Username"`
	// 期望: foreign key students.auth_name references users(username)
}

type CasterStudent struct {
	StuId    string `gorm:"primaryKey;comment:学号"`
	Name     string `gorm:"not null;comment:姓名"`
	AuthName string `gorm:"not null;unique;comment:用户名"`
}

func TestCaster(db *gorm.DB) bool {
	log.Println("Migrating model Caster......")
	err := db.AutoMigrate(&CasterUser{}, &CasterStudent{})
	if err != nil {
		log.Println("Caster migration failed: ", err)
		return false
	}
	query := db.Debug().Create(&CasterUser{
		Username: "caster",
		Password: "c@5tEr",
		Student: &CasterStudent{
			StuId:    "12345678",
			Name:     "魔术师",
			AuthName: "caster",
		},
	})
	if query.Error != nil {
		log.Println("Caster create error: ", query.Error)
		return false
	}
	if query.RowsAffected == 0 {
		log.Println("Caster did not finish create.")
		return false
	}
	var result *CasterUser
	db.Debug().Where(&CasterUser{Username: "caster"}).Find(&result)
	js, _ := json.Marshal(result)
	log.Println(string(js))
	log.Println("Caster test finished.")
	return true
}
