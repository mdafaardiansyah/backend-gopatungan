package main

import (
	"Gopatungan/auth"
	"Gopatungan/handler"
	"Gopatungan/user"
	"fmt"
	"github.com/gin-gonic/gin"
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
	authService := auth.NewService() //done testing postman

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.rbY_zh7v_we9_S3bMNX3V-rp7YLS944LOYZd524aKSE")

	if err != nil {
		fmt.Println("Error di Validate Token")
	}
	if token.Valid {
		fmt.Println("Token Valid")
	} else {
		fmt.Println("Token Invalid")
	}

	fmt.Println(authService.GenerateToken(1001)) //testing generate token

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)                    //tested
	api.POST("/sessions", userHandler.Login)                        //tested
	api.POST("/email_checkers", userHandler.CheckEmailAvailability) //tested
	api.POST("/avatars", userHandler.UploadAvatar)                  //tested

	router.Run()
}
