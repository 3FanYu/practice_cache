package users

import (
	"net/http"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context, db *gorm.DB, input RegisterInput) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create user
	userDAO := dao.NewUserDAO(db)
	user := models.User{Email: input.Email, EncryptedPassword: string(hashedPassword)}
	if err := userDAO.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	go sendVerificationEmail(db, user.Email)

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})

}

func sendVerificationEmail(db *gorm.DB, targetAddress string) {
	// Send verification email logic here
	token := "1234"
	emailDAO := dao.NewEmailDAO(db)
	email := models.Email{TargetAddress: targetAddress, VerifyLink: "localhost:3000/v1/user/verify?token=" + token}
	emailDAO.CreateEmail(&email)
}
