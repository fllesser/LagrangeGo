package logic

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/entity"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strings"
)

func Sgst(client *client.QQClient, event *message.GroupMessage) {
	msg := event.ToString()
	if strings.HasPrefix(msg, "sgst") {
		memberInfo := client.GetCachedMemberInfo(client.Uin, event.GroupUin)
		if memberInfo.Permission == entity.Owner {
			title := strings.TrimSpace(strings.TrimPrefix(msg, "sgst"))
			_ = client.GroupSetSpecialTitle(event.GroupUin, event.Sender.Uin, title)
		}
	}
}
