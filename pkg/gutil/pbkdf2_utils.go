package gutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

func Key(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}

func Xor(k1 []byte, k2 []byte) []byte {
	l1 := len(k1)
	l2 := len(k2)
	l := l1
	if l < l2 {
		l = l2
	}
	rlt := make([]byte, l)
	for i := 0; i < l; i++ {
		var a byte = 0
		if i < l1 {
			a = k1[i]
		}
		var b byte = 0
		if i < l2 {
			b = k2[i]
		}
		rlt[i] = a ^ b
	}
	return rlt
}

func EncryptOfPbkdf2(key []byte, salt []byte, iterations int, length int) string {
	rootKeyByte := Key(key, salt, iterations, length, sha256.New)
	return hex.EncodeToString(rootKeyByte)
}
