package main

import (
	"log"
	"net/http"

	"github.com/battlesnakeio/starter-snake-go/api"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	dump(decoded)

	moveDirection := "left"
	yourHeadCoord := decoded.You.Body[0]

	if yourHeadCoord.X < 3 {
		moveDirection = "up"
	}
	if yourHeadCoord.Y < 3 {
		moveDirection = "right"
	}
	if yourHeadCoord.X > 8 && yourHeadCoord.Y < 8 {
		moveDirection = "down"
	}
	if decoded.Turn > 100 {
		moveDirection = "left"
	}

	respond(res, api.MoveResponse{
		Move: moveDirection,
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
