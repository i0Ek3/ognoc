package ognoc

import (
	"fmt"
)

func show() {
	fmt.Printf("------------------------\n")
}

func main() {
	str := stringCheck(plaintext)
	pwd1 := Caeser(str, offset)

	// FIXME:
	pwd2 := Transform(pwd1, "x")

	// FIXME: lowUp/upLow return wrong result: single char
	pwd3 := CommonT(str, "spec")
	pwd4 := CommonT(str, "low")
	pwd5 := CommonT(str, "up")

	// FIXME: insert position fixed
	pwd6 := FormatN(pwd1, "number", "inner", pwdLen)
	pwd7 := FormatN(pwd1, "letter", "inner", pwdLen)
	pwd8 := FormatN(pwd1, "spechar", "inner", pwdLen)

	// FIXME: insert position fixed
	pwd9 := FormatN(pwd1, "number", "pre", pwdLen)
	pwd10 := FormatN(pwd1, "letter", "pre", pwdLen)
	pwd11 := FormatN(pwd1, "spechar", "pre", pwdLen)

	// FIXME: insert position fixed
	pwd12 := FormatN(pwd1, "number", "post", pwdLen)
	pwd13 := FormatN(pwd1, "letter", "post", pwdLen)
	pwd14 := FormatN(pwd1, "spechar", "post", pwdLen)

	// cutN
	pwd15 := FormatN("hahSDhashdfhasdfhasdfh", "number", "pre", pwdLen)
	pwd16 := FormatN("hahSDhashdfhasdfhasdfh", "number", "inner", pwdLen)
	pwd17 := FormatN("hahSDhashdfhasdfhasdfh", "number", "post", pwdLen)

	// fillN
	pwd18 := FormatN("hahSDh", "letter", "pre", pwdLen)
	pwd19 := FormatN("hahSDh", "letter", "inner", pwdLen)
	pwd20 := FormatN("hahSDh", "letter", "post", pwdLen)

	pwd21 := FormatN("hahSDh", "spechar", "pre", pwdLen)
	pwd22 := FormatN("hahSDh", "spechar", "inner", pwdLen)
	pwd23 := FormatN("hahSDh", "spechar", "post", pwdLen)

	pwd24 := FormatN("hahSDh", "number", "pre", pwdLen)
	pwd25 := FormatN("hahSDh", "number", "post", pwdLen)

	show()
	fmt.Println(str)
	show()
	fmt.Println(pwd1)
	show()
	fmt.Println(pwd2)
	show()
	fmt.Println(pwd3)
	show()
	fmt.Println(pwd4)
	show()
	fmt.Println(pwd5)
	show()
	fmt.Println(pwd6)
	fmt.Println(pwd7)
	fmt.Println(pwd8)
	show()
	fmt.Println(pwd9)
	fmt.Println(pwd10)
	fmt.Println(pwd11)
	show()
	fmt.Println(pwd12)
	fmt.Println(pwd13)
	fmt.Println(pwd14)
	show()
	fmt.Println(pwd15)
	fmt.Println(pwd16)
	fmt.Println(pwd17)
	show()
	fmt.Println(pwd18)
	fmt.Println(pwd19)
	fmt.Println(pwd20)
	show()
	fmt.Println(pwd21)
	fmt.Println(pwd22)
	fmt.Println(pwd23)
	show()
	fmt.Println(pwd24)
	fmt.Println(pwd25)

    // result:
    // ------------------------
    // abcdefghi
    // ------------------------
    // cdefghiab
    // ------------------------

    // ------------------------
    // abcdefghi@h
    // ------------------------
    // A
    // ------------------------
    // a
    // ------------------------
    // ------------------------
    // 1076742cdefghiab
    // prtnuzicdefghiab
    // >]{$+]<cdefghiab
    // ------------------------
    // cdefghiab0744375
    // cdefghiabjstdmlp
    // cdefghiab&#}@@&^
    // ------------------------
    // ashdfhasdfhasdfh
    // hahSDhashdfhasdf
    // ------------------------
    // tiomuykeqrhahSDh
    // hahSDhvzvjlmhrzp
    // ------------------------
    // _..)&!@)+#hahSDh
    // hahSDh?.<%&@}-+]
    // ------------------------
    // 4730216262hahSDh
    // hahSDh0247462390
}
