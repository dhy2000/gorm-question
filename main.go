package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm-question/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

func InitDatabase() {
	type DbInit struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	}
	var init DbInit
	var err error
	if err = viper.UnmarshalKey("database", &init); err != nil {
		log.Fatalln("error parsing database config.")
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s", init.Host, init.Port, init.User, init.Password, init.Dbname)
	conn := postgres.Open(dsn)
	Db, err = gorm.Open(conn, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalln("connect database failed.")
	}
	log.Println("database connected......")
}

/*
	用户通过用户名和密码登录, 以用户名作为主键，必须唯一
	学生包括学号，姓名。学号为主键。
	系统中存在多种身份的用户(例如学生与教师), 本例中暂时只考虑学生
	一个用户与一个学生或者一个教师绑定, 也就是让学生/教师分别有一个用户名外键

	建表的基本 SQL 语句:

	create table users (
		username varchar primary key,
		password varchar not null
	);

	create table students (
		stu_id varchar primary key,
		name varchar not null,
		username varchar not null unique references user(username)
	);

	本项目中提供了 4 中 gorm 模型的构建方法, 其模型描述与测试函数位于 `models` 包中。每个文件代表一种模型的写法。

	在 archer 和 assassin 两个示例中, students 的外键及被其 references 的 users 主键名称相同, 均为 username.
	而 berserker 和 caster 两个示例中，students 的外键名采用 auth_name 与 users 的主键区分开。
*/

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("failed to read config.")
	}
	InitDatabase()
	// 以下分别对四种写法的 gorm 模型进行测试，测试内容包括 AutoMigrate, 插入一条数据以及查询刚刚插入的数据
	// 为了确保插入数据的顺利, 每次运行本程序前都需要 drop all tables.
	log.Println("Test begin ......")
	results := make(map[string]bool)
	results["archer"] = models.TestArcher(Db)
	results["assassin"] = models.TestAssassin(Db)
	results["berserker"] = models.TestBerserker(Db)
	results["caster"] = models.TestCaster(Db)
	log.Println("Test end.")
	fmt.Println("run results: ", results)
}

/*
测试结果:

表结构(Migrate):
	archer 创建的表结构以及外键约束与期望是不一致的:
		- users 表的主键 username 作为了外键并指向了 students.username
		- 表的依赖关系颠倒，期望 students 以来 users 表, 但模型 migrate 出来是相反的
	assassin, berserker, caster 的表结构正确

数据插入:
	assassin, berserker, caster 插入均正确, archer 的表结构已经是错误的, 插入结果无意义

数据查询:
	assassin, berserker, caster 均无法正确查询到结构体内嵌的子结构体
*/
