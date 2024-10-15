package personal

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/entity"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strconv"
	"strings"
)

var emojis = make(map[uint32]bool, 200)

func init() {
	emojiArr := [167]uint32{4, 5, 8, 9, 10, 12, 14, 16, 21, 23, 24, 25, 26, 27, 28, 29, 30, 32, 33, 34,
		38, 39, 41, 42, 43, 49, 53, 60, 63, 66, 74, 75, 76, 78, 79, 85, 89, 96, 97, 98, 99, 100, 101,
		102, 103, 104, 106, 109, 111, 116, 118, 120, 122, 123, 124, 125, 129, 144, 147, 171, 173, 174,
		175, 176, 179, 180, 181, 182, 183, 201, 203, 212, 214, 219, 222, 227, 232, 240, 243, 246, 262, 264,
		265, 266, 267, 268, 269, 270, 271, 272, 273, 277, 278, 281, 282, 284, 285, 287, 289, 290, 293, 294,
		297, 298, 299, 305, 306, 307, 314, 315, 318, 319, 320, 322, 324, 326, 9728, 9749, 9786, 10024, 10060,
		10068, 127801, 127817, 127822, 127827, 127836, 127838, 127847, 127866, 127867, 127881, 128027, 128046,
		128051, 128053, 128074, 128076, 128077, 128079, 128089, 128102, 128104, 128147, 128157, 128164, 128166,
		128168, 128170, 128235, 128293, 128513, 128514, 128516, 128522, 128524, 128527, 128530, 128531, 128532,
		128536, 128538, 128540, 128541, 128557, 128560, 128563}
	for id := range emojiArr {
		emojis[uint32(id)] = true
	}
}

func replyEmoji(c *client.QQClient, event *message.GroupMessage) {
	if event.Sender.Uin != c.Uin {
		return
	}
	for _, elem := range event.GetElements() {
		if elem.Type() == message.Face {
			id := elem.(*message.FaceElement).FaceID
			if _, ok := emojis[uint32(id)]; ok {
				_ = c.GroupSetReaction(event.GroupUin, event.Id, strconv.Itoa(int(id)), true)
			}
		}
	}
}

func sgst(c *client.QQClient, event *message.GroupMessage) {
	msg := event.ToString()
	if strings.HasPrefix(msg, "sgst") {
		memberInfo := c.GetCachedMemberInfo(c.Uin, event.GroupUin)
		if memberInfo.Permission == entity.Owner {
			title := strings.TrimSpace(strings.TrimPrefix(msg, "sgst"))
			_ = c.GroupSetSpecialTitle(event.GroupUin, event.Sender.Uin, title)
		}
	}
}

func forwardMarkdown(c *client.QQClient, event *message.GroupMessage) {
	if event.Sender.Uin != c.Uin || event.ToString() != "/test" {
		return
	}
	_, err := c.SendGroupMessage(event.GroupUin,
		[]message.IMessageElement{&message.ForwardMessage{Nodes: []*message.ForwardNode{{
			GroupId:    event.GroupUin,
			SenderId:   c.Uin,
			SenderName: c.NickName(),
			Time:       0,
			Message:    []message.IMessageElement{&message.TextElement{Content: "# 标题 \\n## 简介很开心 \\n内容[🔗腾讯](https://www.qq.com)"}},
		}}}})
	if err != nil {
		return
	}
}
