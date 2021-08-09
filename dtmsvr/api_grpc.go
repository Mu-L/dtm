package dtmsvr

import (
	"context"
	"log"
	"strings"

	"github.com/yedf/dtm/dtmcli"
	pb "github.com/yedf/dtm/dtmpb"
	"google.golang.org/grpc"
)

// dtmServer is used to implement helloworld.GreeterServer.
type dtmServer struct {
	pb.UnimplementedDtmServer
}

func (s *dtmServer) Call(ctx context.Context, in *pb.DtmRequest) (*pb.DtmReply, error) {
	log.Printf("dtmServer Received: %v", in)
	dynamicCallPb(ctx, in, in.Extra["BusiFunc"], []byte(in.Data))
	return &pb.DtmReply{DtmResult: "SUCCESS", DtmMessage: "ok"}, nil
}

func (s *dtmServer) Submit(ctx context.Context, in *pb.DtmRequest) (*pb.DtmReply, error) {
	svcSubmit(TransFromDtmRequest(in), in.WaitResult)
	return &pb.DtmReply{DtmResult: "SUCCESS", DtmMessage: "ok"}, nil
}

func dynamicCallPb(ctx context.Context, in *pb.DtmRequest, pbAddr string, data []byte) error {
	fs := strings.Split(pbAddr, "/")
	grpcAddr := fs[0]
	method := "/" + strings.Join(fs[1:], "/")
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(dtmcli.GrpcClientLog))
	dtmcli.FatalIfError(err)
	reply := &pb.BusiReply{}
	err = conn.Invoke(ctx, method, &pb.BusiRequest{Info: &pb.DtmTransInfo{Gid: in.Gid}}, reply)
	dtmcli.FatalIfError(err)
	return err
}
