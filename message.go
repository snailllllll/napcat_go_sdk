package napcat_go_sdk

type ReceiveMessage struct {
	Time          int         `json:"time"`
	SelfId        int         `json:"self_id"`
	PostType      string      `json:"post_type"`
	MetaEventType string      `json:"meta_event_type"`
	SubType       string      `json:"sub_type"`
	MessageId     int         `json:"message_id"`
	MessageSeq    int         `json:"message_seq"`
	RealId        int         `json:"real_id"`
	RealSeq       string      `json:"real_seq"`
	MessageType   MessageFrom `json:"message_type"`
	Sender        struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
	} `json:"sender"`
	RawMessage string `json:"raw_message"`
	Font       int    `json:"font"`
	Message    []struct {
		Type MsgType `json:"type"`
		Data struct {
			Text string `json:"text"`
		} `json:"data"`
	} `json:"message"`
	MessageFormat string `json:"message_format"`
	GroupId       *int   `json:"group_id"`
	TargetId      int    `json:"target_id"`
	Raw           struct {
		MsgId      string `json:"msgId"`
		MsgRandom  string `json:"msgRandom"`
		MsgSeq     string `json:"msgSeq"`
		CntSeq     string `json:"cntSeq"`
		ChatType   int    `json:"chatType"`
		MsgType    int    `json:"msgType"`
		SubMsgType int    `json:"subMsgType"`
		SendType   int    `json:"sendType"`
		SenderUid  string `json:"senderUid"`
		PeerUid    string `json:"peerUid"`
		ChannelId  string `json:"channelId"`
		GuildId    string `json:"guildId"`
		GuildCode  string `json:"guildCode"`
		FromUid    string `json:"fromUid"`
		FromAppid  string `json:"fromAppid"`
		MsgTime    string `json:"msgTime"`
		MsgMeta    struct {
		} `json:"msgMeta"`
		SendStatus     int    `json:"sendStatus"`
		SendRemarkName string `json:"sendRemarkName"`
		SendMemberName string `json:"sendMemberName"`
		SendNickName   string `json:"sendNickName"`
		GuildName      string `json:"guildName"`
		ChannelName    string `json:"channelName"`
		Elements       []struct {
			ElementType    int    `json:"elementType"`
			ElementId      string `json:"elementId"`
			ElementGroupId int    `json:"elementGroupId"`
			ExtBufForUI    struct {
			} `json:"extBufForUI"`
			TextElement struct {
				Content        string      `json:"content"`
				AtType         int         `json:"atType"`
				AtUid          string      `json:"atUid"`
				AtTinyId       string      `json:"atTinyId"`
				AtNtUid        string      `json:"atNtUid"`
				SubElementType int         `json:"subElementType"`
				AtChannelId    string      `json:"atChannelId"`
				LinkInfo       interface{} `json:"linkInfo"`
				AtRoleId       string      `json:"atRoleId"`
				AtRoleColor    int         `json:"atRoleColor"`
				AtRoleName     string      `json:"atRoleName"`
				NeedNotify     int         `json:"needNotify"`
			} `json:"textElement"`
			FaceElement            interface{} `json:"faceElement"`
			MarketFaceElement      interface{} `json:"marketFaceElement"`
			ReplyElement           interface{} `json:"replyElement"`
			PicElement             interface{} `json:"picElement"`
			PttElement             interface{} `json:"pttElement"`
			VideoElement           interface{} `json:"videoElement"`
			GrayTipElement         interface{} `json:"grayTipElement"`
			ArkElement             interface{} `json:"arkElement"`
			FileElement            interface{} `json:"fileElement"`
			LiveGiftElement        interface{} `json:"liveGiftElement"`
			MarkdownElement        interface{} `json:"markdownElement"`
			StructLongMsgElement   interface{} `json:"structLongMsgElement"`
			MultiForwardMsgElement interface{} `json:"multiForwardMsgElement"`
			GiphyElement           interface{} `json:"giphyElement"`
			WalletElement          interface{} `json:"walletElement"`
			InlineKeyboardElement  interface{} `json:"inlineKeyboardElement"`
			TextGiftElement        interface{} `json:"textGiftElement"`
			CalendarElement        interface{} `json:"calendarElement"`
			YoloGameResultElement  interface{} `json:"yoloGameResultElement"`
			AvRecordElement        interface{} `json:"avRecordElement"`
			StructMsgElement       interface{} `json:"structMsgElement"`
			FaceBubbleElement      interface{} `json:"faceBubbleElement"`
			ShareLocationElement   interface{} `json:"shareLocationElement"`
			TofuRecordElement      interface{} `json:"tofuRecordElement"`
			TaskTopMsgElement      interface{} `json:"taskTopMsgElement"`
			RecommendedMsgElement  interface{} `json:"recommendedMsgElement"`
			ActionBarElement       interface{} `json:"actionBarElement"`
			PrologueMsgElement     interface{} `json:"prologueMsgElement"`
			ForwardMsgElement      interface{} `json:"forwardMsgElement"`
		} `json:"elements"`
		Records             []interface{} `json:"records"`
		EmojiLikesList      []interface{} `json:"emojiLikesList"`
		CommentCnt          string        `json:"commentCnt"`
		DirectMsgFlag       int           `json:"directMsgFlag"`
		DirectMsgMembers    []interface{} `json:"directMsgMembers"`
		PeerName            string        `json:"peerName"`
		FreqLimitInfo       interface{}   `json:"freqLimitInfo"`
		Editable            bool          `json:"editable"`
		AvatarMeta          string        `json:"avatarMeta"`
		AvatarPendant       string        `json:"avatarPendant"`
		FeedId              string        `json:"feedId"`
		RoleId              string        `json:"roleId"`
		TimeStamp           string        `json:"timeStamp"`
		ClientIdentityInfo  interface{}   `json:"clientIdentityInfo"`
		IsImportMsg         bool          `json:"isImportMsg"`
		AtType              int           `json:"atType"`
		RoleType            int           `json:"roleType"`
		FromChannelRoleInfo struct {
			RoleId string `json:"roleId"`
			Name   string `json:"name"`
			Color  int    `json:"color"`
		} `json:"fromChannelRoleInfo"`
		FromGuildRoleInfo struct {
			RoleId string `json:"roleId"`
			Name   string `json:"name"`
			Color  int    `json:"color"`
		} `json:"fromGuildRoleInfo"`
		LevelRoleInfo struct {
			RoleId string `json:"roleId"`
			Name   string `json:"name"`
			Color  int    `json:"color"`
		} `json:"levelRoleInfo"`
		RecallTime   string `json:"recallTime"`
		IsOnlineMsg  bool   `json:"isOnlineMsg"`
		GeneralFlags struct {
		} `json:"generalFlags"`
		ClientSeq      string      `json:"clientSeq"`
		FileGroupSize  interface{} `json:"fileGroupSize"`
		FoldingInfo    interface{} `json:"foldingInfo"`
		MultiTransInfo interface{} `json:"multiTransInfo"`
		SenderUin      string      `json:"senderUin"`
		PeerUin        string      `json:"peerUin"`
		MsgAttrs       struct {
		} `json:"msgAttrs"`
		AnonymousExtInfo interface{} `json:"anonymousExtInfo"`
		NameType         int         `json:"nameType"`
		AvatarFlag       int         `json:"avatarFlag"`
		ExtInfoForUI     interface{} `json:"extInfoForUI"`
		PersonalMedal    interface{} `json:"personalMedal"`
		CategoryManage   int         `json:"categoryManage"`
		MsgEventInfo     interface{} `json:"msgEventInfo"`
		SourceType       int         `json:"sourceType"`
		Id               int         `json:"id"`
	} `json:"raw"`
}

