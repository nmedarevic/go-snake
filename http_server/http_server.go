package http_server

import (
	"fmt"
	"log"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

func handleIndexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

// func ListenAndServe(addr string, handler Handler) error
// type SnakeServer struct {
// 	store PlayerStore
// }

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type SnakeServer struct {
	router *http.ServeMux
}

func (s *SnakeServer) routes() {
	s.router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s *SnakeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// player := strings.TrimPrefix(r.URL.Path, "/snake")

	fmt.Fprint(w, "20")
	// return
	// if player == "Pepper" {
	// 	fmt.Fprint(w, "20")
	// 	return
	// }

	// if player == "Floyd" {
	// 	fmt.Fprint(w, "10")
	// 	return
	// }
}

func RunServer() {
	srvr := SnakeServer{
		router: http.NewServeMux(),
	}

	srvr.routes()
	// mux := http.NewServeMux()
	log.Fatal(http.ListenAndServe(":3000", srvr.router))
}

// func main() {
// 	fmt.Println("Hello, World!")
// 	RunServer()
// }
