package Ognoc

import (
    "fmt"
)

const (
    MSGMISSING = "message missing"
)

func Error(msg string) {
    fmt.Println(msg)
}
