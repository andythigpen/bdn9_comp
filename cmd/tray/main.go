package main

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/andythigpen/bdn9_comp/v2/cmd"
	"github.com/andythigpen/bdn9_comp/v2/cmd/tray/icon"
	device1 "github.com/andythigpen/bdn9_comp/v2/device"
	pb "github.com/andythigpen/bdn9_comp/v2/proto"
	"github.com/andythigpen/bdn9_comp/v2/service"
	"github.com/getlantern/systray"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var listening = false
var device device1.BDN9Device
var grpcServer *grpc.Server
var server pb.BDN9ServiceServer
var listener net.Listener
var handler eventHandler
var muted = false
var slackPid int32 = 0
var teamsPid int32 = 0

type eventHandler struct{}

func findPid(windowName string, windowTitle string) (int32, error) {
	pids, err := robotgo.FindIds(windowName)
	if err != nil {
		return 0, err
	}
	if len(pids) == 0 {
		return 0, fmt.Errorf("No window with name %s found", windowName)
	}
	if len(pids) > 1 {
		foundPids := make([]int32, 0)
		for _, pid := range pids {
			title := robotgo.GetTitle(pid)
			if strings.Compare(strings.ToLower(windowTitle), strings.ToLower(title)) == 0 {
				foundPids = append(foundPids, pid)
			}
		}
		if len(foundPids) == 0 {
			return 0, fmt.Errorf("No windows with title %s found", windowTitle)
		}
		if len(foundPids) > 1 {
			return 0, fmt.Errorf("Multiple windows with title %s found", windowTitle)
		}
		return foundPids[0], nil
	}
	return pids[0], nil
}

func focusWindow(windowName string, windowTitle string) error {
	pid, err := findPid(windowName, windowTitle)
	if err != nil {
		return err
	}
	return robotgo.ActivePID(pid)
}

