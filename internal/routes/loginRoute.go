package routes

import (
	"ShelterChatBackend/Api/database"
	"ShelterChatBackend/Api/database/structs"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginRoute(ctx *gin.Context) {
	var requestObj structs.User
	var responseObj structs.User
	err := json.NewDecoder(ctx.Request.Body).Decode(&requestObj)
	if err != nil {
		panic(err)
	}

	if response := database.DB.Table("users").Where("email = ?", requestObj.Email).Or("username = ?", requestObj.Username).First(&responseObj); errors.Is(response.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusConflict, gin.H{
			"Status": "Conflict, No Matching Credentials found",
		})
		ctx.Done()
		return
	}

	if ok := bcrypt.CompareHashAndPassword([]byte(responseObj.Password), []byte(requestObj.Password)); ok != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"Status": "Conflict, No Matching Credentials found",
		})
		ctx.Done()
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": "Accepted, Logged you in Succesfully",
	})

}
