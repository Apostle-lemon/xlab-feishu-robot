package dispatcher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

func decrypt(encrypt string, key string) (string, error) {
	// copied from: https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-subscription-configure-/encrypt-key-encryption-configuration-case

	buf, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", fmt.Errorf("base64StdEncode Error[%v]", err)
	}
	if len(buf) < aes.BlockSize {
		return "", errors.New("cipher  too short")
	}
	keyBs := sha256.Sum256([]byte(key))
	block, err := aes.NewCipher(keyBs[:sha256.Size])
	if err != nil {
		return "", fmt.Errorf("AESNewCipher Error[%v]", err)
	}
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]
	// CBC mode always works in whole blocks.
	if len(buf)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(buf, buf)
	n := strings.Index(string(buf), "{")
	if n == -1 {
		n = 0
	}
	m := strings.LastIndex(string(buf), "}")
	if m == -1 {
		m = len(buf) - 1
	}
	return string(buf[n : m+1]), nil
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
