package ognoc

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	fmt.Printf("----> plaintext = %s\n", plaintext)
	fmt.Printf("----> Caesar(%s, %d) = %s\n", plaintext, offset, Caesar(plaintext, 2))
	fmt.Printf("----> Transform(%s, 'x', %d) = %s\n", plaintext, pwdLen, Transform(plaintext, "x", pwdLen))
	fmt.Printf("----> CommonT(%s, 'low', %d) = %s\n", plaintext, pwdLen, CommonT(plaintext, "low", pwdLen))
	fmt.Printf("----> CommonT(%s, 'up', %d) = %s\n", plaintext, pwdLen, CommonT(plaintext, "up", pwdLen))
	fmt.Printf("----> CommonT(%s, 'spec', %d) = %s\n", plaintext, pwdLen, CommonT(plaintext, "spec", pwdLen))
	fmt.Printf("----> CommonT(%s, 'low', %d) = %s\n", plaintext, pwdLen, CommonT(plaintext, "low", pwdLen))
	fmt.Printf("----> StringCheck('a 3bd!!!!Ef@2RLLLK') = %s\n", StringCheck("a 3bd!!!!Ef@2RLLLK"))
	fmt.Printf("----> FormatN('abcdef', 'number', 'pre', %d) = %s\n", pwdLen, FormatN("abcdef", "number", "pre", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'number', 'inner', %d) = %s\n", pwdLen, FormatN("abcdef", "number", "inner", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'number', 'post', %d) = %s\n", pwdLen, FormatN("abcdef", "number", "post", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'spechar', 'pre', %d) = %s\n", pwdLen, FormatN("abcdef", "spechar", "pre", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'spechar', 'inner', %d) = %s\n", pwdLen, FormatN("abcdef", "spechar", "inner", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'spechar', 'post', %d) = %s\n", pwdLen, FormatN("abcdef", "spechar", "post", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'letter', 'pre', %d) = %s\n", pwdLen, FormatN("abcdef", "letter", "pre", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'letter', 'inner', %d) = %s\n", pwdLen, FormatN("abcdef", "letter", "inner", pwdLen))
	fmt.Printf("----> FormatN('abcdef', 'letter', 'post', %d) = %s\n", pwdLen, FormatN("abcdef", "letter", "post", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'number', 'pre', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "number", "pre", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'number', 'inner', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "number", "inner", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'number', 'post', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "number", "post", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'spechar', 'pre', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "spechar", "pre", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'spechar', 'inner', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "spechar", "inner", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'spechar', 'post', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "spechar", "post", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'letter', 'pre', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "letter", "pre", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'letter', 'inner', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "letter", "inner", pwdLen))
	fmt.Printf("----> FormatN('abcdefghijkl', 'letter', 'post', %d) = %s\n", pwdLen, FormatN("abcdefghijkl", "letter", "post", pwdLen))
}
