package handler

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/redis"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

//H is my connection
var H = hub{
	broadcast:  make(chan message),
	register:   make(chan subscription),
	unregister: make(chan subscription),
	rooms:      make(map[string]map[*connection]bool),
}

type hub struct {
	rooms      map[string]map[*connection]bool
	broadcast  chan message
	register   chan subscription
	unregister chan subscription
}

type message struct {
	Data   string `json:"message"`
	room   string
	Sender string `json:"sender"`
}

type subscription struct {
	conn *connection
	room string
}

type connection struct {
	ws       *websocket.Conn
	username string
	send     chan []byte
}

//OpenWebSocket opens a connection to our nice websocket
func OpenWebSocket(c *gin.Context, roomID string) {
	uuid, err := redis.GetLoggedInUUID(c.Request)
	if err != nil {
		switch err {
		case dbutils.ErrUnauthorized:
			c.String(401, err.Error())
		case dbutils.ErrInternalServer:
			c.String(500, err.Error())
		default:
			c.String(500, "internal server error")
		}
		return
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(500, "Server Error")
		return
	}
	var user dbutils.User
	err = dbutils.SelectUserByUUID(uuid, &user)
	if err != nil {
		c.String(500, "Server Error")
		return
	}
	co := &connection{send: make(chan []byte, 256), ws: ws, username: user.Username}
	s := subscription{co, roomID}
	H.register <- s
	go s.writePump()
	go s.readPump()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:3000" || r.Header.Get("Origin") == "http://localhost:8080"
	},
}

func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

func (s subscription) readPump() {
	c := s.conn
	defer func() {
		H.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := message{string(msg), s.room, c.username}
		H.broadcast <- m
	}
}

func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (H *hub) Run() {
	for {
		select {
		case s := <-H.register:
			connections := H.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				H.rooms[s.room] = connections
			}
			H.rooms[s.room][s.conn] = true
		case s := <-H.unregister:
			connections := H.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(H.rooms, s.room)
					}
				}
			}
		case m := <-H.broadcast:
			connections := H.rooms[m.room]
			data, _ := json.Marshal(m)
			for c := range connections {
				select {
				case c.send <- data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(H.rooms, m.room)
					}
				}
			}
		}
	}
}
