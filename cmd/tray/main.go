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

func main() {
	onExit := func() {
		if device != nil && device.IsOpen() {
			device.Close()
		}
	}

	systray.Run(onReady, onExit)
}

func startServer() error {
	bindAddress := viper.GetString("bind")
	if len(bindAddress) == 0 {
		bindAddress = "localhost:17432"
	}
	lis, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	device, err = cmd.OpenSerialDevice()
	if err != nil {
		return err
	}
	server := service.NewService(device)
	pb.RegisterBDN9ServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	listening = true
	return nil
}

func onReady() {
	cmd.InitConfig()
	systray.SetTitle("BDN9")
	systray.SetTooltip("Keyboard daemon")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the app")

	mStatus := systray.AddMenuItem("Status: initializing", "initializing")

	listen := func() {
		if err := startServer(); err != nil {
			mStatus.SetTitle(fmt.Sprintf("Error: %s", err))
		} else {
			mStatus.SetTitle("Status: running")
		}
	}

	go func() {
		for {
			select {
			case <-mQuitOrig.ClickedCh:
				systray.Quit()
				return
			case <-mStatus.ClickedCh:
				if !listening {
					listen()
				}
			}
		}
	}()

	go listen()
}
