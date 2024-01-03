package routes

import "github.com/kataras/iris/v12"

func Register(ctx iris.Context) {
	var registerInput RegisterUserInput
	err := ctx.ReadJSON(&registerInput)
	if err != nil {
      
	}

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
