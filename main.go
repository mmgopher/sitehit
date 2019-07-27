package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"site-hit/logger"
	"site-hit/server"
	"site-hit/service"
)


func init() {

	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func main() {
	logger.Info("Start application")
	hitCounter := service.NewHitCounter(service.GobWrite{})
	srv := server.New(http.NewServeMux(),hitCounter)
	srv.Start()
}
