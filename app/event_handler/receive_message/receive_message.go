package receiveMessage

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type messageHandler func(*MessageEvent)

// for more detailed information, see https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
type MessageEvent struct {
	Sender struct {
		Sender_id struct {
			Union_id string `json:"union_id"`
			Open_id  string `json:"open_id"`
			User_id  string `json:"user_id"`
		} `json:"sender_id"`
		Sender_type string `json:"sender_type"`
		Tenant_key  string `json:"tenant_key"`
	} `json:"sender"`
	Message struct {
		Message_id   string `json:"message_id"`
		Root_id      string `json:"root_id"`
		Parent_id    string `json:"parent_id"`
		Create_time  string `json:"create_time"`
		Chat_id      string `json:"chat_id"`
		Chat_type    string `json:"chat_type"`
		Message_type string `json:"message_type"`
		Content      string `json:"content"`
		Metions      struct {
			Key string `json:"key"`
			Id  struct {
				Union_id string `json:"union_id"`
				Open_id  string `json:"open_id"`
				User_id  string `json:"user_id"`
			} `json:"id"`
			Name       string `json:"name"`
			Tenant_key string `json:"tenant_key"`
		} `json:"metions"`
	} `json:"message"`
}

// dispatch message, according to Chat type
func Receive(event map[string]any) {
	messageevent := MessageEvent{}
	map2struct(event, &messageevent)
	switch messageevent.Message.Chat_type {
	case "p2p":
		p2p(&messageevent)
	case "group":
		group(&messageevent)
	default:
		logrus.WithFields(logrus.Fields{"chat type": messageevent.Message.Chat_type}).Warn("Receive message, but this chat type is not supported")
	}
}

func map2struct(m map[string]interface{}, stru interface{}) {
	bytes, _ := json.Marshal(m)
	json.Unmarshal(bytes, stru)
}