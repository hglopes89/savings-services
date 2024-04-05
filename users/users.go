package users

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func UsersDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=users port=5432 sslmode=disable TimeZone=UTC"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

type User struct {
	ID   uint
	Name string
}

func GetUsers() []User {
	db := UsersDB()
	var users []User
	db.Find(&users)
	return users
}

func CreateUser(user User) {
	db := UsersDB()
	db.Create(&user)
}

func CreateUserMock() {
	db := UsersDB()
	db.Create(&User{Name: "John"})
	db.Create(&User{Name: "Doe"})
	db.Create(&User{Name: "Jane"})
	db.Create(&User{Name: "h"})
	db.Create(&User{Name: "a"})
	db.Create(&User{Name: "d"})
	db.Create(&User{Name: "i"})
	db.Create(&User{Name: "s"})
}
