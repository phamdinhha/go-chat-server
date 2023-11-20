package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/pkg/http_error"
	"github.com/phamdinhha/go-chat-server/pkg/http_response"
	"github.com/phamdinhha/go-chat-server/pkg/websocket"
)

// var RegisterWebsocketRoute = func(router *mux.Router, cfg *config.Config) {
// 	pool := websocket.NewPool()
// 	go pool.Start()

// 	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		jwtToken := r.URL.Query().Get("token")
// 		jwtSecret := cfg.JWTConfig.JWTSecret
// 		parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != cfg.JWTConfig.SigningMethod {
// 				return nil, http_error.ErrInvalidCredentials
// 			}
// 			return []byte(jwtSecret), nil
// 		})
// 		if err != nil {
// 			handleWebsocketAuthenticationErr(w, err)
// 			return
// 		}
// 		claims, ok := parsedToken.Claims.(jwt.MapClaims)
// 		if !ok || !parsedToken.Valid {
// 			handleWebsocketAuthenticationErr(w, http_error.ErrInvalidCredentials)
// 			return
// 		}
// 		serverWS(pool, w, r, claims)
// 	})
// }

type wsController struct {
	cfg *config.Config
}

func NewWSController(cfg *config.Config) *wsController {
	return &wsController{
		cfg: cfg,
	}
}

func (ws *wsController) WSHandler(w http.ResponseWriter, r *http.Request) {
	pool := websocket.NewPool()
	go pool.Start()
	jwtToken := r.URL.Query().Get("token")
	jwtSecret := ws.cfg.JWTConfig.JWTSecret
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != ws.cfg.JWTConfig.SigningMethod {
			return nil, http_error.ErrInvalidCredentials
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		handleWebsocketAuthenticationErr(w, err)
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		handleWebsocketAuthenticationErr(w, http_error.ErrInvalidCredentials)
		return
	}
	serverWS(pool, w, r, claims)
}

func (ws *wsController) WSGinHandler(c *gin.Context) {
	ws.WSHandler(c.Writer, c.Request)
}

var RegisterWebsocketRoute = func(router *gin.RouterGroup, cfg *config.Config) {
	wsController := NewWSController(cfg)
	router.GET("", wsController.WSGinHandler)
}

func serverWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request, claims jwt.MapClaims) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	client := &websocket.Client{
		Connection: conn,
		Pool:       pool,
		Email:      claims["email"].(string),
		UserID:     claims["user_id"].(string),
	}
	pool.Register <- client
	requestBody := make(chan []byte)
	go client.Read(requestBody)
	// go br.ReadMessage()
}

func handleWebsocketAuthenticationErr(w http.ResponseWriter, err error) {
	log.Println("websocket error: ", err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := http_response.ErrorCtxResponse(&fiber.Ctx{}, err)
	data, _ := json.Marshal(res)
	w.Write(data)
}
