package main

import (
	"sync"

	"golang.org/x/net/websocket"
)

var connection = make(map[*websocket.Conn]struct{})
var mu sync.Mutex

func addConn(conn *websocket.Conn) {
	mu.Lock()
	connection[conn] = struct{}{}
	mu.Unlock()
}

func delConn(conn *websocket.Conn) {
	mu.Lock()
	delete(connection, conn)
	mu.Unlock()
}

func getConns() []*websocket.Conn {
	conns := make([]*websocket.Conn, 0)
	mu.Lock()
	for k := range connection {
		conns = append(conns, k)
	}
	mu.Unlock()
	return conns
}

func getLen() int {
	mu.Lock()
	l := len(connection)
	mu.Unlock()
	return l
}
