package gutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSAEncrypt
//
//	@Description: RSA加密
//	@param plainTextBytes
//	@param publicKeyBytes
//	@return []byte
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

// RSADecrypt
//
//	@Description: RSA解密
//	@param cipherTextBytes
//	@param privateKeyBytes
//	@return []byte
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

// RSAGenerate
//
//	@Description: 密钥生成
//	@param bits
//	@return []byte
//	@return []byte
//	@return error
func RSAGenerate(bits int) ([]byte, []byte, error) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	prk := pem.EncodeToMemory(&privateBlock)
	if prk == nil {
		return nil, nil, errors.New("generate private key fail")
	}
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, nil, err
	}
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	puk := pem.EncodeToMemory(&publicBlock)
	if prk == nil {
		return nil, nil, errors.New("generate public key fail")
	}
	return prk, puk, nil
}
