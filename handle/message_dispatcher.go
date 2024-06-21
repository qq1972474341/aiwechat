package handle

import (
	"aiwechat/application/utils"
	"aiwechat/model"
	"github.com/eatmoreapple/openwechat"
	"github.com/gorilla/websocket"
)

// MessageDispatcher 分发处理前端发送的请求
func MessageDispatcher(ws *websocket.Conn, bot *openwechat.Bot, messageModel *model.RequestModel) {
	defer func() {
		if err := recover(); err != nil {
			utils.Logger.Errorln("请求分发处理失败", err)
		}
	}()
	//分发处理不同类型的消息
	switch messageModel.Operation {
	case model.SendMessage:
		SendMessage(bot, messageModel)
	case model.SendGroupMessage:
		SendGroupMessage(bot, messageModel)
	case model.GetFriendList:
		GetFriendList(ws, bot)
		GetFriendHeadImgList(ws, bot)
	case model.GetGroupList:
		GetGroupList(ws, bot)
		GetGroupHeadImgList(ws, bot)
	case model.GetHeadImg:
		GetHeadImg(ws, bot, messageModel)
	case model.GetGroupHeadImg:
		GetGroupHeadImg(ws, bot, messageModel)
	case model.AcceptFriendRequest:
		AcceptFriendRequest(ws, bot, messageModel)
	default:
		utils.Logger.Errorln("收到未定义的请求类型", messageModel)
	}
}

// ReceiveMessageAdapter 分发处理接收到的消息
func ReceiveMessageAdapter(ws *websocket.Conn, bot *openwechat.Bot, msg *openwechat.Message) {
	defer func() {
		if err := recover(); err != nil {
			utils.Logger.Errorln("消息接收分发处理失败", err)
		}
	}()
	if msg.IsSendByFriend() {
		ReceiveFriendMessage(ws, msg)
	} else if msg.IsSendByGroup() {
		ReceiveGroupMessage(ws, msg)
	} else if msg.IsFriendAdd() {
		utils.Logger.Infoln("收到未定义的消息类型", msg)
		ReceiveFriendAdd(ws, msg)
	}
}
