package dispatcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"xlab-feishu-robot/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary feishu event dispatcher
// @Tags feishu_event
// @Accept json
// @Success 200 {string} OK
// @Router /feishu_events [post]
func Dispatcher(c *gin.Context) {
	// Handler for Feishu Event Http Callback

	// [steps]
	// - decrypt if needed
	// - return to test event
	// - check data
	// - dispatch events

	// see: https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM

	// get raw body (bytes)
	rawBody, _ := ioutil.ReadAll(c.Request.Body)

	// decrypt data if ENCRYPT is on
	var requestStr string
	if encryptKey := config.C.Feishu.EncryptKey; encryptKey != "" {
		rawBodyJson := make(map[string]any)
		json.Unmarshal(rawBody, &rawBodyJson)
		rawRequestStr, _ := rawBodyJson["encrypt"].(string)
		var err error
		requestStr, err = decrypt(rawRequestStr, encryptKey)
		if err != nil {
			logrus.Error("Cannot decrypt request")
		}
	} else {
		requestStr = string(rawBody)
	}

	var req FeishuEventRequest
	deserializeRequest(requestStr, &req)
	logrus.Debug("Feishu Robot received a request: ", req)

	// return to server test event
	if req.Challenge != "" {
		c.JSON(http.StatusOK, gin.H{"challenge": req.Challenge})
		return
	}

	if !validateRequest(c, req.Token, string(rawBody)) {
		logrus.Error("Cannot validate event: ", req)
		c.String(http.StatusBadRequest, "验证错误")
		return
	}

	if eventRepeatDetect(req.EventId) {
		logrus.Warning("Repeated event: ", req)
		c.String(http.StatusBadRequest, "事件重复")
		return
	}

	if handler, exists := eventMap[req.EventType]; exists {
		handler(req.Event)
		c.String(http.StatusOK, "OK")
		return
	} else {
		logrus.Error("Failed to find event handler: ", req)
		c.String(http.StatusBadRequest, "无对应处理函数")
		return
	}
}

func validateRequest(c *gin.Context, token string, rawBodyStr string) bool {
	// check the token and hash in event request header

	if token != config.C.Feishu.VerificationToken {
		return false
	}
	timestamp := c.Request.Header.Get("X-Lark-Request-Timestamp")
	nonce := c.Request.Header.Get("X-Lark-Request-Nonce")
	signature := c.Request.Header.Get("X-Lark-Signature")

	return signature == calculateSignature(timestamp, nonce, config.C.Feishu.EncryptKey, rawBodyStr)
}

func eventRepeatDetect(eventId string) bool {
	if _, repeated := eventIdList[eventId]; repeated {
		return true
	} else {
		eventIdList[eventId] = true
		return false
	}
}
