package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kataras/iris/v12"
)

func Autocomplete(ctx iris.Context) {
	limit := "10"

	locationQuery := ctx.URLParam("location")
	limitQuery := ctx.URLParam("limit")
	apiKey := os.Getenv("LOCATION_TOKEN")

	if limitQuery != "" {
		limit = limitQuery
	}

	thirdPartyLocationURL := "https://api.locationiq.com/v1/autocomplete?key=" + apiKey + "&q=" + locationQuery + "&limit=" + limit
	locationHelper(ctx, thirdPartyLocationURL)

}

func Search(ctx iris.Context) {

	SearchQuery := ctx.URLParam("location")

	apiKey := os.Getenv("LOCATION_TOKEN")

	thirdPartyLocationURL := "https://us1.locationiq.com/v1/search?key=" + apiKey + "&q=" + SearchQuery + "&format=json"
	locationHelper(ctx, thirdPartyLocationURL)
}

func locationHelper(ctx iris.Context, url string) {

	response, err := http.Get(url)

	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("location Request Error").DetailErr(err))
		return
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("location Request Error").DetailErr(err))
		return
	}

	var objMap []map[string]interface{}

	err = json.Unmarshal(databytes, &objMap)
	if err != nil {
		ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
			Title("JSON Unmarshal Error").DetailErr(err))
		return
	}
	ctx.JSON(objMap)

}
