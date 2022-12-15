package database

import (
	"ShelterChatBackend/Api/database/structs"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	//"os"
	//"fmt"
)

//TODO: Change everything back to postgresql

var(
	DB *gorm.DB
	err error
)


func SetupDatabase(){

	//dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Berlin", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("USER_PWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	dsn := "/test_data_securecord?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := DB.AutoMigrate(&structs.User{}); err != nil {
		panic(err)
	}
}