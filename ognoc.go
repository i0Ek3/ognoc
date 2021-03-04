package Ognoc

import (
    u "Ognoc/utils"
    e "Ognoc/error"
)

// const defines some numbers
const (
    // iterTimes is default times of iteration
    iterTimes = 100
    // outputLen is default length of output string
    outputLen = 8
)

// Ognoc is the main struct of Ognoc
type Ognoc struct {
    // openned indicates if random channel enable
    openned bool
    // size indicates the size of string
    size    int
}

// Ognocer defines the fuction of Ognoc
type Ognocer interface {
    // generate generates the string with given string
    generate(given string) string
    // randomit shuffles the original string
    randomit(given string) string
}

func (o *Ognoc) init() {
    o.openned = true
    o.size = 8
}

func (o *Ognoc) generate(given string) string {
    newstring := ""
    if u.Empty(given) {
        return e.MSGMISSING
    } else {
        if o.openned {
            o.randomit(given)
        } else {
            // TODO: implement generate algorithm
            newstring += given
        }
    }
    return newstring
}

func (o *Ognoc) randomit(given string) string {
    // TODO: random shuffle
    return ""
}
