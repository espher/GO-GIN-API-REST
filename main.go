package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type user_data struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []user_data{
	{Id: 1, Name: "Pedro", Age: 20},
	{Id: 2, Name: "Hufo", Age: 22},
	{Id: 3, Name: "Paco", Age: 35},
	{Id: 4, Name: "Luis", Age: 44},
}

func home(u *gin.Context) {
	u.IndentedJSON(http.StatusOK, "default")
}

func getUsers(u *gin.Context) {
	u.IndentedJSON(http.StatusOK, users)
}

func postUsers(u *gin.Context) {
	var new_user user_data

	if err := u.BindJSON(&new_user); err != nil {
		return
	}

	users = append(users, new_user)
	u.IndentedJSON(http.StatusCreated, users)
}

func getUser(u *gin.Context) {
	id := u.Param("id")
	id_int, err := strconv.Atoi(id)

	if err != nil {
		return
	}

	for _, v := range users {
		if v.Id == id_int {
			u.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	u.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.GET("/", home)

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", postUsers)

	router.Run("localhost:9898")
}
