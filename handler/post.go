package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
	"wechat_api/log"
	"wechat_api/model"
	"wechat_api/service"
)

func Post(c *gin.Context) {
	log.Infof("Post call")

	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Post err:%+v", err)
			c.String(200, "success")
		}
	}()

	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Errorf("Post read body err:%v", err)
		return
	}

	receiveMsg := &model.ReceiveMsg{}
	err = xml.Unmarshal(data, receiveMsg)
	if err != nil {
		log.Errorf("Post unmarshal msg err:%v", err)
		return
	}
	log.Infof("Post receive msg :%+v", receiveMsg)
	log.Infof("Post receive msg content :%s", receiveMsg.Content)

	//receiveMsg.Content = strings.Trim(receiveMsg.Content, "[")
	//receiveMsg.Content = strings.Trim(receiveMsg.Content, "]")
	rs := []rune(receiveMsg.Content)
	rsE := make([]rune, 0)
	for i, r := range rs {
		if i == 0 || i == len(rs)-1 {
			continue
		}
		rsE = append(rsE, r)
	}

	receiveMsg.Content = string(rsE)

	log.Infof("Post receive msg content :%s", receiveMsg.Content)

	// todo msgid去重
	// todo 默认返回success

	msg, err := service.ReplyMsg(receiveMsg)
	if err != nil {
		msg = err.Error()
	}

	replyMsg := &model.ReplyMsg{
		ToUserName:   receiveMsg.FromUserName,
		FromUserName: receiveMsg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      msg,
	}

	replyData, err := xml.Marshal(replyMsg)
	if err != nil {
		log.Errorf("Post Marshal err:%v", err)
		return
	}
	c.String(200, string(replyData))
}
