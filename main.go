package main

import (
	"Gopatungan/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:1234@tcp(127.0.0.1:3306)/gopatungan_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}
	userInput.Name = "Tes Simpan dari Service"
	userInput.Job = "Tester"
	userInput.Email = "test@example.com"
	userInput.Password = "password"

	userService.RegisterUser(userInput)
}
