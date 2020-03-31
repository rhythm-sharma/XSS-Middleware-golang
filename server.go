package main

import (
	"/github.com/rrhythmsharma/goLang-ECHO/xss-middleware.go"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)


// it will store the  data coming from user input 
var (
	users = map[int]*user{}
	seq   = 1
)


//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	
	if err := c.Bind(u); err != nil {
		return err
	}
	
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}


func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}


func getAllUsers(c echo.Context) error {
	println(users)
	return c.JSON(http.StatusOK, users)
}


func updateUser(c echo.Context) error {	
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}


func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}


//----------
// Main function
//----------
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.GET("/users/all", getAllUsers)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}