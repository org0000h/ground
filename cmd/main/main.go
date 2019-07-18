package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/org0000h/ground/internal/app/service"
)

func main() {
	router := httprouter.New()
	router.GET("/", service.Index)
	router.GET("/hello/:name", service.Hello)

	router.GET("/user/:uid", service.GetUser)
	router.POST("/addUser/:uid", service.AddUser)
	router.DELETE("/deluser/:uid", service.DeleteUser)
	router.PUT("/moduser/:uid", service.ModifyUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}
