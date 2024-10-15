// nolint
package main

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/LagrangeDev/LagrangeGo/personal"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var logger = personal.Logger

func main() {
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := &auth.DeviceInfo{
		Guid:          "cfcd208495d565ef66e7dff9f98764da",
		DeviceName:    "Lagrange-DCFCD07E",
		SystemKernel:  "Windows 10.0.22631",
		KernelVersion: "10.0.22631",
	}

	qqClient := client.NewClient(0, appInfo, "https://sign.lagrangecore.org/api/sign/25765")
	qqClient.SetLogger(personal.ProtocolLogger{})
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
	// 加载所有功能
	personal.Handlers.Sub(qqClient)
	personal.Handlers = nil

	err = qqClient.Login("", "qrcode.png")
	if err != nil {
		logger.Errorln("login err:", err)
		return
	}

	personal.CheckAlive(qqClient)

	defer qqClient.Release()

	defer func() {
		data, err = qqClient.Sig().Marshal()
		if err != nil {
			logger.Errorln("marshal sig.bin err:", err)
			return
		}
		err = os.WriteFile("sig.bin", data, 0644)
		if err != nil {
			logger.Errorln("write sig.bin err:", err)
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
