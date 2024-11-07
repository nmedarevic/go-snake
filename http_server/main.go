package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"snakegame.com/http/start"
	snake "snakegame.com/snake"
)

type SnakeGameUnit struct {
	// snake    snake.Snake
	playerId string
}

type SnakeGame interface {
	calculateTick(playerId string) *SnakeGameUnit
}

type Server struct {
	router        *http.ServeMux
	calculateTick func(playerId string) *SnakeGameUnit
}

func (s *Server) HandleTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

type HelloJSON struct {
	A int
}

func (s *Server) HandleJSONResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	helloJson := HelloJSON{
		A: 1,
	}
	json.NewEncoder(w).Encode(helloJson)
}

func (s *Server) routes() {
	s.router.HandleFunc("/test", s.HandleTest)
	s.router.HandleFunc("/test-a", s.HandleJSONResponse)
}

func RunServer() {
	srvr := Server{
		router: http.NewServeMux(),
	}

	srvr.routes()
	// mux := http.NewServeMux()
	log.Fatal(http.ListenAndServe(":3000", srvr.router))
}

func main() {
	fmt.Println("Hello, World!")

	snakeBody := make([]snake.Point, 0)
	snakeBody = append(snakeBody, snake.Point{X: int32(0), Y: int32(0)})
	snakeBody = append(snakeBody, snake.Point{X: int32(1), Y: int32(0)})
	snakeBody = append(snakeBody, snake.Point{X: int32(2), Y: int32(0)})

	snek := snake.Snake{
		Body: &snakeBody,
		Id:   snake.RandSeq(10),
	}

	snek.Print()

	data := map[string]interface{}{
		"snake": snek.Body,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
	}
	fmt.Printf("json data: %s\n", jsonData)

	component := start.Hello(string(jsonData))

	http.Handle("/", templ.Handler(component))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
