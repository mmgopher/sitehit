package main

import (
	"net/http"
	"site-hit/server"
	"site-hit/service"
)

func main() {

	hitCounter := service.NewHitCounter()
	s := server.New(http.NewServeMux(),hitCounter)
	s.Start()
}
