package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/phamdinhha/go-chat-server/config"
	"github.com/phamdinhha/go-chat-server/internal/websocket"
	"github.com/phamdinhha/go-chat-server/pkg/http_error"
	"github.com/phamdinhha/go-chat-server/pkg/websocket_server"
)

type wsController struct {
	cfg    *config.Config
	wsPool *websocket.Pool
}

func NewWSController(cfg *config.Config, wsPool *websocket.Pool) *wsController {
	return &wsController{
		cfg:    cfg,
		wsPool: wsPool,
	}
}

func (ws *wsController) WSHandler(w http.ResponseWriter, r *http.Request) {
	// pool := websocket.NewPool()
	// go pool.Start()
	rID := r.URL.Query().Get("room_id")
	roomID, err := uuid.Parse(rID)
	if err != nil {
		handleWebsocketAuthenticationErr(w, err)
		return
	}
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
	serverWS(ws.wsPool, roomID, w, r, claims)
}

func (ws *wsController) WSGinHandler(c *gin.Context) {
	ws.WSHandler(c.Writer, c.Request)
}

var RegisterWebsocketRoute = func(router *gin.RouterGroup, cfg *config.Config) {
	pool := websocket.NewPool()
	go pool.Start()
	wsController := NewWSController(cfg, pool)
	router.GET("", wsController.WSGinHandler)
}

func serverWS(pool *websocket.Pool, roomID uuid.UUID, w http.ResponseWriter, r *http.Request, claims jwt.MapClaims) {
	conn, err := websocket_server.Upgrade(w, r)
	if err != nil {
		log.Println(err)
		return
	}
	client := &websocket.Client{
		ID:         uuid.New().String(),
		RoomID:     roomID,
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
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := map[string]interface{}{
		"data":   []string{},
		"errors": err.Error(),
	}
	data, _ := json.Marshal(res)
	w.Write(data)
}
