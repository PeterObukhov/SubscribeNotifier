package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDb() *gorm.DB {
	connString := `host=%s user=%s password=%s dbname=%s port=%s sslmode=%s`
	dsn := fmt.Sprintf(
		connString,
		configuration.DbConfig.Host,
		configuration.DbConfig.User,
		configuration.DbConfig.Password,
		configuration.DbConfig.DbName,
		configuration.DbConfig.Port,
		configuration.DbConfig.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Subscription{})

	return db
}

func addUserOrSubscription(db gorm.DB, user User) bool {
	var userId uint
	err := db.Model(user).Select("Id").Where("login = ?", user.Login).Find(&userId).Error
	if err != nil {
		panic(err)
	}

	var res *gorm.DB
	if userId == 0 {
		res = db.Create(&user)
	} else {
		res = addSubscription(db, userId, user.Subscriptions[0])
	}

	if res.Error != nil {
		return false
	} else {
		return true
	}
}

func addSubscription(db gorm.DB, userId uint, subscription Subscription) *gorm.DB {
	subscription.UserID = userId
	res := db.Create(&subscription)
	return res
}
