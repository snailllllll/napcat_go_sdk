package napcat_go_sdk

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn             *websocket.Conn  //websocket连接
	Handler          []HandlerMessage //消息处理器
	responseChannels sync.Map         //等待响应的通道
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

func (client *WebSocketClient) SendMessage(message Message[any]) (string, error) {
	msg := message.SendWebSocketMsg()
	// 创建响应通道并存储到sync.Map中

	echo_id := message.Echo
	responseChan := make(chan []byte)
	client.responseChannels.Store(echo_id, responseChan)

	err := client.conn.WriteJSON(msg)

	// 阻塞等待通道有数据写入，读取响应并关闭通道
	select { // 阻塞等待通道有数据写入，超时时间设置为 20 秒
	case response := <-responseChan:
		// 关闭通道
		close(responseChan)
		// 返回响应
		return string(response), err
	case <-time.After(20 * time.Second):
		// 超时后关闭通道
		close(responseChan)
		// 返回超时错误
		return "", fmt.Errorf("timeout")
	}

}

func (client *WebSocketClient) ReadMessage() (string, error) {
	_, message, err := client.conn.ReadMessage()
	var receiveMessage ReceiveMessage
	err = json.Unmarshal(message, &receiveMessage)
	if err != nil {
		var replyStatus apiResponse
		err = json.Unmarshal(message, &replyStatus)
		if err != nil {
			return "", err
		}
		// 根据 echo 将响应发送回原始通道
		if replyStatus.Echo != "" {
			if v, ok := client.responseChannels.Load(replyStatus.Echo); ok {
				responseChan := v.(chan []byte)
				responseChan <- message
				return string(message), nil
			}
		}
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
