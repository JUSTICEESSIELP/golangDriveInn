package utils

import "github.com/kataras/iris/v12"

func CreateError(status int, title string, detail string, ctx iris.Context) {
	ctx.StopWithError(status, iris.NewProblem().Title(title).Detail(detail))
}

func createValidationError(ctx iris.Context) {
	CreateError(409, "Validation error", "Enter all the required inputs", ctx)

}

func CreateInternalServerError(ctx iris.Context) {

}
