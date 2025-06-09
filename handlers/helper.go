package handlers

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "users-api-gin/db"
    "users-api-gin/models"
)

func GetUser(c *gin.Context){
	var users []models.User
	db.DB.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}
func AddUser(c *gin.Context){
	var user models.User

	if err := c.BindJSON(&user); err != nil{
		return
	}
	if user.Firstname == "" || user.Lastname == "" || user.Course == "" || user.Track == "" || user.Status == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
    return
}
	db.DB.Create(&user)
	c.IndentedJSON(http.StatusCreated, user)
}
func GetUserByID(c *gin.Context){
	id := c.Param("id")
	var user models.User

	result := db.DB.First(&user, "id=?", id)
	if result.Error != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, user)
}

func GetUserByName(c *gin.Context) {
	inputName := strings.TrimSpace(strings.ToLower(c.Param("fullname")))
	var users []models.User
	db.DB.Find(&users)
	var matched []models.User

	for _, n := range users {
		fullName := strings.TrimSpace(strings.ToLower(n.Firstname + " " + n.Lastname))
		fName := strings.TrimSpace(strings.ToLower(n.Firstname))
		lName := strings.TrimSpace(strings.ToLower(n.Lastname))

		if inputName == fullName || inputName == fName || inputName == lName {
			matched = append(matched, n)
		}
	}

	if len(matched) > 0 {
		c.IndentedJSON(http.StatusOK, matched)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Name not found"})
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Check if user exists
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind new data
	var updatedData models.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	db.DB.Model(&user).Updates(updatedData)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// Check if user exists
	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete user
	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
