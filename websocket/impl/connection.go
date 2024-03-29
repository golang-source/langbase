package impl

import (
"errors"
"github.com/golang-source/websocket"
"sync"
)

type Connection struct {
wsConn    *websocket.Conn
inChan    chan []byte
outChan   chan []byte
closeChan chan byte

mutex    sync.Mutex
isClosed bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
conn = &Connection{
wsConn:    wsConn,
inChan:    make(chan []byte, 1000),
outChan:   make(chan []byte, 1000),
closeChan: make(chan byte, 1),
}

//启动读协程
go conn.readLoop()
//启动写协程
go conn.writeLoop()
return
}

//API
func (conn *Connection) ReadMessage() (data []byte, err error) {
select {
case data = <-conn.inChan:
case <-conn.closeChan:
err = errors.New("connection is closed")
}
return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
select {
case conn.outChan <- data:
case <-conn.closeChan:
err = errors.New("connection is closed")
}
return
}

func (conn *Connection) Close() {
//线程安全，可重入
conn.wsConn.Close()

//要保证close方法只被执行一次
//close(conn.closeChan)
conn.mutex.Lock()
if !conn.isClosed {
close(conn.closeChan)
conn.isClosed = true
}
conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
var (
data []byte
err  error
)

for {

if _, data, err = conn.wsConn.ReadMessage(); err != nil {
goto ERR
}

//可能会阻塞，等待inChan,因为inChan只有1000的容量
select {
case conn.inChan <- data:
case <-conn.closeChan:
//当closeChan被关闭的时候
goto ERR
}
}
ERR:
conn.Close()
}

func (conn *Connection) writeLoop() {

var (
data []byte
err  error
)

for {
select {
case data = <-conn.outChan:
case <-conn.closeChan:
goto ERR
}

if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
goto ERR
}
}
ERR:
conn.Close()
}