type MsgType string

const (
	TEXT   MsgType = "text"
	FACE   MsgType = "face"
	IMAGE  MsgType = "image"
	REPLY  MsgType = "reply"
	AT     MsgType = "at"
	RECORD MsgType = "record"
	VIDEO  MsgType = "video"
	DICE   MsgType = "dice"
	RPS    MsgType = "rps"
	NODE   MsgType = "node"
)

type Action string

const (
	// 发送私聊
	SEND_PRIVATE_MSG Action = "send_private_msg"
	// 发送戳一戳
	FRIEND_POKE Action = "friend_poke"
	// 发送群聊信息
	SEND_GROUP_MSG Action = "send_group_msg"
	// 群聊戳一戳
	GROUP_POKE Action = "group_poke"
)

type MessageFrom string

const (
	// PRIVATE 私聊信息
	PRIVATE MessageFrom = "private"
	// GROUP 群聊信息
	GROUP MessageFrom = "group"
)

type Msg struct {
	Type MsgType `json:"type"`
	Data MsgData `json:"data"`
}

type MsgData struct {
	Text *string `json:"text"`
	Id   *int    `json:"id"`
	File *string `json:"file"`
	Data *string `json:"data"`
	Name *string `json:"name"`
}

type replyStatus struct {
	Status  string `json:"status"`
	Retcode int    `json:"retcode"`
	Data    struct {
		MessageId int `json:"message_id"`
	} `json:"data"`
	Message string `json:"message"`
	Wording string `json:"wording"`
	Echo    string `json:"echo"`
}

