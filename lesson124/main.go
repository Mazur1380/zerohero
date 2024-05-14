package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {

	wsHandler := websocket.Handler(handler)

	http.Handle("/", wsHandler)
	http.ListenAndServe(":80", nil)
}

func handler(conn *websocket.Conn) {
	conn.SetDeadline(time.Now().Add(1 * time.Minute))
	addConn(conn)
	greet := fmt.Sprintf("Привет в чате: %d людей", getLen())
	err := websocket.Message.Send(conn, greet)
	if err != nil {
		log.Println(err)
		delConn(conn)
		return
	}

	for {
		var msg string
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			log.Println(err)
			delConn(conn)
			return
		}
		conn.SetDeadline(time.Now().Add(1 * time.Minute))
		cs := getConns()
		for _, c := range cs {
			if c == conn {
				continue
			}
			err = websocket.Message.Send(c, msg)
			if err != nil {
				log.Println(err)
				delConn(c)
				return

			}

		}
	}
}