func sendMuteKeys(pid int32, windowName string, windowTitle string, muteKeys []string) error {
	if pid == 0 {
		if err := focusWindow(windowName, windowTitle); err != nil {
			return err
		}
	} else if err := robotgo.ActivePID(pid); err != nil {
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

func (h eventHandler) HandleEvent(d device1.BDN9Device, ev device1.Event) {
	var err error
	slackWindowName := viper.GetString("slackWindowName")
	slackWindowTitle := viper.GetString("slackWindowTitle")
	teamsWindowName := viper.GetString("teamsWindowName")
	teamsWindowTitle := viper.GetString("teamsWindowTitle")

	fmt.Printf("ev: %v\n", ev)
	switch ev {
	case device1.EVENT_FOCUS_SLACK:
		err = focusWindow(slackWindowName, slackWindowTitle)
		if err != nil {
			fmt.Printf("No program found: %s", err)
			return
		}
	case device1.EVENT_FOCUS_TEAMS:
		err = focusWindow(teamsWindowName, teamsWindowTitle)
		if err != nil {
			fmt.Printf("No program found: %s", err)
			return
		}
	case device1.EVENT_MUTE_SLACK:
		slackMuteKeys := viper.GetStringSlice("slackMuteKeys")
		err = sendMuteKeys(slackPid, slackWindowName, slackWindowTitle, slackMuteKeys)
		if err != nil {
			fmt.Printf("No program found: %s", err)
			return
		}
	case device1.EVENT_MUTE_TEAMS:
		teamsWindowName := viper.GetString("teamsWindowName")
		teamsMuteKeys := viper.GetStringSlice("teamsMuteKeys")
		err := sendMuteKeys(teamsPid, teamsWindowName, teamsWindowTitle, teamsMuteKeys)
		if err != nil {
			fmt.Printf("No program found: %s", err)
			return
		}
	case device1.EVENT_START_SLACK:
		slackPid, err = findPid(slackWindowName, slackWindowTitle)
		if err != nil {
			fmt.Printf("%s\n", err)
			// try again after waiting a second
			go func() {
				time.Sleep(1 * time.Second)
				slackPid, err = findPid(slackWindowName, slackWindowTitle)
				if err != nil {
					fmt.Printf("%s\n", err)
				}
			}()
		}
	case device1.EVENT_START_TEAMS:
		teamsPid, err = findPid(teamsWindowName, teamsWindowTitle)
		if err != nil {
			fmt.Printf("%s\n", err)
			// try again after waiting a second
			go func() {
				time.Sleep(1 * time.Second)
				teamsPid, err = findPid(teamsWindowName, teamsWindowTitle)
				if err != nil {
					fmt.Printf("%s\n", err)
				}
			}()
		}
	case device1.EVENT_END_CALL:
		slackPid = 0
		teamsPid = 0
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

func connectUSBDevice() error {
	var err error
	handler = eventHandler{}
	device, err = cmd.OpenUSBDevice(handler)
	if err != nil {
		return fmt.Errorf("Failed to open device: %s", err)
	}
	return nil
}

func onReady() {
	cmd.InitConfig()
	systray.SetTooltip("Keyboard daemon")
	systray.SetTemplateIcon(icon.Data, icon.Data)

	mClearIndicators := systray.AddMenuItem("Clear indicators", "Clear indicators")
	systray.AddSeparator()
	mFocusSlack := systray.AddMenuItem("Focus Slack", "Focus Slack")
	mFocusTeams := systray.AddMenuItem("Focus Teams", "Focus Teams")
	mLayers := systray.AddMenuItem("Layers", "Layers")
	mDefaultLayer := mLayers.AddSubMenuItem("Default", "Default")
	mCalls := systray.AddMenuItem("Calls", "Calls")
	mStartSlack := mCalls.AddSubMenuItem("Start Slack Call", "Start Slack Call")
	mStartTeams := mCalls.AddSubMenuItem("Start Teams Call", "Start Teams Call")
	mMuteToggle := mCalls.AddSubMenuItem("Toggle Mute", "Toggle Mute")
	mEndCall := mCalls.AddSubMenuItem("End Call", "End a call in progress")
	systray.AddSeparator()
	mUSB := systray.AddMenuItem("USB: initializing", "initializing")
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
		server = service.NewService(device)
		pb.RegisterBDN9ServiceServer(grpcServer, server)
		mServer.SetTitle(fmt.Sprintf("Server: listening on %s", bindAddress))
		mServer.SetTooltip("listening")
		mClearIndicators.Enable()
		mStartSlack.Enable()
		mStartTeams.Enable()
		mMuteToggle.Enable()
		mEndCall.Enable()
		// this is a blocking call
		if err := grpcServer.Serve(listener); err != nil {
			mServer.SetTitle(fmt.Sprintf("Failed to serve: %s", err))
			return
		}
	}

	initUSB := func() {
		mClearIndicators.Disable()
		mStartSlack.Disable()
		mStartSlack.Disable()
		mMuteToggle.Disable()
		mEndCall.Disable()
		if device != nil && device.IsOpen() {
			device.Close()
		}
		if err := connectUSBDevice(); err != nil {
			mUSB.SetTitle(fmt.Sprintf("Error: %s", err))
		} else {
			mUSB.SetTitle("USB: connected")
			mUSB.SetTooltip("connected")
			initServer()
		}
	}

	go func() {
		for {
			select {
			case <-mQuitOrig.ClickedCh:
				systray.Quit()
				return
			case <-mClearIndicators.ClickedCh:
				if server == nil {
					continue
				}
				ctx := context.Background()
				for l := pb.Layer_LAYER_DEFAULT; l < pb.Layer_LAYER_MAX; l++ {
					for k := 0; k < pb.KEY_UPPER_MAX; k++ {
						server.DisableIndicator(ctx, &pb.DisableIndicatorRequest{
							Layer: pb.Layer(l),
							Key:   uint32(k),
						})
					}
				}
			case <-mFocusSlack.ClickedCh:
				slackWindowName := viper.GetString("slackWindowName")
				slackWindowTitle := viper.GetString("slackWindowTitle")
				if err := focusWindow(slackWindowName, slackWindowTitle); err != nil {
					fmt.Printf("%s\n", err)
				}
			case <-mFocusTeams.ClickedCh:
				teamsWindowName := viper.GetString("teamsWindowName")
				teamsWindowTitle := viper.GetString("teamsWindowTitle")
				if err := focusWindow(teamsWindowName, teamsWindowTitle); err != nil {
					fmt.Printf("%s\n", err)
				}
			case <-mDefaultLayer.ClickedCh:
				if server == nil {
					continue
				}
				ctx := context.Background()
				server.ActivateLayer(ctx, &pb.ActivateLayerRequest{
					Layer: pb.Layer_LAYER_DEFAULT,
				})
			case <-mStartSlack.ClickedCh:
				if server == nil {
					continue
				}
				ctx := context.Background()
				server.ActivateLayer(ctx, &pb.ActivateLayerRequest{
					Layer: pb.Layer_LAYER_SLACK,
				})
				muted = false
				server.SetMuteStatus(ctx, &pb.SetMuteStatusRequest{Muted: muted})
			case <-mStartTeams.ClickedCh:
				if server == nil {
					continue
				}
				ctx := context.Background()
				server.ActivateLayer(ctx, &pb.ActivateLayerRequest{
					Layer: pb.Layer_LAYER_TEAMS,
				})
				muted = true
				server.SetMuteStatus(ctx, &pb.SetMuteStatusRequest{Muted: muted})
			case <-mMuteToggle.ClickedCh:
				if server == nil {
					continue
				}
				ctx := context.Background()
				muted = !muted
				server.SetMuteStatus(ctx, &pb.SetMuteStatusRequest{Muted: muted})
			case <-mEndCall.ClickedCh:
				if server == nil {
					continue
				}
				ctx := context.Background()
				server.EndCall(ctx, &pb.EndCallRequest{})
			case <-mUSB.ClickedCh:
				initUSB()
			case <-mServer.ClickedCh:
				initServer()
			}
		}
	}()

	go initUSB()
}
