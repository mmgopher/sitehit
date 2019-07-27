package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"site-hit/config"
	"site-hit/service"
)

type Response struct{
	Counter int `json:"site_hit_counter"`
}

type Server struct {
	router     *http.ServeMux
	hitCounter service.HitCounter
}

func (s *Server) routes() {
	s.router.HandleFunc("/", s.handleHits())
}


func (s *Server) handleHits() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseWithJson(w,http.StatusOK,Response{s.hitCounter.IncrementAndGetCounter()})

	})
}

func (s *Server) Start() {
	s.routes()
	log.Fatal(http.ListenAndServe(config.GetConfiguration().HttpAddr, s.router))
}

func New(mux *http.ServeMux, hitCounter service.HitCounter) *Server {
	return &Server{router: mux, hitCounter: hitCounter}
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
