package controller_users

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/auth"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/AndrewSalko/salkodev.edms.go/email"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRegistrationRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user UserRegistrationRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	validationErr := validate.Struct(user)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	//перевести email до lower-case
	emailNormalized := strings.ToLower(user.Email)

	users := database_users.Users()

	count, err := users.CountDocuments(ctx, bson.M{"email": emailNormalized})

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error detected while fetching the email"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user with email already exists"})
		return
	}

	passwordHashed := auth.HashPassword(user.Password)

	userInfo := database_users.UserInfo{Name: user.Name, Email: user.Email, Password: passwordHashed}

	_, err = database_users.CreateUser(ctx, userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cancel()

	//потрібно надіслати Email з особливим посиланням-підтвердженням (токен для підтвердження)
	emailConfirmToken, errConfirm := auth.GenerateTokenForUserRegistration(emailNormalized)
	if errConfirm != nil {
		msg := fmt.Sprintf("error generate confirmation token: %s", errConfirm.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	//TODO: зробити шаблон email для підтв.реєстрації
	//TODO: подумати де брати справжній сервер (домен)
	emailBody := "Click on link to finish registration http://localhost:8080/users/confirmregistration?token=" + emailConfirmToken
	email.SendMail(emailNormalized, "SalkoDev EDMS registration", emailBody)

	resultData := gin.H{"message": "Check your email and confirm registration"}

	c.JSON(http.StatusOK, resultData)
}
