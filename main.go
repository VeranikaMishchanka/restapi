package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


// USERSSS
type user struct {
    ID     uuid    `json:"id"`
    FirstName  string  `json:"firstname"`
    LastName string  `json:"lastname"`
    Email string `json:"email"`
    Age unit `json:"age"`
    CreatedTime time `json:"createdtime"`
}


var users []user


func main() {
    router := gin.Default()
    router.POST("/users", postUser)
    router.GET("/users", getUserByID)
    router.PUT("/users", updateUserByID)
    router.DELETE("/users", deleteUserByID)
    router.Run("localhost:8080")
}

//POST
func postUser(c *gin.Context) {
    var newUser user

    if err := c.BindJSON(&newUser); err != nil {
        return
    }


    users=append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

//GET[id]

func getUserByID(c *gin.Context) {
    id := c.Param("id")

    // Loop ?
    for _, a := range users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})

//PUT[id]
func updateUserByID(c *gin.Context) {
    id := c.Params("id")
    //TBD
    //if not in slice error?

    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.JSON(http.StatusOK, user)
}

//DELETE[id]
func deleteUserByID(c *gin.Context) {
    id := c.Param("id")

    users:= Remove(user, users)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound,  gin.H{"message": "user not found"})
    } else {
	c.IndentedJSON(http.StatusOK, nil)
	    }
