package app

import (
	"fmt"
	"invertory/auth"
	"invertory/domain"
	"invertory/errs"
	"invertory/handler"
	"invertory/logger"
	"invertory/service"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	//membuat setup cek untuk koneksi
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	logger.Info("load environment variables...")

	sanityCheck()
	//*membuat setup untuk repository
	inventoryRepository := domain.NewInventoryRepositoryDB(getClientDB())
	usersRepositoryDB := domain.NewUsersRepositoryDB(getClientDB())
	//*membuat setup untuk service
	serviceInventory := service.NewInventoryService(&inventoryRepository)
	usersService := service.NewUsersService(&usersRepositoryDB)
	authService := auth.NewService()
	//*membuat setup untuk handler
	inventoryHandler := handler.NewInventoryHandler(serviceInventory)
	usersHandler := handler.NewUsersHandler(usersService, authService)

	//*membuat setup router dengan Gin
	router := gin.Default()
	//* router url crud
	router.GET("/inventory", inventoryHandler.GetInventoryAll)
	router.GET("/inventory/:itemid", inventoryHandler.GetInventoryByID)
	router.POST("/inventory", inventoryHandler.CreateInventory)
	router.PUT("/inventory/:itemid", inventoryHandler.UpdateInventory)
	router.DELETE("/inventory/:itemid", inventoryHandler.DeleteInventory)

	//* url router register and login
	router.POST("/register", usersHandler.CreateUsers)
	router.POST("/login", usersHandler.LoginUsers)
	//*router port
	router.Run(":9000")
}

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environment variable %s not defined. terminating application...", envKey))
		}
	}

	logger.Info("environment variables loaded...")

}

func getClientDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("success connect to database...")

	return db
}

func authMiddleware(authService auth.Service, usersService service.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := errs.NewAuthenticationError("Unauthorized error")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := errs.NewAuthenticationError("Unauthorized error")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := errs.NewAuthenticationError("Unauthorized error")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		usersId := int(claim["user_id"].(float64))
		users, errss := usersService.GetUsersByID(usersId)
		if errss != nil {
			response := errs.NewAuthenticationError("Unauthorized error")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUsers", users)
	}
}
