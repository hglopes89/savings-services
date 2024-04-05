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
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func GetUsers() []User {
	db := UsersDB()
	var users []User
	db.Find(&users)
	return users
}
