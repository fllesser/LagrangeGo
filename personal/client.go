package personal

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/sirupsen/logrus"
	"time"
)

func CheckAlive(c *client.QQClient) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logrus.Errorf("定时器发生错误，%v", r)
			}
			ticker.Stop() // 意外退出时关闭定时器
		}()
		status, lastStatus := c.Online.Load(), false
		statusContent := map[bool]string{true: "online", false: "offline"}
		Logger.Infof("Lgr -> [%v] %v", c.Uin, statusContent[status])
		for range ticker.C {
			lastStatus, status = status, c.Online.Load()
			if lastStatus != status {
				Logger.Infof("Lgr -> [%v] %v", c.Uin, statusContent[status])
			}
		}
	}()
}

func groupLog(_ *client.QQClient, event *message.GroupMessage) {
	Logger.Infof(fromProtocol+"message.group[gid:%v,uid:%v,msg:%v]", event.GroupUin, event.Sender.Uin, event.ToString())
}

func privateLog(_ *client.QQClient, event *message.PrivateMessage) {
	Logger.Infof(fromProtocol+"message.private[uid:%v,msg:%v]", event.Sender.Uin, event.ToString())
}
