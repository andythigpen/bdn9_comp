package main

import (
	"fmt"
	"net"

	"github.com/andythigpen/bdn9_comp/v2/cmd"
	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/andythigpen/bdn9_comp/v2/serial"
	"github.com/andythigpen/bdn9_comp/v2/service"
	"github.com/getlantern/systray"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var listening = false
var device serial.BDN9SerialDevice
var grpcServer *grpc.Server
var listener net.Listener

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
	device, err = cmd.OpenSerialDevice()
	if err != nil {
		return fmt.Errorf("Failed to open serial device: %s", err)
	}
	return nil
}

func onReady() {
	cmd.InitConfig()
	systray.SetTitle("BDN9")
	systray.SetTooltip("Keyboard daemon")

	mSerial := systray.AddMenuItem("Serial: initializing", "initializing")
	mServer := systray.AddMenuItem("Server: initializing", "initializing")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the app")

	bindAddress := viper.GetString("bind")
	if len(bindAddress) == 0 {
		bindAddress = "localhost:17432"
	}

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
