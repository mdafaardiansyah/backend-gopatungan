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

	router := gin.New()

	// Ambil URL frontend dari environment variable
	//frontendURL := os.Getenv("FRONTEND_URL")
	//if frontendURL == "" {
	//	log.Fatal("FRONTEND_URL environment variable is not set")
	//}
	router.Use(CORSMiddleware())

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

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

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, api_key, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Vary", "Origin") // Add this line

		if ctx.Request.Method == "OPTIONS" {
			log.Println("OPTIONS")
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
