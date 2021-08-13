package utils

// 加密类
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

// md5验证
func MD5Str(src string) string {
	h := md5.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

// hmacsha256验证
func HmacSha256(src, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}

//生成随机长度的string类型
//0  // 纯数字
//1  // 小写字母
//2  // 大写字母
//3  // 数字、大小写字母
func Krand(size int, kind int) (str string) {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	str = string(result)
	return
}


func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)

	origData = PKCS5UnPadding(origData)

	return origData, nil
}
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
