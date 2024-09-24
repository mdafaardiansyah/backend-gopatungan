package main

import (
	"Gopatungan/handler"
	"Gopatungan/helper"
	"Gopatungan/internal/auth"
	campaign2 "Gopatungan/internal/campaign"
	"Gopatungan/internal/payment"
	transaction2 "Gopatungan/internal/transaction"
	user2 "Gopatungan/internal/user"
	"Gopatungan/pkg/config"
	"Gopatungan/pkg/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func main() {
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// Load .env file
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Buat koneksi ke database
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user2.NewRepository(db)
	campaignRepository := campaign2.NewRepository(db)
	transactionRepository := transaction2.NewRepository(db)

	userService := user2.NewService(userRepository)
	campaignService := campaign2.NewService(campaignRepository)
	authService := auth.NewService() //done testing postman
	paymentService := payment.NewService()
	transactionService := transaction2.NewService(transactionRepository, campaignRepository, paymentService)

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	router := gin.Default()
	//router.Use(cors.Default())

	// Custom CORS configuration
	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowOrigins = []string{"http://localhost:3000"} // Replace with your frontend's origin
	defaultConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(defaultConfig))

	router.Static("/images", "./assets/images")

	api := router.Group("/api/v1")

	//User
	api.POST("/users", userHandler.RegisterUser)                                             //tested
	api.POST("/sessions", userHandler.Login)                                                 //tested
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)                          //tested
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar) //tested -middleware
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser) //tested -middleware

	//Campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)                                                  // tested
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)                                               // tested
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)     //tested
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)  //tested
	api.POST("/campaigns-images", authMiddleware(authService, userService), campaignHandler.UploadImage) //tested

	//Transaction
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions) // tested
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)                   //tested
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user2.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
} //tested middleware postman
