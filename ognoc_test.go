package main

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	o := fmt.Println
	oo := fmt.Printf
	oo("----> plaintext = \"%s\"\n", plaintext)

	t.Run("test caesar and transform", func(t *testing.T) {
		oo("----> Caesar(\"%s\", %d) = %s\n", plaintext, offset, Caesar(plaintext, offset))
		oo("----> Transform(\"%s\", %s, %d) = %s\n", plaintext, algoType, pwdLen, Transform(plaintext, algoType, pwdLen))
	})

	t.Run("test other", func(t *testing.T) {
		oo("----> CommonT(\"abcefg\", \"low\", %d) = %s\n", offset, CommonT("abcefg", "low", offset))
		oo("----> CommonT(\"abcefg\", \"up\", %d) = %s\n", offset, CommonT("abcefg", "up", offset))
		oo("----> CommonT(\"abcefg\", \"spec\", %d) = %s\n", offset, CommonT("abcefg", "spec", offset))
		oo("----> FormatString(\"a 3bd!!!!Ef@2RLLLK\") = %s\n", FormatString("a 3bd!!!!Ef@2RLLLK"))
		oo("----> FormatN(\"abcdef\", \"number\", \"pre\", %d) = %s\n", pwdLen, FormatN("abcdef", "number", "pre", pwdLen))
		oo("----> FormatN(\"abcdef\", \"number\", \"inner\", %d) = %s\n", pwdLen, FormatN("abcdef", "number", "inner", pwdLen))
		oo("----> FormatN(\"abcdef\", \"number\", \"post\", %d) = %s\n", pwdLen, FormatN("abcdef", "number", "post", pwdLen))
		oo("----> FormatN(\"abcdef\", \"spechar\", \"pre\", %d) = %s\n", pwdLen, FormatN("abcdef", "spechar", "pre", pwdLen))
		oo("----> FormatN(\"abcdef\", \"spechar\", \"inner\", %d) = %s\n", pwdLen, FormatN("abcdef", "spechar", "inner", pwdLen))
		oo("----> FormatN(\"abcdef\", \"spechar\", \"post\", %d) = %s\n", pwdLen, FormatN("abcdef", "spechar", "post", pwdLen))
		oo("----> FormatN(\"abcdef\", \"letter\", \"pre\", %d) = %s\n", pwdLen, FormatN("abcdef", "letter", "pre", pwdLen))
		oo("----> FormatN(\"abcdef\", \"letter\", \"inner\", %d) = %s\n", pwdLen, FormatN("abcdef", "letter", "inner", pwdLen))
		oo("----> FormatN(\"abcdef\", \"letter\", \"post\", %d) = %s\n", pwdLen, FormatN("abcdef", "letter", "post", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"number\", \"pre\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "number", "pre", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"number\", \"inner\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "number", "inner", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"number\", \"post\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "number", "post", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"spechar\", \"pre\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "spechar", "pre", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"spechar\", \"inner\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "spechar", "inner", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"spechar\", \"post\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "spechar", "post", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"letter\", \"pre\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "letter", "pre", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"letter\", \"inner\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "letter", "inner", pwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"letter\", \"post\", %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "letter", "post", pwdLen))
	})

	t.Run("test cipher complicated", func(t *testing.T) {
		o(Detect("Ab@12sdf$0#*b)"))
		o(Detect("aaaaaaaaaa"))
		o(Detect("abncdd"))
		o(Detect("AAAAA"))
		o(Detect("ADSFJKDFSJSJDF"))
		o(Detect("1231"))
		o(Detect("123545613467417"))
		o(Detect("@%$%#$&&"))
		o(Detect("*##*"))
		o(Detect("asdfasdf12314568"))
		o(Detect("asdfASDdf12314568"))
		o(Detect("21LSDKF"))
		o(Detect("21lsdkf"))
		o(Detect("21LSDKF**&"))
	})

    t.Run("test GenerateRandom", func(t *testing.T) {
        o(GenerateRandom(mixed))
        o(GenerateRandom(alphabets))
	})
}
