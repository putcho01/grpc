package main

import (
	"context"
	"fmt"
	"log"
	"net"

	personpb "github.com/putcho01/grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	personpb.UnimplementedPersonServer
}

// Helloメソッドを実装することでprotoファイルで定義した内容と連携
func (*server) Hello(ctx context.Context, in *personpb.HelloRequest) (*personpb.HelloResponse, error) {
	// Getxxxメソッドも自動に作成されています。
	name := in.GetName()
	email := in.GetEmail()
	age := in.GetAge()
	result := fmt.Sprintf("Hello,%s (%d). email is %s", name, age, email)

	// 受け取ったメッセージを連結したレスポンスを返す
	return &personpb.HelloResponse{
		Message: result,
	}, nil
}

func main() {
	// リッスンするportをを設定
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	// サーバをインスタンス化
	s := grpc.NewServer(opts...)

	personpb.RegisterPersonServer(s, &server{})

	// 起動確認様にログ出力
	log.Println("Waiting for gRPC request ....")
	log.Println("--------------")

	// サービス起動
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
