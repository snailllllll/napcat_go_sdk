package napcat_go_sdk

type HandlerMessage interface {
	// HandleMessage 消息处理器
	HandleMessage(message *ReceiveMessage)
}

type MessageHandlerEvent interface {
	// PrivateMessageEvent 处理私聊信息
	PrivateMessageEvent(receiveMessage *ReceiveMessage)

	// GroupMessageEvent 处理群聊信息
	GroupMessageEvent(receiveMessage *ReceiveMessage)
}
