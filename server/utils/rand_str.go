package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/url"
	"sort"
	"strings"
)

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}

	return string(result), nil
}

// GenerateSignature 生成签名的函数
func GenerateSignature(params map[string]string, appSecret string) string {
	// 删除 signature 参数（因为不参与签名）
	delete(params, "signature")

	// 提取并按 ASCII 顺序排序参数名
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接成 key=value&key2=value2… 格式的字符串
	var stringA strings.Builder
	for i, k := range keys {
		stringA.WriteString(fmt.Sprintf("%s=%s", k, url.QueryEscape(params[k])))
		if i < len(keys)-1 {
			stringA.WriteString("&")
		}
	}

	// 在 stringA 末尾添加 appSecret
	stringSignTemp := stringA.String() + "&appSecret=" + appSecret

	// 进行 MD5 运算
	hash := md5.New()
	hash.Write([]byte(stringSignTemp))
	signature := strings.ToUpper(hex.EncodeToString(hash.Sum(nil))) // 转换为大写

	return signature
}
