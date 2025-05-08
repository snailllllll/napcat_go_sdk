package napcat_go_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn    *websocket.Conn  //websocket连接
	Handler []HandlerMessage //消息处理器
}

func NewWebSocketClient(url string, port uint, token *string) (*WebSocketClient, error) {

	connUrl := ""
	if token == nil {
		connUrl = fmt.Sprintf("ws://%s:%d", url, port)
	} else {
		connUrl = fmt.Sprintf("ws://%s:%d?access_token=%s", url, port, *token)
	}

	conn, _, err := websocket.DefaultDialer.Dial(connUrl, nil)

	return &WebSocketClient{conn: conn, Handler: make([]HandlerMessage, 0)}, err
}

func (client *WebSocketClient) SendMessage(message SendMsg) error {
	err := client.conn.WriteJSON(message.SendWebSocketMsg())
	return err
}

func (client *WebSocketClient) ReadMessage() (string, error) {
	_, message, err := client.conn.ReadMessage()
	var receiveMessage ReceiveMessage
	err = json.Unmarshal(message, &receiveMessage)
	if err != nil {
		var replyStatus replyStatus
		err = json.Unmarshal(message, &replyStatus)
		if err != nil {
			return "", err
		}
		return string(message), err
	}
	if client.Handler != nil {
		for _, handlerMessage := range client.Handler {
			if handlerMessage != nil {
				//异步执行处理器中对于消息的处理
				go func() {
					h := handlerMessage
					h.HandleMessage(&receiveMessage)
				}()
			}
		}
	}

	return string(message), err

}
