package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func main() {
	r := gin.Default()
	r.GET("/wx", wx)

	r.Run("0.0.0.0:80")
}

func wx(c *gin.Context) {
	fmt.Printf("wx call")

	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	fmt.Printf("wx url %s %s %s %s", signature, timestamp, nonce, echostr)

	if signature == "" || timestamp == "" || nonce == "" || echostr == "" {
		return
	}

	token := "7777"

	s := []string{token, timestamp, nonce}
	sort.Strings(s)
	ss := s[0] + s[1] + s[2]

	sha := sha1.New()
	sha.Write([]byte(ss))
	bs := sha.Sum(nil)

	bss := string(bs)

	fmt.Printf("bs:%s bss:%s", bs, bss)

	if bss == signature {
		c.String(http.StatusOK, echostr)
	}
}
