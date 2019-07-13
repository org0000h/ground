package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/org0000h/ground/internal/app/api"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	router.GET("/user/:uid", getUser)
	router.POST("/addUser/:uid", addUser)
	router.DELETE("/deluser/:uid", deleteUser)
	router.PUT("/moduser/:uid", modifyUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}