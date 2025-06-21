package realtime

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Hub struct {
	clients map[string]*websocket.Conn
	mu      sync.RWMutex
}

var hub = Hub{
	clients: make(map[string]*websocket.Conn),
}

func (h *Hub) AddClient(userID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[userID] = conn
}

func (h *Hub) RemoveClient(userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, userID)
}

func (h *Hub) GetClient(userID string) (*websocket.Conn, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	conn, exists := h.clients[userID]
	return conn, exists
}
