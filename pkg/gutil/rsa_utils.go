package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

//
// RSAEncrypt
//  @Description: RSA加密
//  @param plainTextBytes
//  @param publicKeyBytes
//  @return []byte
//
func RSAEncrypt(plainTextBytes []byte, publicKeyBytes []byte) []byte {
	//pem解码
	block, _ := pem.Decode(publicKeyBytes)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherTextBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainTextBytes)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherTextBytes
}

//
// RSADecrypt
//  @Description: RSA解密
//  @param cipherTextBytes
//  @param privateKeyBytes
//  @return []byte
//
func RSADecrypt(cipherTextBytes []byte, privateKeyBytes []byte) []byte {
	//pem解码
	block, _ := pem.Decode(privateKeyBytes)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainTextBytes, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherTextBytes)
	//返回明文
	return plainTextBytes
}
