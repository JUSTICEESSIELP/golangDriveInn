package utils

import "github.com/kataras/iris/v12"

func CreateError(status int, title string, detail string, ctx iris.Context) {
	ctx.StopWithError(status, iris.NewProblem().Title(title).Detail(detail))
}

func CreateValidationError(ctx iris.Context) {

}

func CreateInternalServerError(ctx iris.Context) {
	CreateError(
		iris.StatusInternalServerError,
		"Internal Server Error",
		"Internal server Error",
		ctx,
	)

}

func CreateNotFound(ctx iris.Context) {
	CreateError(
		iris.StatusNotFound, "Not found", "Resource not found", ctx,
	)
}
