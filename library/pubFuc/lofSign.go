package pubFuc

import (
	"crypto/sha256"
	"czc-server/library/logger"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckSign(secret, timestamp, nonce, sign string) bool {
	ts, _ := strconv.ParseInt(timestamp, 10, 64)
	// 如果请求头时间与当前时间相差超过180秒，则认为请求不合法 (结合x-rio-nonce，可以防止请求重放)
	now := time.Now().Unix()
	if ts > now+180 || ts < now-180 {
		logger.Print("time expired")
		//return false
	}
	signData := fmt.Sprintf("%s%s%s%s", timestamp, secret, nonce, timestamp)
	logger.Print("signdata: %s", signData)
	res := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(signData))))
	logger.Print("signres: %s", res)
	return res == sign
}

func GetSign(secret, timestamp, nonce string) string {
	signData := fmt.Sprintf("%s%s%s%s", timestamp, secret, nonce, timestamp)
	res := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(signData))))
	return res
}
