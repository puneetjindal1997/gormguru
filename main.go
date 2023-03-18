package main

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Welcome to your channel go guru ji

// Object Relational Mapping (ORM) is a technique used in creating a "bridge" between object-oriented programs and, in most cases, relational databases.

// When interacting with a database using OOP languages, you'll have to perform different operations like creating, reading, updating, and deleting (CRUD) data from a database. By design, you use SQL for performing these operations in relational databases.

// What is Orm?
// An ORM tool is software designed to help OOP developers interact with relational databases. So instead of creating your own ORM software from scratch, you can make use of these tools.

// "SELECT id, name, email, country, phone_number FROM users WHERE id = 20"  =>  users.GetById(20)

// less headache
// less code
// orm is here for you guys

// Hibernate, Apache OpenJPA => java
// Django, web2py => python
// Laravel, cakephp => php

// Advantages
// It speeds up development time for teams.
// Decreases the cost of development.
// Handles the logic required to interact with databases.
// Improves security. ORM tools are built to eliminate the possibility of SQL injection attacks.
// You write less code when using ORM tools than with SQL.

// Disadvantages
// Learning how to use ORM tools can be time consuming.
// They are likely not going to perform better when very complex queries are involved.
// ORMs are generally slower than using SQL.

// \dt
// \c database_Name

type UserTest struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"<-:create"`
	Email string `gorm:"<-:create"`
}

var db *gorm.DB
var err error

func init() {
	dsn := `host=localhost 
			user=test1 
			password=password 
			dbname=test 
			port=5432 
			sslmode=disable 
			TimeZone=Asia/Shanghai`
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	fmt.Println(db, err)
	if err != nil {
		return
	}
	db.AutoMigrate(&UserTest{})
	fmt.Println("Created!")
}

func main() {

	// createDbRecord(db)
	fetchWithConditionRecord(db)
	// updateRecord(db)
	// fetchAllRecord(db)
	// deleteRecord(db)
}

func createDbRecord(db *gorm.DB) {
	u := UserTest{Name: "Ram", Email: "ram@gmail.com"}
	resp := db.Create(&u)
	fmt.Println(resp.Error, resp.RowsAffected)
}

func fetchWithConditionRecord(db *gorm.DB) {
	var u UserTest
	id := "1; drop database name;"
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	tc := db.Where("id=?", idInt).Find(&u)
	fmt.Println(tc.RowsAffected)
	fmt.Println(u)
}

func fetchAllRecord(db *gorm.DB) {
	var u []UserTest
	db.Find(&u)
	fmt.Println(u)
}

func updateRecord(db *gorm.DB) {
	resp := db.Table("user_tests").Where("id=?", 2).Updates(map[string]interface{}{"name": "xyz", "email": "xyz@gmail.com"})
	fmt.Println(resp, resp.Error, resp.RowsAffected)
}

func deleteRecord(db *gorm.DB) {
	resp := db.Where("id=?", 2).Delete(&UserTest{})
	fmt.Println(resp, resp.Error, resp.RowsAffected)
}
