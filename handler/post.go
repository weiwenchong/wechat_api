package handler

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
	"wechat_api/model"
)

func Post(c *gin.Context) {
	fmt.Printf("Post call")

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Post read body err:%v", err)
		return
	}
	fmt.Printf("Post data:%v", data)

	receiveMsg := &model.ReceiveMsg{}
	err = xml.Unmarshal(data, receiveMsg)
	if err != nil {
		fmt.Printf("Post unmarshal msg err:%v", err)
		return
	}

	fmt.Printf("Post receive msg :%v", receiveMsg)

	replyMsg := &model.ReplyMsg{
		ToUserName:   receiveMsg.FromUserName,
		FromUserName: receiveMsg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      receiveMsg.Content,
	}

	replyData, err := xml.Marshal(replyMsg)
	if err != nil {
		fmt.Printf("Post Marshal err:%v", err)
		return
	}
	c.String(200, string(replyData))
}
