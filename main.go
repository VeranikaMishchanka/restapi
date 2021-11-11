package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)


type user struct {
    ID     string    `json:"id"`
    FirstName  string  `json:"firstname"`
    LastName string  `json:"lastname"`
    Email string `json:"email"`
    Age int `json:"age"`
    CreatedTime time.Time `json:"createdtime"`
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

    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

//GET[id]

func getUserByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
//PUT[id]
func updateUserByID (c *gin.Context) {
        id := c.Param("id")
        for _, a := range users {
            if a.ID == id {
                c.BindJSON(&a)
                c.IndetedJSON(http.StatusOK, a)
                return
            }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
    }

//DELETE[id]
func deleteUserByID(c *gin.Context) {
        id := c.Param("id")
        for i, a := range users {
            if a.ID == id {
                users = append(users[:i])
                c.IndetedJSON(http.StatusNoContent, a)
                return
            }
        }
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
    }
