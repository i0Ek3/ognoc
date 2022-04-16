# ognoc

> That's really no tears without a coffin.

`ognoc` is a cryptosystem which offers you kinda ability to set strong password to protect your personal information.


## Why `ognoc`?

`ognoc` is a variation of Congo, they are anagram. I use it as the name of cryptosystem to create confusion, just like our password, while you can remember it, you can't guess its meaning. Of course, Ognoc more than that, it can do more, welcome to try it.


## Feature

- General Policy
  - general transformation
    - lowcase + upcase + special characters
  - two version support
    - only contain numbers and alphabets
    - contain special characters
  - Caesar cipher
  - SHA transformation
- Complicated Policy
  - Multiple transformation
- PQD


## Getting Started

### Usage

```console
» ./ognoc -h
Usage of ./ognoc:
  -c string
        the color of generate password(black, white, blue, red, yellow, green, cyan, magenta) (default "blue")
  -f string
        the base fill content(number, letter, spechar) (default "spechar")
  -l int
        the length of generated password (default 12)
  -n int
        the offset you want to move (default 2)
  -p string
        specific the position to insert(pre, inner, post) (default "inner")

» ./ognoc -f="number" -l=15 -n=4 -p="post" -c="magenta"
efghijklm919263
```


### Install

`$ go get https://github.com/i0Ek3/ognoc`


### Import

```Go
package main

import (
        "github.com/i0Ek3/ognoc"
)

func main() {
    pwd1 := ognoc.Caesar(plaintext, offset)
    pwd2 := ognoc.Transform(cipher, algoType, n...)

    fmt.Printf(pwd1)
    fmt.Printf(Detect(pwd1))
    fmt.Printf(pwd2)
    fmt.Printf(Detect(pwd2))

    fmt.Println(ognoc.Generate(plaintext, position, fill, color, offset, pwdLen))
}
```

More details please run command `go doc`.


## Password Quality Detection(PQD)

In ognoc, we use easy scoring mechanism to detect password quality, please check [this post](https://github.com/rlmao/pass-strength) to know more.


## Some Issues We Met Here

- syntax error: cannot use ... with non-final parameter ns | `func test(a int, b ...int) {}`
- cannot assign to string, cause of strings a immutable | `r := []rune(string)`
- cannot use type byte as type rune in assignment | `var r rune; var b byte; r = rune(b)`
- non-constant array bound length | `b := make([]byte, length)`


## Bugfix

- [ ] wierd CommonT() issue: sometime have right result sometime not
- [ ] fix invariable position in fillN function
- [ ] generation result continuity problem
- [ ] output the random length with default value 


## TODO

- [x] cmd and argument control support
- [ ] password strongability detection
  - [x] easy scoring mechanism
  - [ ] accroding entrophy to calculate the complicated of password
  - [ ] or use regex to detect the characters coverage
- [ ] output cipher randomly with given or ungiven string
  - [ ] do not use the result convert by caesar transform 
- [ ] support more general policy
  - [ ] keyword transform
- [ ] support more complicated policy
  - [ ] build personal transform algorithm: gogogo()
- [ ] generic support
- [ ] logger system
- [ ] access database to store cipher
  - [ ] MySQL/PostgreSQL


## Credit

- http://www.atoolbox.net/Category.php?Id=27

- http://www.metools.info/code/c28.html

- https://yqqy.top/regex-lookground-verify-passowrd/

- https://icode.best/i/24537132742865

- [https://github.com/rlmao/pass-strength](https://github.com/rlmao/pass-strength)
