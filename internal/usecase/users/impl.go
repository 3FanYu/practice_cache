package users

import (
	"net/http"

	"github.com/3fanyu/glossika/internal/dao"
	"github.com/3fanyu/glossika/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewUsecase(userDAO UserDAO, emailDAO EmailDAO) UserUsecase {
	return &impl{userDAO: userDAO, emailDAO: emailDAO}
}

type impl struct {
	userDAO  UserDAO
	emailDAO EmailDAO
}

func (im *impl) CreateUser(c *gin.Context, input dao.RegisterInput) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	// Create user
	user := models.User{Email: input.Email, EncryptedPassword: string(hashedPassword)}
	if err := im.userDAO.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	go sendVerificationEmail(im.emailDAO, user.Email)

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})

}

func (im *impl) Auth(c *gin.Context, input dao.AuthInput) {
	// Find the user by email
	user, err := im.userDAO.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
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

func sendVerificationEmail(dao EmailDAO, targetAddress string) {
	// Send verification email logic here
	token := "1234"
	email := models.Email{TargetAddress: targetAddress, VerifyLink: "localhost:3000/v1/user/verify?token=" + token}
	dao.CreateEmail(&email)
}
