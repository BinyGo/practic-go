package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"

	pb "github.com/practic-go/net/grpc/tag-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/practic-go/net/grpc/tag-service/pkg/bapi"
)

var port string
var grpcPort string
var httpPort string

func init() {
	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.StringVar(&grpcPort, "grpcPort", "8000", "gRPC 启动端口号")
	flag.StringVar(&httpPort, "http_port", "9001", "HTTP 启动端口号")
	flag.Parse()
}

type TagServer struct {
	pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}
func main() {
	errs := make(chan error)
	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		log.Fatalf("Run Server err: %v", err)
	}

	/* 	lis, err := net.Listen("tcp", ":"+port)
	   	if err != nil {
	   		log.Fatalf("net.Listen err: %v", err)
	   	}

	   	err = s.Serve(lis)
	   	if err != nil {
	   		log.Fatalf("server.Serve err: %v", err)
	   	} */
}

//grpcurl -plaintext localhost:8000 tag.TagService.GetTagList
func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, NewTagServer())
	reflection.Register(s) //将注册的服务反射出来
	/*
		$go get github.com/fullstorydev/grpcurl/cmd/grpcurl
		$ grpcurl -plaintext localhost:8000 list //查看服务
			grpc.reflection.v1alpha.ServerReflection
			tag.TagService
		$ grpcurl -plaintext localhost:8000 list tag.TagService //查看服务方法
			tag.TagService.GetTagList
		$ grpcurl -plaintext -d '{"name":"Go"}' localhost:8000 tag.TagService.GetTagList 执行rpc方法
	*/
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}

//curl http://127.0.0.1:9001/ping
func RunHttpServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return http.ListenAndServe(":"+port, serveMux)
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8999")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, err
		//return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}
