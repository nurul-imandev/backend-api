package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"nurul-iman-blok-m/auth"
	"nurul-iman-blok-m/database"
	"nurul-iman-blok-m/handler"
	"nurul-iman-blok-m/helper"
	"nurul-iman-blok-m/role"
	"nurul-iman-blok-m/user"
	"strings"
)

func main() {
	db := database.Db()

	userRepository := user.NewRepository(db)
	roleRepository := role.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	roleService := role.NewRoleService(roleRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	roleHandler := handler.NewRoleHandler(roleService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.LoginUser)
	api.POST("/role/add", authMiddleware(authService, userService), roleHandler.SaveRole)
	api.GET("/roles", authMiddleware(authService, userService), roleHandler.GetRoles)

	router.Run(":8080")
}

func authMiddleware(autService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := autService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := uint(claim["user_id"].(float64))

		currentUser, errFindUser := userService.GetUserByID(userId)

		if errFindUser != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", currentUser)
	}

}
