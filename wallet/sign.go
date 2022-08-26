package wallet

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
)

// Sign 签名
func Sign(data map[string]string) string {
	keys := make([]string, 0)
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := ""
	for _, key := range keys {
		if key != "sign" {
			str = str + data[key]
		}
	}
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// VerifySign 验证签名
func VerifySign(data map[string]string) bool {

	signNew := Sign(data)
	sign, ok := data["sign"]
	if !ok {
		return false
	}

	if sign != signNew {
		return false
	}

	return true
}
