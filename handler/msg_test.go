package handler

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"
	"wechat_api/model"
)

var receive = `
<xml>
 <ToUserName><![CDATA[公众号]]></ToUserName>
 <FromUserName><![CDATA[粉丝号]]></FromUserName>
 <CreateTime>1460537339</CreateTime>
 <MsgType><![CDATA[text]]></MsgType>
 <Content><![CDATA[欢迎开启公众号开发者模式]]></Content>
 <MsgId>6272960105994287618</MsgId>
</xml>`

func TestMsg(t *testing.T) {
	receiveMsg := &model.ReceiveMsg{}
	err := xml.Unmarshal([]byte(receive), receiveMsg)
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

	fmt.Printf(string(replyData))
}
