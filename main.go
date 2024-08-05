package main

import (
	"fmt"
	"goProject/handler"
	"goProject/user"
	"log"

	// "time"

	// "os/user"
	// "net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/kunfayakun?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	input := user.LoginInput{
		Email:    "email@domain.com",
		Password: "password",
	}
	user, err := userService.Login(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(user.Name)
	fmt.Println(user.Email)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users/login", userHandler.Login)
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
	// userInput := user.RegisterUserInput{
	// 	Name:       "rico",
	// 	Email:      "rico@gmail.com",
	// 	Occupation: "admin tik",
	// 	Password:   "password",
	// }

	// userService.RegisterUser(userInput)

	// fmt.Println("Connected to database is good")

	// var users []user.User

	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println(user.Occupation)
	// 	fmt.Println(user.Role)
	// 	fmt.Println(user.Token)
	// 	fmt.Println(user.AvatarFileName)
	// 	fmt.Println(user.CreatedAt)
	// 	fmt.Println(user.UpdatedAt)
	// 	fmt.Println("--------------------------------")
	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/kunfayakun?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// 	//input
// 	// handler mapping input to struct
// 	//service mapping to user
// 	//repository save struct user to db
// 	//db
// }