/*
	{
		"action": "",
		"params": {}
	}
*/
type Message[T any] struct {
	/*
	 需要发送的接口
	*/
	Action Action `json:"action"`
	/*
	 对应的参数
	*/
	Params T `json:"params"`
}

func (msg *Message[any]) SendWebSocketMsg() interface{} {
	return msg
}

func (msg *Message[any]) SendHttpMsg() interface{} {
	return msg.Params
}

type UserGroupId struct {
	UserId  *string `json:"user_id"`
	GroupId *string `json:"group_id"`
}

// Poke
/*
群聊poke
{
  "group_id": 0,
  "user_id": 0
}

私聊poke
{
  "user_id": 0
}
*/
type Poke = Message[UserGroupId]

type SendMsgContent struct {
	UserGroupId
	Messages []Msg `json:"message"`
}

// PrivateMsg
/*
私聊文本
{
  "user_id": "",
  "message": [
    {
      "type": "text",
      "data": {
        "text": ""
      }
    },
	{
    	"type": "image",
        "data": {
            	// 本地路径
            	"file": "file://D:/a.jpg"
            	// 网络路径
            	// "file": "http://i0.hdslb.com/bfs/archive/c8fd97a40bf79f03e7b76cbc87236f612caef7b2.png"
            	//base64编码
            	// "file": "base64://xxxxxxxx"
        	}
        },
		{
            "type": "face",
            "data": {
                "id": 37
                //参考 https://bot.q.qq.com/wiki/develop/api-v2/openapi/emoji/model.html#EmojiType
            }
        },
		{
            "type": "record",
            "data": {
                // 本地路径
                "file": "file://D:/a.mp3"

                // 网络路径
                // "file": "http://xxx.mp3"
            }
        }，
		{
            "type": "video",
            "data": {
                // 本地路径
                "file": "file://D:/a.mp4"

                // 网络路径
                // "file": "http://xxx.mp4"
            }
        }，
		//回复的消息
		{
            //第一个必须为reply
            "type": "reply",
            "data": {
                "id": 1263753202
            }
        },
        {
            "type": "text",
            "data": {
                "text": "回复你了"
            }
        },
		{
            "type": "file",
            "data": {
                // 本地路径
                "file": "file://D:/a.jpg",

                // 网络路径
                // "file": "http://i0.hdslb.com/bfs/archive/c8fd97a40bf79f03e7b76cbc87236f612caef7b2.png",

                // "file": "base64或DataUrl编码",

                "name": "a.jpg"
            }
        }
  ]
}
*/
type PrivateMsg = Message[SendMsgContent]

type GroupMsg = Message[SendMsgContent]
