package main

import (
	"go-mongodb/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	// CREATED A NEW INSTANCE!
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	// MAKING ROUTES!
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/delete/:id", uc.DeleteUser)

	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("")

	if err != nil {
		panic(err)
	}

	return s
}
