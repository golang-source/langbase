package main

import (
	"github.com/golang-source/websocket"
	"net/http"
	"time"

	"langbasic/websocket/impl"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)

	if wsConn, err = upGrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		for {
			//撕裂一个空间不断写一个heartbeat
			if err := conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}

	}
ERR:
	conn.Close()
}

func main() {
	http.HandleFunc("/upgrade", wsHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}


