package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
)

/*
	Archer: Student 中嵌套一个 User, students 表的外键和 users 表的主键同名，均为 username
	https://gorm.io/zh_CN/docs/belongs_to.html
*/

type ArcherUser struct {
	Username string `gorm:"primaryKey;comment:用户名"`
	Password string `gorm:"not null;comment:密码"`
}

type ArcherStudent struct {
	StuId    string      `gorm:"primaryKey;comment:学号"`
	Name     string      `gorm:"not null;comment:姓名"`
	Username string      `gorm:"not null;unique;comment:用户名"`
	User     *ArcherUser `gorm:"foreignKey:Username;references:Username"`
	// 期望: foreign key students.username references users(username)
}

func TestArcher(db *gorm.DB) bool {
	log.Println("Migrating model Archer......")
	//err := db.AutoMigrate(&ArcherUser{}, &ArcherStudent{})
	err := db.AutoMigrate(&ArcherStudent{}, &ArcherUser{}) // 按上面的 migrate 顺序无法运行, 创建外键约束时找不到表
	if err != nil {
		log.Println("Archer migration failed: ", err)
		return false
	}
	query := db.Debug().Create(&ArcherStudent{
		StuId:    "12345678",
		Name:     "射手",
		Username: "archer",
		User: &ArcherUser{
			Username: "archer",
			Password: "@2che2",
		},
	})
	if query.Error != nil {
		log.Println("Archer create error: ", query.Error)
		return false
	}
	if query.RowsAffected == 0 {
		log.Println("Archer did not finish create.")
		return false
	}
	var result *ArcherStudent
	db.Debug().Where(&ArcherStudent{StuId: "12345678"}).Find(&result)
	js, _ := json.Marshal(result)
	fmt.Println(string(js))
	log.Println("Archer test finished.")
	return true
}
