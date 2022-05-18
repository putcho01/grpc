package main

import (
	"context"
	"log"

	personpb "github.com/putcho01/grpc/proto"

	"google.golang.org/grpc"
)

// Unaryリクエスト用のメソッドを定義
func postUnaryRequest(c personpb.PersonClient) {
	log.Println("Start Unary request")

	// Helloリクエストを作成
	// はprotoファイルで定義
	req := &personpb.HelloRequest{
		Name:  "Thomas Lathan",
		Age:   30,
		Email: "Withown@example.net",
	}

	// Helloリクエストを実行
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Person request: %v", err)
	}
	log.Printf("Response from Hello Server: %v\n", res)
}

func main() {
	// オプション設定用のスライス
	var opts []grpc.DialOption
	// テストなのでfalse
	tls := false
	if tls {
		// セキュア通信処理を実装
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// サーバ側に接続
	cc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	// クライアントをインスタンス化
	c := personpb.NewPersonClient(cc)
	// リクエスト実行
	postUnaryRequest(c)
}
