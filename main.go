package main

import (
	"Gopatungan/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "root:1234@tcp(127.0.0.1:3306)/gopatungan_db?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//
	//fmt.Println("Connection Opened to Database")
	//
	//var users []user.User
	//length := len(users)
	//
	//fmt.Println(length)
	//
	//db.Find(&users)
	//length = len(users)
	//
	//fmt.Println(length)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//	fmt.Println("==========")
	//}

	router := gin.Default()
	router.GET("/handler", handler)
	router.Run()
}
func handler(c *gin.Context) {
	dsn := "root:1234@tcp(127.0.0.1:3306)/gopatungan_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
