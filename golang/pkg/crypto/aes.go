package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Aes存储秘钥
const (
	NoEncrypt = 0
	AesCbc    = 1
)

// crypto 加密算法
func StorageEncrypt(srcStr []byte, crypto int, signKey []byte) ([]byte, error) {
	switch crypto {
	case AesCbc:
		return cbcEncrypt(srcStr, signKey, iV)
	default:
		return srcStr, nil
	}
}

func StorageDecrypt(encrypt []byte, crypto int, signKey []byte) ([]byte, error) {
	switch crypto {
	case AesCbc:
		return cbcDecrypt(encrypt, signKey, iV)
	default:
		return encrypt, nil
	}
}

func CommonEncrypt(srcStr []byte, crypto int, key []byte) ([]byte, error) {
	switch crypto {
	case AesCbc:
		return cbcEncrypt(srcStr, key, iV)
	default:
		return srcStr, nil
	}
}

func CommonDecrypt(encryptStr []byte, crypto int, key []byte) ([]byte, error) {
	switch crypto {
	case AesCbc:
		return cbcDecrypt(encryptStr, key, iV)
	default:
		return encryptStr, nil
	}
}

// 基础向量
var iV = []byte{
	0x19,
	0x34,
	0x57,
	0x62,
	0x90,
	0xAB,
	0xCB,
	0xEF,
	0x12,
	0x64,
	0x14,
	0x78,
	0x90,
	0xAC,
	0xAE,
	0x45,
}

func cbcEncrypt(originText []byte, key []byte, iv []byte) ([]byte, error) {
	text := make([]byte, len(originText))
	copy(text, originText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	padText := pKCS7Padding(text, block.BlockSize()) // 填充
	blockMode := cipher.NewCBCEncrypter(block, iv)
	result := make([]byte, len(padText)) // 加密
	blockMode.CryptBlocks(result, padText)
	// base64转为字符
	return base64Encode(result), nil
}

func base64Encode(src []byte) []byte {
	enc := base64.StdEncoding
	buf := make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(buf, src)
	return buf
}

func base64Decode(s []byte) ([]byte, error) {
	enc := base64.StdEncoding
	dbuf := make([]byte, enc.DecodedLen(len(s)))
	n, err := enc.Decode(dbuf, s)
	return dbuf[:n], err
}

func cbcDecrypt(encryptStr []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// base64转为二进制
	encrypt, err := base64Decode(encryptStr)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(encrypt))
	blockMode.CryptBlocks(result, encrypt)
	// 去除填充
	result = unPKCS7Padding(result)
	return result, nil
}

// PKCS7Padding 计算待填充的长度
func pKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		// 已对齐，填充一整块数据，每个数据为 blockSize
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		// 未对齐 填充 padding 个数据，每个数据为 padding
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}

// UnPKCS7Padding 取出填充的数据 以此来获得填充数据长度
func unPKCS7Padding(text []byte) []byte {
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}
