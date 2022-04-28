package config

import (
	model "bookshop/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
// Connect sqlite
func ConnectWithDB(){
	database,err := gorm.Open("sqlite3","test.db")
	if err != nil{
		panic("Failed to connect to database")
	}
	database.AutoMigrate(&model.User{},&model.Category{},&model.Post{})
	DB = database
}


// // Close connection
// func CloseDbConnection(db *gorm.DB){
// 	database,err := db.DB()
// 	if err != nil{
// 		panic("Failed to close connection")
// 	}
// 	database.Close()
// }