package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtkey = []byte("maxwest")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func main() {
	r := gin.Default()
	r.GET("/set",set)
	r.GET("/get",get)
	r.Run(":9090")
}
// 颁发token
func set(ctx *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "max",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(200,gin.H{
		"token": tokenString,
	})
}
// 解析token
func get(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"code": 401,
			"msg": "权限不足",
		})
		ctx.Abort()
		return
	}
	token, claims, err := parseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"code": 401,
			"msg": "权限不足",
		})
		ctx.Abort()
		return
	}
	fmt.Println("parsing...")
	fmt.Println(claims.UserId)
}
func parseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	return token, claims, err
}