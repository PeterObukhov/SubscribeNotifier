package main

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = connectToDb()
	startBot()
}
