package main

import (
	"bwastartup/handler"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)

<<<<<<< HEAD
<<<<<<< HEAD
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)

=======
>>>>>>> 0c61164ccc255f9a984ef9b8004fed0f398f2c09
=======
>>>>>>> 0c61164ccc255f9a984ef9b8004fed0f398f2c09
	router.Run(":8080")

}
