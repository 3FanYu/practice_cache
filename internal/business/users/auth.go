package users

import (
	"net/http"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Auth(c *gin.Context, db *gorm.DB, input AuthInput) {
	// Find the user by email
	userDAO := dao.NewUserDAO(db)
	user, err := userDAO.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// Generate JWT token
	token, err := generateJWTToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Set the authorization header
	c.Writer.Header().Set("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{"message": "Signed in successfully"})

}
