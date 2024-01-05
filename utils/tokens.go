package utils

import (
	"context"
	"driveinn_server/storage"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var contextRedis = context.Background()

func CreateTokenPair(id uint) (tkPair *jwt.TokenPair, err error) {
	accessTokenEnv := os.Getenv("ACCESS_TOKEN_SIGNER")
	refreshTokenEnv := os.Getenv("REFRESH_TOKEN_SIGNER")

	accessTokenSigner := jwt.NewSigner(jwt.HS256, []byte(accessTokenEnv), 24*time.Hour)

	refresherTokenSigner := jwt.NewSigner(jwt.HS256, []byte(refreshTokenEnv), 365*24*time.Hour)

	userId := strconv.FormatUint(uint64(id), 10)

	accessTokenclaims := AccessTokenClaims{
		ID: id,
	}

	refreshTokenclaim := jwt.Claims{Subject: userId}

	accessToken, err := accessTokenSigner.Sign(accessTokenclaims)

	if err != nil {
		return nil, err

	}

	refreshToken, err := refresherTokenSigner.Sign(refreshTokenclaim)

	if err != nil {
		return nil, err

	}

	var tokenPair *jwt.TokenPair
	tokenPair.AccessToken = accessToken
	tokenPair.RefreshToken = refreshToken
	storage.Redis.Set(contextRedis, string(refreshToken), "true", 365*24*time.Hour+5*time.Minute).Err()

	// Send the generated token pair to the client.
	// The tokenPair looks like: {"access_token": $token, "refresh_token": $token}

	return tokenPair, nil

	// create a access and refresh token and send to the client
}

func RefreshToken(ctx iris.Context) {
	token := jwt.GetVerifiedToken(ctx)
	tokenStr := string(token.Token)

	validToken, tokenErr := storage.Redis.Get(contextRedis, tokenStr).Result()

	if tokenErr != nil {
		CreateNotFound(ctx)
		return
	}

	if validToken != string("true") {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	storage.Redis.Del(contextRedis, tokenStr)
	userID, parseErr := strconv.ParseUint(token.StandardClaims.Subject, 10, 32)
	if parseErr != nil {
		CreateInternalServerError(ctx)
		return
	}

	tokenPair, tokenPairErr := CreateTokenPair(uint(userID))
	if tokenPairErr != nil {
		CreateInternalServerError(ctx)
		return
	}

	ctx.JSON(iris.Map{
		"accessToken":  string(tokenPair.AccessToken),
		"refreshToken": string(tokenPair.RefreshToken),
	})
	// how to make a new refresh token
}

type ForgetPasswordToken struct {
	ID    uint   `json:"ID"`
	Email string `json:"email"`
}

type AccessTokenClaims struct {
	ID uint `json:"ID"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
