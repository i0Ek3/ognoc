package ognoc

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	o := fmt.Println
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
}
