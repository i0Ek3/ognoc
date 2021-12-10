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



## Getting Started

### Install

`$ go get https://github.com/i0Ek3/ognoc`

### Usage

```Go
package main

import (
		"github.com/i0Ek3/ognoc"
)

func main() {
    pwd1 := ognoc.Caesar(plaintext, offset)
    pwd2 := ognoc.Transform(cipher, algoType, n...)

    fmt.Printf(pwd1)
    fmt.Printf(pwd2)
}
```

More details please run command `go doc`.

## Some Issues We Met Here

- syntax error: cannot use ... with non-final parameter ns | `func test(a int, b ...int) {}`
- cannot assign to string, cause of strings a immutable | `r := []rune(string)`
- cannot use type byte as type rune in assignment | `var r rune; var b byte; r = rune(b)`
- non-constant array bound length | `b := make([]byte, length)`


## TODO

- [ ] bug fix
    - [ ] wierd CommonT() issue: sometime have right result sometime not
    - [ ] fix invariable position in fillN function
    - [ ] generation result continuity problem
- [ ] more general policy
    - [ ] keyword transform
- [ ] generic support
- [ ] password strongability detection
    - [ ] accroding entrophy to calculate the comlicated of password
    - [ ] or use regex to detect the characters coverage
- [ ] logger system
- [ ] cmd support
- [ ] access database to store cipher
    - [ ] MySQL
    - [ ] PostgreSQL
- [ ] comlicated policy
    - [ ] build personal transform algorithm: gogogo()


## Credit

- http://www.atoolbox.net/Category.php?Id=27
- http://www.metools.info/code/c28.html
- https://yqqy.top/regex-lookground-verify-passowrd/
- https://icode.best/i/24537132742865
