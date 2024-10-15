package personal

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
)

var (
	Handlers *CustomHandlers = NewCustomHandlers()
)

func init() {
	Handlers.AddPrivate(privateLog)
	Handlers.AddGroup(groupLog, true)
	Handlers.AddGroup(sgst, true)
	Handlers.AddGroup(replyEmoji, true)
	//Handlers.AddGroup(forwardMarkdown, false)
}

type (
	CustomHandlers struct {
		groupMessageHandlers       []Handler[*message.GroupMessage]
		selfGroupMessageHandlers   []Handler[*message.GroupMessage]
		privateMessageHandlers     []Handler[*message.PrivateMessage]
		selfPrivateMessageHandlers []Handler[*message.PrivateMessage]
	}

	Handler[T any] func(c *client.QQClient, event T)
)

func (h *CustomHandlers) Sub(c *client.QQClient) {
	for _, handler := range h.groupMessageHandlers {
		c.GroupMessageEvent.Subscribe(handler)
		c.SelfGroupMessageEvent.Subscribe(handler)
	}
	for _, handler := range h.privateMessageHandlers {
		c.PrivateMessageEvent.Subscribe(handler)
	}
}

func (h *CustomHandlers) AddGroup(handler Handler[*message.GroupMessage], isSelf bool) {
	if isSelf {
		h.selfGroupMessageHandlers = append(h.selfGroupMessageHandlers, handler)
	}

	h.groupMessageHandlers = append(h.groupMessageHandlers, handler)
}

func (h *CustomHandlers) AddPrivate(handler Handler[*message.PrivateMessage]) {
	h.privateMessageHandlers = append(h.privateMessageHandlers, handler)
}

func NewCustomHandlers() *CustomHandlers {
	return &CustomHandlers{
		groupMessageHandlers:   make([]Handler[*message.GroupMessage], 0),
		privateMessageHandlers: make([]Handler[*message.PrivateMessage], 0),
	}
}
