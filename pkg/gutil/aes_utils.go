package gutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

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

//
// AesEncryptOfECBWithPKCS5Padding
//  @Description: aes 加密
//  @param key
//  @param origData
//  @return []byte
//  @return error
//
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

//
// AesDecryptOfECBWithPKCS5Padding
//  @Description: aes解密
//  @param key
//  @param origData
//  @return []byte
//  @return error
//
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

//
// AesEncryptOfGCMWithNoPadding
//  @Description: aes 加密
//  @param key
//  @param nonce
//  @param origData
//  @return []byte
//  @return error
//
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

//
// AesDecryptOfGCMWithNoPadding
//  @Description: aes解密
//  @param key
//  @param nonce
//  @param ciphertext
//  @return []byte
//  @return error
//
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
