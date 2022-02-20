package main

import (
	"fmt"
	"net"

	"github.com/andythigpen/bdn9_comp/v2/cmd"
	"github.com/andythigpen/bdn9_comp/v2/cmd/tray/icon"
	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/andythigpen/bdn9_comp/v2/service"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var listening = false
var device serial.BDN9SerialDevice
var grpcServer *grpc.Server
var listener net.Listener
var handler serialHandler

type serialHandler struct{}

func sendMuteKeys(windowName string, muteKeys []string) error {
	pids, err := robotgo.FindIds(windowName)
	if err != nil {
		return err
	}
	if len(pids) == 0 {
		return fmt.Errorf("No window with name %s found", windowName)
	}
	if len(pids) > 1 {
		return fmt.Errorf("Multiple windows with name %s found", windowName)
	}
	err = robotgo.ActiveName(windowName)
	if err != nil {
		return err
	}
	key := muteKeys[0]
	muteKeys = muteKeys[1:]
	args := make([]interface{}, len(muteKeys))
	for i := range muteKeys {
		args[i] = muteKeys[i]
	}
	robotgo.KeyTap(key, args...)
	return nil
}

func (h serialHandler) HandleEvent(d serial.BDN9SerialDevice, ev serial.Event) {
	fmt.Printf("ev: %v\n", ev)
	switch ev {
	case serial.EVENT_MUTE_SLACK:
		slackWindowName := viper.GetString("slackWindowName")
		slackMuteKeys := viper.GetStringSlice("slackMuteKeys")
		err := sendMuteKeys(slackWindowName, slackMuteKeys)
		if err != nil {
			fmt.Printf("No program found: %s", err)
			return
		}
	case serial.EVENT_MUTE_TEAMS:
		teamsWindowName := viper.GetString("teamsWindowName")
		teamsMuteKeys := viper.GetStringSlice("teamsMuteKeys")
		err := sendMuteKeys(teamsWindowName, teamsMuteKeys)
		if err != nil {
			fmt.Printf("No program found: %s", err)
			return
		}
	}
}

func main() {
	onExit := func() {
		if device != nil && device.IsOpen() {
			device.Close()
		}
	}

	systray.Run(onReady, onExit)
}

func connectSerialDevice() error {
	var err error
	handler = serialHandler{}
	device, err = cmd.OpenSerialDevice(handler)
	if err != nil {
		return fmt.Errorf("Failed to open serial device: %s", err)
	}
	return nil
}

func onReady() {
	cmd.InitConfig()
	systray.SetTooltip("Keyboard daemon")
	systray.SetTemplateIcon(icon.Data, icon.Data)

	mSerial := systray.AddMenuItem("Serial: initializing", "initializing")
	mServer := systray.AddMenuItem("Server: initializing", "initializing")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the app")

	bindAddress := viper.GetString("bind")

	initServer := func() {
		var err error
		if grpcServer != nil {
			grpcServer.Stop()
		}
		listener, err = net.Listen("tcp", bindAddress)
		if err != nil {
			mServer.SetTitle(fmt.Sprintf("Failed to bind: %s", err))
			return
		}
		var opts []grpc.ServerOption
		grpcServer = grpc.NewServer(opts...)
		server := service.NewService(device)
		pb.RegisterBDN9ServiceServer(grpcServer, server)
		mServer.SetTitle(fmt.Sprintf("Server: listening on %s", bindAddress))
		mServer.SetTooltip("listening")
		// this is a blocking call
		if err := grpcServer.Serve(listener); err != nil {
			mServer.SetTitle(fmt.Sprintf("Failed to serve: %s", err))
			return
		}
	}

	initSerial := func() {
		if device != nil && device.IsOpen() {
			device.Close()
		}
		if err := connectSerialDevice(); err != nil {
			mSerial.SetTitle(fmt.Sprintf("Error: %s", err))
		} else {
			mSerial.SetTitle(fmt.Sprintf("Serial: connected to %s", device.Name()))
			mSerial.SetTooltip("connected")
			initServer()
		}
	}

	go func() {
		for {
			select {
			case <-mQuitOrig.ClickedCh:
				systray.Quit()
				return
			case <-mSerial.ClickedCh:
				initSerial()
			case <-mServer.ClickedCh:
				initServer()
			}
		}
	}()

	go initSerial()
}
