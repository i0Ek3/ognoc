package main

import (
	"fmt"
	//"crypto"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

var (
	alphabets = "abcdefghijklmnopqrstuvwxyz"
	numbers   = "012345679"
	specials  = "!@#$%^&*()_+-=[]{}<>?.|"
)

const (
	plaintext = "aBcDefGhi"
	offset    = 2
	pwdLen    = 16
)

type F func(string, int) string

func Caeser(plaintext string, offset int) (res string) {
	strs := stringCheck(plaintext)
	if offsetCheck(strs, offset) {
		res = caeser(strs, offset)
	} else {
		offset = randomN(len(strs))
		res = caeser(strs, offset)
	}
	return
}

func caeser(plaintext string, offset int) string {
	if offset >= 0 {
		plaintext = plaintext[offset:len(plaintext)] + plaintext[:offset]
	} else {
		newn := len(plaintext) - int(math.Abs(float64(offset)))
		plaintext = plaintext[newn:len(plaintext)] + plaintext[:newn]
	}
	return plaintext
}

func Transform(cipher, algoType string, ns ...int) string {
	// TODO: getLen()
	//n := getN(len(Caeser(cipher, offset)), ns[0])
	n := randomN(len(Caeser(cipher, offset)))
	if offsetCheck(cipher, n) {
		return splitCipher(algoType, cipher, n)
	} else {
		return fmt.Sprintf("Invalid length to transform.")
	}
}

/*func getLen(algoType, cipher string, key ...string) (l int) {
    if len(key) > 0 {
        key = key[0]
    } else {
        key = ""
    }

	switch algoType {
	case "SHA1":
		l = len(crypto.SHA1(cipher)) - 1
	case "SHA224":
		l = len(crypto.SHA224(cipher)) - 1
	case "SHA256":
		l = len(crypto.SHA256(cipher)) - 1
	case "SHA384":
	    l = len(crypto.SHA384(cipher)) - 1
	case "SHA512":
		l = len(crypto.SHA512(cipher)) - 1
	case "SHA3":
		l = len(crypto.SHA3(cipher)) - 1
	case "MD516":
		l = len(crypto.MD5(cipher, 16)) - 1
	case "MD532":
		l = len(crypto.MD5(cipher, 32)) - 1
    case "HmacSHA1":
		l = len(crypto.HmacSHA1(cipher, key)) - 1
	case "HmacSHA224":
		l = len(crypto.HmacSHA224(cipher, key)) - 1
	case "HmacSHA256":
		l = len(crypto.HmacSHA256(cipher, key)) - 1
	case "HmacSHA384":
		l = len(crypto.HmacSHA384(cipher, key)) - 1
	case "HmacSHA512":
		l = len(crypto.HmacSHA512(cipher, key)) - 1
	case "HmacMD5":
		l = len(crypto.HmacMD5(cipher, key)) - 1
	}

    return
}*/

func randomN(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}

// TODO: write getN into T version for common use
func getN(length int, ns ...int) (n int) {
	if len(ns) > 0 {
		n = ns[0]
	} else {
		n = randomN(length)
	}
	return
}

func CommonT(plaintext, policy string, ns ...int) (res string) {
	strs := stringCheck(plaintext)

	//n := getN(len(strs), ns[0])
	n := randomN(len(strs))

	switch policy {
	case "low":
		res = lowUp(strs, n)
	case "up":
		res = upLow(strs, n)
	case "spec":
		res = special(strs, n)
	}
	return
}

func stringCheck(str string) string {
	return removeNonstr(str)
}

func removeNonstr(str string) string {
	var res string
	for i := 0; i < len(str); i++ {
		if unicode.IsLetter(rune(str[i])) {
			res += string(str[i])
		}
	}
	return strings.ToLower(res)
}

func offsetCheck(given string, offset int) bool {
	if int(math.Abs(float64(offset))) < len(given) {
		return true
	}
	return false
}

// TODO: rewrite splitCipher function with given crypto function
// splitCipher splits given algorithm result into new string,
// which length is n and Key is not requested
func splitCipher(algoType, cipher string, n int) (res string) {
	// FIXME: hardcode
	if algoType != "" {
		return cipher[:n]
	}
	return ""
	/*switch algoType {
	case "Hash":
		res = crypto.Hash(cipher)[:n]
	case "SHA1":
		res = crypto.SHA1(cipher)[:n]
	case "SHA224":
		res = crypto.SHA224(cipher)[:n]
	case "SHA256":
		res = crypto.SHA256(cipher)[:n]
	case "SHA3":
		res = crypto.SHA3(cipher)[:n]
	case "SHA3_224":
		res = crypto.SHA3_224(cipher)[:n]
	case "SHA3_256":
		res = crypto.SHA3_256(cipher)[:n]
	case "SHA384":
		res = crypto.SHA384(cipher)[:n]
	case "SHA512":
		res = crypto.SHA512(cipher)[:n]
	case "SHA512_224":
		res = crypto.SHA512_224(cipher)[:n]
	case "SHA512_256":
		res = crypto.SHA512_256(cipher)[:n]
	case "MD516":
		res = crypto.MD516(cipher)[:n]
	case "MD532":
		res = crypto.MD532(cipher)[:n]
	}
	return*/
}

// Key version of splitCipher
// TODO: need to check whether crypto support Hmac type algorithm
/*func splitCipherWithKey(algoType, cipher, key string, n int) (res string) {
	switch algoType {
	case "HmacSHA1":
		res = crypto.HmacSHA1(cipher, key)[:n]
	case "HmacSHA224":
		res = crypto.HmacSHA224(cipher, key)[:n]
	case "HmacSHA256":
		res = crypto.HmacSHA256(cipher, key)[:n]
	case "HmacSHA384":
		res = crypto.HmacSHA384(cipher, key)[:n]
	case "HmacSHA512":
		res = crypto.HmacSHA512(cipher, key)[:n]
	case "HmacMD5":
		res = crypto.HmacMD5(cipher, key)[:n]
	}
	return
}*/

// n means offset, while n = 2, abcdefg => aBcDeFg,
// while n = 3, abcdefg => abCdeFg
func lowUp(str string, offset int) string {
	for i := 0; i < len(str); i += offset {
		str = strings.ToUpper(string(str[i]))
	}
	return str
}

// upLow just the other side compare to lowUp
// n = 2, abcdefg => AbCdEfG, n = 3, abcdefg => AbcDefG
func upLow(str string, offset int) string {
	for i := 0; i < len(str); i += offset {
		str = strings.ToLower(string(str[i]))
	}
	return str
}

// n means how many random special characters in cipher
func special(str string, offset int) string {
	for i := 0; i < len(str); i += offset {
		str += specialize2Spechar(rune(str[i]))
	}
	return str
}

// FormatN formats given string into a fixed string
// which length is n, recommend n between [10, 14]
func FormatN(str, content, where string, n int) (newstr string) {
	if len(str) < n {
		newstr = fillN(str, content, where, n-len(str))
	} else if len(str) > n {
		newstr = cutN(str, where, len(str)-n)
	} else {
		newstr = str
	}
	return
}

func fillN(str, content, where string, n int) string {
	// fill numbers, chars and characters
	// TODO: only characters? cause of formatN()
	if content == "number" {
		return fillNumber(str, where, n)
	} else if content == "spechar" {
		return fillSpechar(str, where, n)
	} else if content == "letter" {
		return fillLetter(str, where, n)
	} else {
		return fmt.Sprintf("Unexpected content!")
	}
}

/*
// fill position: pre, post, in, use an algorithm to choose
// which position should be take in fillX function
func selectOne(n int) {}

// TODO: abstract where function
func where(str, where string) F {}

*/

// TODO: hardcode for now, change later
// the specified content is not accepted temporarily
func fillNumber(str, where string, n int) string {
	var strs string
	switch where {
	case "pre":
		for i := 0; i < n; i++ {
			str = string(numbers[randomN(len(numbers))]) + str
		}
	case "inner":
		for i := 0; i < n; i++ {
			strs += string(numbers[randomN(len(numbers))])
		}
		//str = str[:1] + strs + str[n+1:len(str)-n-1]
		str = str[:1] + strs + str[n+1:]
	case "post":
		for i := 0; i < n; i++ {
			str += string(numbers[randomN(len(numbers))])
		}
	}
	return str
}

func fillSpechar(str, where string, n int) string {
	var strs string
	switch where {
	case "pre":
		for i := 0; i < n; i++ {
			str = string(specials[randomN(len(specials))]) + str
		}
	case "inner":
		for i := 0; i < n; i++ {
			strs += string(specials[randomN(len(specials))])
		}
		//str = str[:1] + strs + str[n+1:len(str)-n-1]
		str = str[:1] + strs + str[n+1:]
	case "post":
		for i := 0; i < n; i++ {
			str += string(specials[randomN(len(specials))])
		}
	}
	return str
}

func fillLetter(str, where string, n int) string {
	var strs string
	switch where {
	case "pre":
		for i := 0; i < n; i++ {
			str = string(alphabets[randomN(len(alphabets))]) + str
		}
	case "inner":
		for i := 0; i < n; i++ {
			strs += string(alphabets[randomN(len(alphabets))])

		}
		//str = str[:1] + strs + str[n+1:len(str)-n-1]
		str = str[:1] + strs + str[n+1:]
	case "post":
		for i := 0; i < n; i++ {
			str += string(alphabets[randomN(len(alphabets))])
		}
	}
	return str
}

func cutN(str, where string, n int) string {
	switch where {
	case "pre":
		str = str[n:]
	case "inner":
		str = str[:n] + str[n*2:]
	case "post":
		str = str[:len(str)-n]
	}
	return str
}

// TODO: these two functions need to transform []string into
// []rune and then transform back to []string instead of rune 2 rune
// specialize2Rune specializes the like-char characters into
// special characters, for example i,j => !, a => @
func specialize2Spechar(r rune) string {
	// TODO: func specialize2Rune(str string) string {
	// considerate transform rune and string each other
	switch r {
	case 'a', 'c':
		return "@"
	case 'i', 'j':
		return "!"
	case 'p':
		return "+"
	case 'x':
		return "*"
	default:
		return string(r)
	}
}

// specialize2Num2Num specializes the like-char characters into
// a number, for example i,j => 1, o => 0
func specialize2Num(r rune) string {
	switch r {
	case 'b':
		return "6"
	case 'e', 'i', 'j', 'l':
		return "1"
	case 'g', 'q':
		return "9"
	case 'o':
		return "0"
	case 'r', 'z':
		return "2"
	case 's':
		return "5"
	default:
		return string(r)
	}
}
