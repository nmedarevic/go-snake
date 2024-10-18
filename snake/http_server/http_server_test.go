package http_server

import (
	// "http_server"

	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockSnakeGame struct {
	mock.Mock
}

func (m *MockSnakeGame) calculateTick(playerId string) *SnakeGameUnit {
	args := m.Called(playerId)
	fmt.Println(args)
	// quote, _ := args.Get(0).(*quote.Quote)
	return &SnakeGameUnit{}
}

func Test_RunServer(t *testing.T) {
	t.Run("Runs the server", func(t *testing.T) {
		srvr := SnakeServer{
			router: http.NewServeMux(),
		}
		srvr.routes()

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		response := makeHttpCall(srvr.router, req)

		if response.Code != http.StatusOK {
			t.Errorf("got %d, want %d", response.Code, http.StatusOK)
		}
	})

	t.Run("Generates a snake game", func(t *testing.T) {
		mockSnakeGameGenerator := MockSnakeGame{}

		srvr := SnakeServer{
			calculateTick: mockSnakeGameGenerator.calculateTick,
		}

		rr := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.URL.RawQuery = "playerId=1"

		mockSnakeGameGenerator.On("calculateTick", "1").Return(&SnakeGameUnit{})

		srvr.HandleTest(rr, req)

		mockSnakeGameGenerator.AssertExpectations(t)
	})
}

func makeHttpCall(router *http.ServeMux, request *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, request)

	return rr
}

// func TestRunServer(t *testing.T) {
// 	t.Run("Run server", func(t *testing.T) {
// 		http_server := &http_server.SnakeServer{}
// 		// http_server.RunServer()

// 		request, _ := http.NewRequest(http.MethodGet, "/snake", nil)
// 		response := httptest.NewRecorder()

// 		http_server.ServeHTTP(response, request)

// 		got := response.Body.String()
// 		want := "20"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }

// func TestGETSnake(t *testing.T) {
// 	http_server := &http_server.SnakeServer{}

// 	t.Run("Get snake playground", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/snake", nil)
// 		response := httptest.NewRecorder()

// 		http_server.ServeHTTP(response, request)

// 		got := response.Body.String()
// 		want := "20"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})

// }

// func TestRunFloyd(t *testing.T) {
// 	http_server := &http_server.SnakeServer{}
// 	t.Run("returns Floyd's score", func(t *testing.T) {
// 		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
// 		response := httptest.NewRecorder()

// 		http_server.ServeHTTP(response, request)

// 		got := response.Body.String()
// 		want := "10"

// 		if got != want {
// 			t.Errorf("got %q, want %q", got, want)
// 		}
// 	})
// }
