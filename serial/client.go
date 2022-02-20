package serial

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/andythigpen/bdn9_comp/v2/proto"
)

type bdn9SerialClient struct {
	device BDN9SerialDevice
}

func NewSerialClient(device BDN9SerialDevice) pb.BDN9ServiceClient {
	return &bdn9SerialClient{device}
}

func (c *bdn9SerialClient) SetRGBMode(ctx context.Context, in *pb.SetRGBModeRequest, opts ...grpc.CallOption) (*pb.SetRGBModeReply, error) {
	return nil, c.device.SetRGBMode(RGBMode(in.Mode))
}

func (c *bdn9SerialClient) SetMatrixHSV(ctx context.Context, in *pb.SetMatrixHSVRequest, opts ...grpc.CallOption) (*pb.SetMatrixHSVReply, error) {
	return nil, c.device.SetMatrixHSV(uint8(in.H), uint8(in.S), uint8(in.V))
}

func (c *bdn9SerialClient) ToggleMatrix(ctx context.Context, in *pb.ToggleMatrixRequest, opts ...grpc.CallOption) (*pb.ToggleMatrixReply, error) {
	return nil, c.device.ToggleMatrix()
}

func (c *bdn9SerialClient) SetIndicatorHSV(ctx context.Context, in *pb.SetIndicatorHSVRequest, opts ...grpc.CallOption) (*pb.SetIndicatorHSVReply, error) {
	return nil, c.device.SetIndicatorHSV(Layer(in.Layer), uint8(in.Key), uint8(in.H), uint8(in.S), uint8(in.V))
}

func (c *bdn9SerialClient) ToggleIndicator(ctx context.Context, in *pb.ToggleIndicatorRequest, opts ...grpc.CallOption) (*pb.ToggleIndicatorReply, error) {
	return nil, c.device.ToggleIndicator(Layer(in.Layer), uint8(in.Key))
}

func (c *bdn9SerialClient) EnableIndicator(ctx context.Context, in *pb.EnableIndicatorRequest, opts ...grpc.CallOption) (*pb.EnableIndicatorReply, error) {
	return nil, c.device.EnableIndicator(Layer(in.Layer), uint8(in.Key))
}

func (c *bdn9SerialClient) DisableIndicator(ctx context.Context, in *pb.DisableIndicatorRequest, opts ...grpc.CallOption) (*pb.DisableIndicatorReply, error) {
	return nil, c.device.DisableIndicator(Layer(in.Layer), uint8(in.Key))
}

func (c *bdn9SerialClient) ActivateLayer(ctx context.Context, in *pb.ActivateLayerRequest, opts ...grpc.CallOption) (*pb.ActivateLayerReply, error) {
	return nil, c.device.ActivateLayer(Layer(in.Layer))
}

func (c *bdn9SerialClient) SetSpeed(ctx context.Context, in *pb.SetSpeedRequest, opts ...grpc.CallOption) (*pb.SetSpeedReply, error) {
	return nil, c.device.SetSpeed(uint8(in.Speed))
}

func (c *bdn9SerialClient) SetMuteStatus(ctx context.Context, in *pb.SetMuteStatusRequest, opts ...grpc.CallOption) (*pb.SetMuteStatusReply, error) {
	return nil, c.device.SetMuteStatus(in.Muted)
}

func (c *bdn9SerialClient) EndCall(ctx context.Context, in *pb.EndCallRequest, opts ...grpc.CallOption) (*pb.EndCallReply, error) {
	return nil, c.device.EndCall()
}

func (c *bdn9SerialClient) Reset(ctx context.Context, in *pb.ResetRequest, opts ...grpc.CallOption) (*pb.ResetReply, error) {
	return nil, c.device.Reset()
}

func (c *bdn9SerialClient) Echo(ctx context.Context, in *pb.EchoRequest, opts ...grpc.CallOption) (*pb.EchoReply, error) {
	return nil, c.device.Echo([]byte{0x1, 0x2, 0x3})
}
