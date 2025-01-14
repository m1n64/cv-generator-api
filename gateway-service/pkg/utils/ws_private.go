package utils

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WebSocketPrivateManager struct {
	clients   map[string]map[*websocket.Conn]bool
	broadcast chan interface{}
	mutex     sync.Mutex
}

func NewWebSocketPrivateManager() *WebSocketPrivateManager {
	return &WebSocketPrivateManager{
		clients:   make(map[string]map[*websocket.Conn]bool),
		broadcast: make(chan interface{}),
	}
}

func (manager *WebSocketPrivateManager) AddClient(conn *websocket.Conn, userID string) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	if manager.clients[userID] == nil {
		manager.clients[userID] = make(map[*websocket.Conn]bool)
	}

	manager.clients[userID][conn] = true
}

func (manager *WebSocketPrivateManager) RemoveClient(conn *websocket.Conn) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	for userID, connections := range manager.clients {
		if _, exists := connections[conn]; exists {
			delete(connections, conn)
			conn.Close()
			if len(connections) == 0 {
				delete(manager.clients, userID)
			}
			break
		}
	}
}

func (manager *WebSocketPrivateManager) BroadcastMessage(message interface{}) {
	var clientsToRemove []struct {
		userID string
		conn   *websocket.Conn
	}

	manager.mutex.Lock()
	for userID, connections := range manager.clients {
		for conn := range connections {
			err := conn.WriteJSON(message)
			if err != nil {
				clientsToRemove = append(clientsToRemove, struct {
					userID string
					conn   *websocket.Conn
				}{userID, conn})
			}
		}
	}
	manager.mutex.Unlock()

	for _, client := range clientsToRemove {
		manager.RemoveClient(client.conn)
	}
}

func (manager *WebSocketPrivateManager) BroadcastToUser(userID string, message interface{}) {
	manager.mutex.Lock()
	connections, exists := manager.clients[userID]
	defer manager.mutex.Unlock()

	if !exists {
		return
	}

	for conn := range connections {
		err := conn.WriteJSON(message)
		if err != nil {
			manager.RemoveClient(conn)
		}
	}
}
