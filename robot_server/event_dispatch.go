package robotServer

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type FeishuEventRequest struct {
	EventId   string
	EventType string
	Token     string
	Event     map[string]any
}

func readFromDict(data map[string]any) FeishuEventRequest {
	eventType := ""
	token := ""
	eventId := ""
	event, _ := data["event"].(map[string]any)

	if _, exist := data["schema"]; exist {
		// v2
		header, _ := data["header"].(map[string]any)
		eventType, _ = header["event_type"].(string)
		token, _ = header["token"].(string)
		eventId, _ = header["event_id"].(string)
	} else {
		// v1
		eventType, _ = event["type"].(string)
		token, _ = event["token"].(string)
		eventId, _ = event["uuid"].(string)
	}

	return FeishuEventRequest{
		EventId:   eventId,
		EventType: eventType,
		Token:     token,
		Event:     event,
	}
}

func feishuEventHandler(c *gin.Context) {
	// Handler for Feishu Event Http Callback

	// [steps]
	// - decrypt if needed
	// - return to test event
	// - check data
	// - dispatch events

	// see: https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM

	// decrypt data if ENCRYPT is on
	var jsobj map[string]any
	if encryptKey := viper.GetString("feishu.ENCRYPT_KEY"); encryptKey != "" {
		rawJson := make(map[string]any)
		c.BindJSON(&rawJson)
		rawJsonStr, _ := rawJson["encrypt"].(string)
		jsonStr, err := decrypt(rawJsonStr, encryptKey)
		if err != nil {
			logrus.Error("Cannot decrypt request")
		}
		json.Unmarshal([]byte(jsonStr), &jsobj)
	} else {
		c.BindJSON(&jsobj)
	}

	// return to server test event
	if challenge, exists := jsobj["challenge"]; exists {
		// TODO: ??
		c.JSON(200, gin.H{"challenge": challenge})
		return
	}

	req := readFromDict(jsobj)
	logrus.Debug("Feishu Robot received a request: ", req)

	if !validateRequest(c, req.Token) {
		logrus.Error("Cannot validate event: ", jsobj)
		c.String(400, "验证错误")
		return
	}

	if eventRepeatDetect(req.EventId) {
		eventMap[req.EventType](req.Event)
	}

	// must return a HTTP:200 (with any body)
	c.String(200, "OK!")
}

func validateRequest(c *gin.Context, token string) bool {
	// check the token and hash in event request header

	if token != viper.GetString("VERIFICATION_TOKEN") {
		return false
	}
	timestamp := c.Request.Header.Get("X-Lark-Request-Timestamp")
	nonce := c.Request.Header.Get("X-Lark-Request-Nonce")
	signature := c.Request.Header.Get("X-Lark-Signature")
	body, _ := ioutil.ReadAll(c.Request.Body)

	return signature == calculateSignature(timestamp, nonce, viper.GetString("ENCRYPT_KEY"), string(body))
}

func calculateSignature(timestamp, nonce, encryptKey, bodystring string) string {
	// copied from: https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-security-verification

	var b strings.Builder
	b.WriteString(timestamp)
	b.WriteString(nonce)
	b.WriteString(encryptKey)
	b.WriteString(bodystring) //bodystring指整个请求体，不要在反序列化后再计算
	bs := []byte(b.String())
	h := sha256.New()
	h.Write(bs)
	bs = h.Sum(nil)
	sig := fmt.Sprintf("%x", bs)
	return sig
}

var eventList = make(map[string]bool)

func eventRepeatDetect(eventId string) bool {
	// 该函数用于检测事件是否已经被处理过,请在任意事件处理函数内部使用该该函数来记录event_id以及判定

	if _, repeated := eventList[eventId]; repeated {
		return true
	} else {
		eventList[eventId] = true
		return false
	}
}
