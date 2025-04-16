package napcat_go_sdk

type MsgType string

const (
	TEXT  MsgType = "text"
	FACE  MsgType = "face"
	IMAGE MsgType = "image"
	REPLY MsgType = "reply"
	AT    MsgType = "at"
)

type Action string

const (
	SEND_PRIVATE_MSG Action = "send_private_msg"
	FRIEND_POKE      Action = "friend_poke"
	SEND_GROUP_MSG   Action = "send_group_msg"
)

type MessageType string

const (
	PRIVATE MessageType = "private"
	GROUP   MessageType = "group"
)
