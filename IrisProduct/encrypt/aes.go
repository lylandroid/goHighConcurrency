package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/kataras/iris/core/errors"
)

//高级加密标准（Adevanced Encryption Standard）

//16,24,32位字符串的话，分别对应AES-128,AES-192,AES-256 加密方法
//key很重要，不能泄露
var PwdKey = []byte("DIS**#KKKDJJSKDI")

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	//Repeat()函数的功能是把切换[]byte{byte(padding)复制padding个，
	// 然后合并成新的字节切片返回
	paddingTxt := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingTxt...)
}

//填充反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充字符串长度
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)], nil
}

//加密
func AesEncrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取快大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypt := make([]byte, len(origData))
	//执行加密
	blockMode.CryptBlocks(encrypt, origData)
	return encrypt, nil
}

//解密
func AesDecrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取快大小
	blockSize := block.BlockSize()
	//对数据进行填充，让数据长度满足需求
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypt := make([]byte, len(origData))
	blockMode.CryptBlocks(decrypt, origData)
	//执行加密
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return decrypt, nil
}

//加密base64
func EnPwdCode(pwd string) (string, error) {
	result, err := AesEncrypt([]byte(pwd), PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

//解密base64
func DePwdCode(pwd string) ([]byte, error) {
	//解密Base64字符串
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(pwdByte, PwdKey)
}
