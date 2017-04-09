package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user")
	fmt.Fprintf(w, "This is an authenticated request")
	fmt.Fprintf(w, "Claim content:\n")
	for k, v := range user.(*jwt.Token).Claims.(jwt.MapClaims) {
		fmt.Fprintf(w, "%s :\t%#v\n", k, v)
	}

	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(claims["empId"])

})

func main() {
	r := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtPwd), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})

	r.Handle("/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(myHandler),
	))
	http.Handle("/", r)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login2", loginCustomer)
	http.ListenAndServe(":3001", nil)
}

const (
	jwtPwd = "Account_8CFB2EC534E14D56"
)

func login(w http.ResponseWriter, r *http.Request) {
	mySigningKey := []byte(jwtPwd)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(err)
	fmt.Fprintf(w, "%v", ss)
}

func loginCustomer(w http.ResponseWriter, r *http.Request) {
	mySigningKey := []byte(jwtPwd)

	type MyCustomClaims struct {
		EmpId string `json:"empId"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		"1234",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "account",
			Audience:  "Account",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(err)
	fmt.Fprintf(w, "%v", ss)
}
