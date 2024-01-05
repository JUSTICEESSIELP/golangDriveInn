package routes

import (
	"driveinn_server/models"
	"driveinn_server/storage"
	"driveinn_server/utils"
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx iris.Context) {
	var registerInput RegisterUserInput
	err := ctx.ReadJSON(&registerInput)
	if err != nil {

	}

}

var newUser models.User

func getAndHandleUserExist(email string, newUserModel *models.User) (exist bool, err error) {
	userExistQuery := storage.DB.Where("email = ?", strings.ToLower(email)).Limit(1).Find(&newUserModel)

	fmt.Print("USER RESULT FROM DB", userExistQuery)

	if userExistQuery.Error != nil {
		return false, userExistQuery.Error
	}

	userExist := userExistQuery.RowsAffected > 0
	if userExist == true {
		return true, nil

	}

	return false, nil

}

func hashPassword(password string) (hashedPassword string, err error) {
	bytesPassword := []byte(password)

	bytesHashed, error := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)

	if error != nil {
		return "", error
	}
	return string(bytesHashed), nil
}

func getUserById(id string, ctx iris.Context) *models.User {

	var user models.User
	userExistQuery := storage.DB.Where("id = ?", id).Find(&user)

	if userExistQuery.Error != nil {
		utils.CreateInternalServerError(ctx)
		return nil

	}

	if userExistQuery.RowsAffected == 0 {
		utils.CreateError(iris.StatusNotFound, "Not found", "user not found", ctx)
	}

	return &user

}

func returnUserFromToken(user models.User, ctx iris.Context) {

}

func Login(ctx iris.Context) {

}

type RegisterUserInput struct {
	UserName string `json:"username" validate:"required,max-256"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}

type LoginUserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=256"`
}
