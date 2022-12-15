package routes

import (
	"ShelterChatBackend/Api/database"
	"ShelterChatBackend/Api/database/structs"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterRoute(ctx *gin.Context) {
	var registerObj structs.User
	var responseObj structs.User

	err := json.NewDecoder(ctx.Request.Body).Decode(&registerObj)
	if err != nil {
		panic(err)
	}

	registerObj.UUID = uuid.NewString()

	pwdHashed, err := bcrypt.GenerateFromPassword([]byte(registerObj.Password), 8)
	if err != nil {
		panic(err)
	}
	registerObj.Password = string(pwdHashed)

	response := database.DB.Table("users").Where("email = ?", registerObj.Email).Or("username = ?", registerObj.Username).Limit(1).First(&responseObj)

	if errors.Is(response.Error, gorm.ErrRecordNotFound){
		res := database.DB.Table("users").Create(&registerObj)
		log.Println(res)

		log.Println(registerObj.Password, registerObj.Email, registerObj.UUID, registerObj.Username)
		ctx.JSON(http.StatusCreated, gin.H{
			"Created": registerObj.UUID,
		})
		ctx.Done()
	} else if response.Error == nil {
		log.Println(response.Error)
		ctx.JSON(http.StatusConflict, gin.H{
			"Status": "Conflict, Username or Email already Taken",
			"Conflicted Name": &responseObj.Username,
		})
		ctx.Done()
	}

}


