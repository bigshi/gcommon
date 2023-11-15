/**
 * Create Time:2023/11/14
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gutil

import (
	"math/rand"
	"time"
)

var numbers = []rune("01234567890")
var numbersAndAlphabet = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandNumberStr
//
//	@Description: 随机数字字符串
//	@param n
//	@return string
func RandNumberStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	length := len(numbers)
	for i := range b {
		b[i] = numbers[rand.Intn(length)]
	}
	return string(b)
}

// RandNumberAndAlphabetStr
//
//	@Description: 随机数字/字母字符串
//	@param n
//	@return string
func RandNumberAndAlphabetStr(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	length := len(numbersAndAlphabet)
	for i := range b {
		b[i] = numbersAndAlphabet[rand.Intn(length)]
	}
	return string(b)
}
