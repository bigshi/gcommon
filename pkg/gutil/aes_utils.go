package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// =================== AES/ECB/PKCS5Padding ==================================================================
// =================== AES/ECB/PKCS5Padding ==================================================================
// =================== AES/ECB/PKCS5Padding ==================================================================
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	paddingNum := int(src[length-1])
	return src[:length-paddingNum]
}

// aes 加密
func AesEncryptOfECBWithPKCS5Padding(key []byte, origData []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//padding
	origData = pkcs5Padding(origData, block.BlockSize())
	//存储每次加密的数据
	//分组分块加密
	buffer := bytes.NewBufferString("")
	tmpData := make([]byte, block.BlockSize()) //存储每次加密的数据
	for index := 0; index < len(origData); index += block.BlockSize() {
		block.Encrypt(tmpData, origData[index:index+block.BlockSize()])
		buffer.Write(tmpData)
	}
	return buffer.Bytes(), nil
}

// aes解密
func AesDecryptOfECBWithPKCS5Padding(key []byte, origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBufferString("")
	tmpData := make([]byte, block.BlockSize())
	for index := 0; index < len(origData); index += block.BlockSize() {
		block.Decrypt(tmpData, origData[index:index+block.BlockSize()])
		buffer.Write(tmpData)
	}
	return pkcs5UnPadding(buffer.Bytes()), nil
}

// =================== AES/GCM/NoPadding ==================================================================
// =================== AES/GCM/NoPadding ==================================================================
// =================== AES/GCM/NoPadding ==================================================================
// aes 加密
func AesEncryptOfGCMWithNoPadding(key []byte, nonce []byte, origData []byte) ([]byte, error) {
	// key只能是 16 24 32长度
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 补全
	gcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}
	return gcm.Seal(nil, nonce, origData, nil), nil
}

// aes解密
func AesDecryptOfGCMWithNoPadding(key []byte, nonce []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCMWithNonceSize(block, 16)
	if err != nil {
		return nil, err
	}
	src, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return src, nil
}
