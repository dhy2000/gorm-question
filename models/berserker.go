package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
)

/*
	Berserker: 与 Archer 类似的 belongs to 联系, 但是 students 表的外键字段与 users 的主键字段不相同
*/

type BerserkerUser struct {
	Username string `gorm:"primaryKey;comment:用户名(唯一)"`
	Password string `gorm:"not null;comment:用户密码"`
}

type BerserkerStudent struct {
	StuId    string         `gorm:"primaryKey;comment:学号"`
	Name     string         `gorm:"not null;comment:姓名"`
	AuthName string         `gorm:"not null;unique;comment:用户名"`
	User     *BerserkerUser `gorm:"foreignKey:AuthName;references:Username"`
	// 期望: foreign key students.auth_name references users(username)
}

func TestBerserker(db *gorm.DB) bool {
	log.Println("Migrating model Berserker......")
	err := db.AutoMigrate(&BerserkerUser{}, &BerserkerStudent{})
	if err != nil {
		log.Println("Berserker migration failed: ", err)
		return false
	}
	query := db.Debug().Create(&BerserkerStudent{
		StuId:    "12345678",
		Name:     "狂战士",
		AuthName: "berserker",
		User: &BerserkerUser{
			Username: "berserker",
			Password: "bEr5ErkEr",
		},
	})
	if query.Error != nil {
		log.Println("Berserker create error: ", query.Error)
		return false
	}
	if query.RowsAffected == 0 {
		log.Println("Berserker did not finish create.")
		return false
	}
	var result *BerserkerStudent
	db.Debug().Where(&BerserkerStudent{AuthName: "berserker"}).Find(&result)
	js, _ := json.Marshal(result)
	log.Println(string(js))
	log.Println("Berserker test finished.")
	return true
}
