package main

import (
	"fmt"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	key = "639247c036f549958f508131f030823a"
)

func main() {

	e := echo.New()

	e.GET("/login", login)

	r := e.Group("/check")
	r.Use(middleware.JWT([]byte(key)))
	r.GET("", check)

	e.Start(":5000")

}

func login(c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["empId"] = "xiao.xinmiao"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(t)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func check(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["empId"].(string)
	return c.String(http.StatusOK, "welcome "+name+"!")

}
