package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

const token = "7777"

func Auth(c *gin.Context) {
	fmt.Printf("wx call")

	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	if signature == "" || timestamp == "" || nonce == "" || echostr == "" {
		return
	}

	s := []string{token, timestamp, nonce}
	sort.Strings(s)
	ss := s[0] + s[1] + s[2]

	sha := sha1.New()
	sha.Write([]byte(ss))
	bs := sha.Sum(nil)

	bss := hex.EncodeToString(bs)

	if bss == signature {
		c.String(http.StatusOK, echostr)
	}
}
