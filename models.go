package main

type User struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Login         string `gorm:"unique"`
	ChatId        int64
	Subscriptions []Subscription `gorm:"foreignKey:UserID"`
}

type Subscription struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint
	Name   string
	Price  float32
	Date   int
}
