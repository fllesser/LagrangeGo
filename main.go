// nolint
package main

import (
	"fmt"
	"github.com/LagrangeDev/LagrangeGo/logic"
	"github.com/LagrangeDev/LagrangeGo/message"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/LagrangeDev/LagrangeGo/utils"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var (
	dumpspath = "dump"
)

func main() {
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := &auth.DeviceInfo{
		Guid:          "cfcd208495d565ef66e7dff9f98764da",
		DeviceName:    "Lagrange-DCFCD07E",
		SystemKernel:  "Windows 10.0.22631",
		KernelVersion: "10.0.22631",
	}

	qqClient := client.NewClient(0, appInfo, "https://sign.lagrangecore.org/api/sign/25765")
	qqClient.SetLogger(protocolLogger{})
	qqClient.UseDevice(deviceInfo)
	data, err := os.ReadFile("sig.bin")
	if err != nil {
		logrus.Warnln("read sig error:", err)
	} else {
		sig, err := auth.UnmarshalSigInfo(data, true)
		if err != nil {
			logrus.Warnln("load sig error:", err)
		} else {
			qqClient.UseSig(sig)
		}
	}

	qqClient.PrivateMessageEvent.Subscribe(privateLog)
	groupSub(qqClient, groupLog)
	groupSub(qqClient, logic.Sgst)

	err = qqClient.Login("", "qrcode.png")
	if err != nil {
		logrus.Errorln("login err:", err)
		return
	}

	checkAlive(qqClient)

	defer qqClient.Release()

	defer func() {
		data, err = qqClient.Sig().Marshal()
		if err != nil {
			logger.Errorln("marshal sig.bin err:", err)
			return
		}
		err = os.WriteFile("sig.bin", data, 0644)
		if err != nil {
			logrus.Errorln("write sig.bin err:", err)
			return
		}
		logger.Infoln("sig saved into sig.bin")
	}()

	// set up the main stop channel
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	for {
		switch <-mc {
		case os.Interrupt, syscall.SIGTERM:
			return
		}
	}
}

func checkAlive(c *client.QQClient) {
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
		logger.Infof("Lgr[%v] %v", c.Uin, statusContent[status])
		for range ticker.C {
			lastStatus, status = status, c.Online.Load()
			if lastStatus != status {
				logger.Infof("Lgr[%v] %v", c.Uin, statusContent[status])
			}
		}
	}()
}

func groupSub(c *client.QQClient, handler func(client *client.QQClient, event *message.GroupMessage)) {
	c.GroupMessageEvent.Subscribe(handler)
	c.SelfGroupMessageEvent.Subscribe(handler)
}

func groupLog(c *client.QQClient, event *message.GroupMessage) {
	logger.Infof(fromProtocol+"message.group[gid:%v,uid:%v,msg:%v]", event.GroupUin, event.Sender.Uin, event.ToString())
}

func privateLog(c *client.QQClient, event *message.PrivateMessage) {
	logger.Infof(fromProtocol+"message.private[uid:%v,msg:%v]", event.Sender.Uin, event.ToString())
}

// protocolLogger from https://github.com/Mrs4s/go-cqhttp/blob/a5923f179b360331786a6509eb33481e775a7bd1/cmd/gocq/main.go#L501
type protocolLogger struct{}

const fromProtocol = "Lgr -> "

func (p protocolLogger) Info(format string, arg ...any) {
	logger.Infof(fromProtocol+format, arg...)
}

func (p protocolLogger) Warning(format string, arg ...any) {
	logger.Warnf(fromProtocol+format, arg...)
}

func (p protocolLogger) Debug(format string, arg ...any) {
	logger.Debugf(fromProtocol+format, arg...)
}

func (p protocolLogger) Error(format string, arg ...any) {
	logger.Errorf(fromProtocol+format, arg...)
}

func (p protocolLogger) Dump(data []byte, format string, arg ...any) {
	msg := fmt.Sprintf(format, arg...)
	if _, err := os.Stat(dumpspath); err != nil {
		err = os.MkdirAll(dumpspath, 0o755)
		if err != nil {
			logger.Errorf("出现错误 %v. 详细信息转储失败", msg)
			return
		}
	}
	dumpFile := path.Join(dumpspath, fmt.Sprintf("%v.dump", time.Now().Unix()))
	logger.Errorf("出现错误 %v. 详细信息已转储至文件 %v 请连同日志提交给开发者处理", msg, dumpFile)
	_ = os.WriteFile(dumpFile, data, 0o644)
}

const (
	// 定义颜色代码
	colorReset  = "\x1b[0m"
	colorRed    = "\x1b[31m"
	colorYellow = "\x1b[33m"
	colorGreen  = "\x1b[32m"
	colorBlue   = "\x1b[34m"
	colorWhite  = "\x1b[37m"
)

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&ColoredFormatter{})
	logger.SetOutput(colorable.NewColorableStdout())
}

type ColoredFormatter struct{}

func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取当前时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// 根据日志级别设置相应的颜色
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = colorBlue
	case logrus.InfoLevel:
		levelColor = colorGreen
	case logrus.WarnLevel:
		levelColor = colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = colorRed
	default:
		levelColor = colorWhite
	}

	return utils.S2B(fmt.Sprintf("[%s] [%s%s%s]: %s\n",
		timestamp, levelColor, strings.ToUpper(entry.Level.String()), colorReset, entry.Message)), nil
}
