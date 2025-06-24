package middleware

import (
	"github.com/g-s-pai/go-order-service/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

func JWTMiddleware(ctx iris.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.StopWithStatus(401)
		return
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		ctx.StopWithStatus(401)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	ctx.Values().Set("userID", claims["user_id"])
	ctx.Next()
}

func GenerateJWT(userId string) (string, error) {
	claims := models.CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 96)), // Expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecretKey := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(jwtSecretKey))
}