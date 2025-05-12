package napcat_go_sdk

type SendMsg interface {
	SendWebSocketMsg() interface{}
	SendHttpMsg() interface{}
}

type ReceiveResponseMessage interface {
	ReceiveResponseMessage(bytes []byte) error
}
