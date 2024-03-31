package users

import (
	"net/http"
	"regexp"

	"github.com/3fanyu/glossika/internal/dao"
	uc "github.com/3fanyu/glossika/internal/usecase/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, uc uc.UserUsecase) {
	userGroup := router.Group("/v1/user")
	{
		userGroup.POST("/register", CreateUser(uc))
		userGroup.POST("/sign_in", AuthenticateUser(uc))
		userGroup.GET("/verify", VerifyUser(uc))
	}
}

func CreateUser(uc uc.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dao.RegisterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !validatePassword(input.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password is invalid"})
			return
		}
		uc.CreateUser(c, input)
	}
}

func AuthenticateUser(uc uc.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dao.AuthInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		uc.Auth(c, input)
	}
}

func VerifyUser(uc uc.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input dao.VerifyInput
		input.Email = c.Query("email")
		input.Token = c.Query("token")
		if input.Email == "" || input.Token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and token are required"})
			return
		}
		uc.Verify(c, input)
	}
}

func validatePassword(password string) bool {
	// check length
	if len(password) < 6 || len(password) > 16 {
		return false
	}

	// define regular expression
	var (
		upperCase = `[A-Z]`                                       // atleast one upper case
		lowerCase = `[a-z]`                                       // atleast one lower case
		special   = `[()\[\]{}<>+\-*/?,.:;"'_\\|~` + "`!@#$%^&=]" // atleast one special character
	)

	// create regular expression object
	upperCaseRegex := regexp.MustCompile(upperCase)
	lowerCaseRegex := regexp.MustCompile(lowerCase)
	specialRegex := regexp.MustCompile(special)

	// validate password
	return upperCaseRegex.MatchString(password) &&
		lowerCaseRegex.MatchString(password) &&
		specialRegex.MatchString(password)
}
