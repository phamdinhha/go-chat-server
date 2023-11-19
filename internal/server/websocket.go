package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/pkg/http_error"
	"github.com/phamdinhha/go-chat-server/pkg/http_response"
	"github.com/phamdinhha/go-chat-server/pkg/websocket"
)

var RegisterWebsocketRoute = func(router *mux.Router) {
	pool := websocket.NewPool()
	go pool.Start()

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.URL.Query().Get("token")
		jwtSecret := cfg.JWTConfig.JWTSecret
		parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != cfg.JWTConfig.SigningMethod {
				return nil, http_error.ErrInvalidCredentials
			}
		})
		if err != nil {
			handleWebsocketAuthenticationErr(w, err)
			return
		}
		websocket.ServeWs(pool, w, r)
	})
}

func handleWebsocketAuthenticationErr(w http.ResponseWriter, err error) {
	log.Println("websocket error: ", err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := http_response.ErrorCtxResponse(&fiber.Ctx{}, err)
	data, err := json.Marshal(res)
	w.Write(data)
}
