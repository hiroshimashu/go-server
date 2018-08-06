package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	//socket for client
	socket *websocket.Conn

	//channel that is used for recieving messages
	send chan []byte

	//room that the client is participating in.
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
