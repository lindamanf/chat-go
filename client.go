package main

import (
	"github.com/gorilla/websocket"
)

// チャットを行う1ユーザー
type client struct {
	socket *websocket.Conn
	// メッセージが送られるチャネル
	send chan []byte
	// クライアントが参加するチャットルーム
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
