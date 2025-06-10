package handlers

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
	"users-api-gin/db"
	"users-api-gin/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Helper: Generate a random alphanumeric ID
func generateUserID(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, n)
	for i := range id {
		id[i] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

// Helper: JSON error response
func respondWithError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}

// Register a new user
func Register(c *gin.Context) {
	var input struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if input.Username == "" {
		respondWithError(c, http.StatusBadRequest, "Username is required")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not hash password")
		return
	}

	userID := generateUserID(8)
	user := models.User{
		UserID:    userID,
		Username:  input.Username,
	}

	account := models.Account{
		UserID:   userID,
		Username: input.Username,
		Password: string(hashed),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		respondWithError(c, http.StatusInternalServerError, "Could not create user")
		return
	}

	if err := db.DB.Create(&account).Error; err != nil {
		respondWithError(c, http.StatusBadRequest, "Username already exists")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"user_id": userID,
	})
}

// Login a user
func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	var account models.Account
	if err := db.DB.Where("username = ?", input.Username).First(&account).Error; err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(input.Password)); err != nil {
		respondWithError(c, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": account.UserID,
	})
}

// Get all users
func GetUser(c *gin.Context) {
	var users []models.User
	db.DB.Preload("Account").Find(&users)
	c.JSON(http.StatusOK, users)
}

// Add a user directly (admin use)
func AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if user.Firstname == "" || user.Lastname == "" || user.Course == "" || user.Track == "" {
		respondWithError(c, http.StatusBadRequest, "All fields are required")
		return
	}

	db.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		respondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, user)
}

// Get user by name (first, last, or full name)
func GetUserByName(c *gin.Context) {
	inputName := strings.TrimSpace(strings.ToLower(c.Param("fullname")))
	var users []models.User
	var matched []models.User

	db.DB.Find(&users)

	for _, u := range users {
		full := strings.ToLower(strings.TrimSpace(u.Firstname + " " + u.Lastname))
		first := strings.ToLower(strings.TrimSpace(u.Firstname))
		last := strings.ToLower(strings.TrimSpace(u.Lastname))

		if inputName == full || inputName == first || inputName == last {
			matched = append(matched, u)
		}
	}

	if len(matched) == 0 {
		respondWithError(c, http.StatusNotFound, "Name not found")
		return
	}

	c.JSON(http.StatusOK, matched)
}

// Update an existing user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		respondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	var updated models.User
	if err := c.ShouldBindJSON(&updated); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid input")
		return
	}

	db.DB.Model(&user).Updates(updated)
	c.JSON(http.StatusOK, user)
}

// Delete user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
		respondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
