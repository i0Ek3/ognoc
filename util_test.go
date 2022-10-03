package ognoc

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	o := fmt.Println
	oo := fmt.Printf
	oo("----> Plaintext = \"%s\"\n", Plaintext)

	t.Run("test caesar and transform", func(t *testing.T) {
		oo("----> Caesar(\"%s\", %d) = %s\n", Plaintext, Offset, Caesar(Plaintext, Offset))
		oo("----> Transform(\"%s\", %s, %d) = %s\n", Plaintext, AlgoType, PwdLen, Transform(Plaintext, AlgoType, PwdLen))
	})

	t.Run("test other", func(t *testing.T) {
		oo("----> CommonT(\"abcefg\", \"low\", %d) = %s\n", Offset, CommonT("abcefg", "low", Offset))
		oo("----> CommonT(\"abcefg\", \"up\", %d) = %s\n", Offset, CommonT("abcefg", "up", Offset))
		oo("----> CommonT(\"abcefg\", \"spec\", %d) = %s\n", Offset, CommonT("abcefg", "spec", Offset))
		oo("----> FormatString(\"a 3bd!!!!Ef@2RLLLK\") = %s\n", FormatString("a 3bd!!!!Ef@2RLLLK"))
		oo("----> FormatN(\"abcdef\", \"number\", \"pre\", %d) = %s\n", PwdLen, FormatN("abcdef", "number", "pre", PwdLen))
		oo("----> FormatN(\"abcdef\", \"number\", \"inner\", %d) = %s\n", PwdLen, FormatN("abcdef", "number", "inner", PwdLen))
		oo("----> FormatN(\"abcdef\", \"number\", \"post\", %d) = %s\n", PwdLen, FormatN("abcdef", "number", "post", PwdLen))
		oo("----> FormatN(\"abcdef\", \"spechar\", \"pre\", %d) = %s\n", PwdLen, FormatN("abcdef", "spechar", "pre", PwdLen))
		oo("----> FormatN(\"abcdef\", \"spechar\", \"inner\", %d) = %s\n", PwdLen, FormatN("abcdef", "spechar", "inner", PwdLen))
		oo("----> FormatN(\"abcdef\", \"spechar\", \"post\", %d) = %s\n", PwdLen, FormatN("abcdef", "spechar", "post", PwdLen))
		oo("----> FormatN(\"abcdef\", \"letter\", \"pre\", %d) = %s\n", PwdLen, FormatN("abcdef", "letter", "pre", PwdLen))
		oo("----> FormatN(\"abcdef\", \"letter\", \"inner\", %d) = %s\n", PwdLen, FormatN("abcdef", "letter", "inner", PwdLen))
		oo("----> FormatN(\"abcdef\", \"letter\", \"post\", %d) = %s\n", PwdLen, FormatN("abcdef", "letter", "post", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"number\", \"pre\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "number", "pre", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"number\", \"inner\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "number", "inner", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"number\", \"post\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "number", "post", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"spechar\", \"pre\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "spechar", "pre", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"spechar\", \"inner\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "spechar", "inner", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"spechar\", \"post\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "spechar", "post", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"letter\", \"pre\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "letter", "pre", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"letter\", \"inner\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "letter", "inner", PwdLen))
		oo("----> FormatN(\"abcdefghijkl\", \"letter\", \"post\", %d) = %s\n", PwdLen, FormatN("abcdefghijkl", "letter", "post", PwdLen))
	})

	t.Run("test GenerateRandom", func(t *testing.T) {
		o(GenerateRandom(Mixed))
		o(GenerateRandom(Alphabet))
	})
}
