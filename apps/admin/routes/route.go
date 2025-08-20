package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// @Summary Get Admin API Info
		// @Description Returns a welcome message for the Admin API
		// @Tags Admin
		// @Accept json
		// @Produce json
		// @Success 200 {object} map[string]string
		// @Router /api/v1 [get]
		v1.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Admin API"})
		})
		RegisterUserRoutes(v1)
	}
}

// RegisterUserRoutes array vs slice group
func RegisterUserRoutes(router *gin.RouterGroup) {
	handler := Users{}

	group := router.Group("/user")
	group.GET("/get", handler.GetUsers)
	group.POST("/create", handler.CreateUser)
}

var listOfUsers []User
var lastID int

type Users struct{}

// GetUsers godoc
// @Summary List all users
// @Description Retrieve a list of all user items
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} User "List of users"
// @Router /api/v1/user/get [get]
func (e Users) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, listOfUsers)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user item to the list
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User item to create"
// @Success 201 {object} User "Created user item"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Router /api/v1/user/create [post]
func (e Users) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lastID++
	user.ID = lastID
	listOfUsers = append(listOfUsers, user)
	c.JSON(http.StatusCreated, user)
}

type User struct {
	ID        int    `json:"id" example:"1"`
	FirstName string `json:"firstname" example:"This is the user firstname" binding:"required"`
	LastName  string `json:"lastname" example:"This is the user lastname" binding:"required"`
	Age       int    `json:"age" example:"24" binding:"required"`
}
