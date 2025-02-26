package main

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	GetConfiguration()
	db = connectToDb()
	startBot()
}
