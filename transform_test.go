package ognoc

import (
	"fmt"
	"testing"
)

func TestTransform(t *testing.T) {
	fmt.Printf("----> plaintext = %s\n", plaintext)
	fmt.Printf("----> Caesar(%s, %d) = %s\n", plaintext, offset, Caesar(plaintext, 2))
	fmt.Printf("----> Transform(%s, %s, %d) = %s\n", plaintext, algoType, pwdLen, Transform(plaintext, algoType, pwdLen))
	fmt.Printf("----> CommonT(abcefgipxors, 'low', %d) = %s\n", offset, CommonT("abcefgipxors", "low", offset))
	fmt.Printf("----> CommonT(abcefgipxors, 'up', %d) = %s\n", offset, CommonT("abcefgipxors", "up", offset))
	fmt.Printf("----> CommonT(abcefgipxors, 'spec', %d) = %s\n", offset, CommonT("abcefgipxors", "spec", offset))
	fmt.Printf("----> FormatString('a 3bd!!!!Ef@2RLLLK') = %s\n", FormatString("a 3bd!!!!Ef@2RLLLK"))
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
