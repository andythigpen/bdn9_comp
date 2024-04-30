package service

import (
	"context"

	"github.com/andythigpen/bdn9_comp/v2/device"
	pb "github.com/andythigpen/bdn9_comp/v2/proto"
)

type bdn9Service struct {
	pb.UnimplementedBDN9ServiceServer
	device device.BDN9Device
}

func NewService(device device.BDN9Device) pb.BDN9ServiceServer {
	return &bdn9Service{
		UnimplementedBDN9ServiceServer: pb.UnimplementedBDN9ServiceServer{},
		device:                         device,
	}
}

func (s *bdn9Service) SetRGBMode(ctx context.Context, request *pb.SetRGBModeRequest) (*pb.SetRGBModeReply, error) {
	if err := s.device.SetRGBMode(pb.RGBMode(request.Mode)); err != nil {
		return nil, err
	}
	return &pb.SetRGBModeReply{}, nil
}

func (s *bdn9Service) SetMatrixHSV(ctx context.Context, request *pb.SetMatrixHSVRequest) (*pb.SetMatrixHSVReply, error) {
	if err := s.device.SetMatrixHSV(uint8(request.H), uint8(request.S), uint8(request.V)); err != nil {
		return nil, err
	}
	return &pb.SetMatrixHSVReply{}, nil
}

func (s *bdn9Service) ToggleMatrix(ctx context.Context, request *pb.ToggleMatrixRequest) (*pb.ToggleMatrixReply, error) {
	if err := s.device.ToggleMatrix(); err != nil {
		return nil, err
	}
	return &pb.ToggleMatrixReply{}, nil
}

func (s *bdn9Service) SetIndicatorRGB(ctx context.Context, request *pb.SetIndicatorRGBRequest) (*pb.SetIndicatorRGBReply, error) {
	if err := s.device.SetIndicatorRGB(pb.Layer(request.Layer), uint8(request.Key), uint8(request.R), uint8(request.G), uint8(request.B)); err != nil {
		return nil, err
	}
	return &pb.SetIndicatorRGBReply{}, nil
}

func (s *bdn9Service) ToggleIndicator(ctx context.Context, request *pb.ToggleIndicatorRequest) (*pb.ToggleIndicatorReply, error) {
	if err := s.device.ToggleIndicator(pb.Layer(request.Layer), uint8(request.Key)); err != nil {
		return nil, err
	}
	return &pb.ToggleIndicatorReply{}, nil
}

func (s *bdn9Service) EnableIndicator(ctx context.Context, request *pb.EnableIndicatorRequest) (*pb.EnableIndicatorReply, error) {
	if err := s.device.EnableIndicator(pb.Layer(request.Layer), uint8(request.Key)); err != nil {
		return nil, err
	}
	return &pb.EnableIndicatorReply{}, nil
}

func (s *bdn9Service) DisableIndicator(ctx context.Context, request *pb.DisableIndicatorRequest) (*pb.DisableIndicatorReply, error) {
	if err := s.device.DisableIndicator(pb.Layer(request.Layer), uint8(request.Key)); err != nil {
		return nil, err
	}
	return &pb.DisableIndicatorReply{}, nil
}

func (s *bdn9Service) ActivateLayer(ctx context.Context, request *pb.ActivateLayerRequest) (*pb.ActivateLayerReply, error) {
	if err := s.device.ActivateLayer(pb.Layer(request.Layer)); err != nil {
		return nil, err
	}
	return &pb.ActivateLayerReply{}, nil
}

func (s *bdn9Service) SetSpeed(ctx context.Context, request *pb.SetSpeedRequest) (*pb.SetSpeedReply, error) {
	if err := s.device.SetSpeed(uint8(request.Speed)); err != nil {
		return nil, err
	}
	return &pb.SetSpeedReply{}, nil
}

func (s *bdn9Service) SetMuteStatus(ctx context.Context, request *pb.SetMuteStatusRequest) (*pb.SetMuteStatusReply, error) {
	if err := s.device.SetMuteStatus(request.Muted); err != nil {
		return nil, err
	}
	return &pb.SetMuteStatusReply{}, nil
}

func (s *bdn9Service) EndCall(ctx context.Context, request *pb.EndCallRequest) (*pb.EndCallReply, error) {
	if err := s.device.EndCall(); err != nil {
		return nil, err
	}
	return &pb.EndCallReply{}, nil
}

func (s *bdn9Service) Reset(ctx context.Context, request *pb.ResetRequest) (*pb.ResetReply, error) {
	if err := s.device.Reset(); err != nil {
		return nil, err
	}
	return &pb.ResetReply{}, nil
}

func (s *bdn9Service) Echo(ctx context.Context, request *pb.EchoRequest) (*pb.EchoReply, error) {
	if err := s.device.Echo([]byte{0x1, 0x2, 0x3}); err != nil {
		return nil, err
	}
	return &pb.EchoReply{}, nil
}
