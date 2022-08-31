package dispatcher

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Dispatcher(c *gin.Context) {
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

func eventRepeatDetect(eventId string) bool {
	if _, repeated := eventIdList[eventId]; repeated {
		return true
	} else {
		eventIdList[eventId] = true
		return false
	}
}
