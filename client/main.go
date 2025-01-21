package main

import (
	"context"
	"fmt"
	v1 "github.com/Havens-blog/e-cas-service/api/grpc/v1"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	/*cc, err := viper.Load()
	if err != nil {
		panic("load config failed")
	}
	conn, err := consul.NewDiscovery(cc.Consul, cc.Consul.Discovery.GoGrpcLayout)
	if err != nil {
		panic(err)
	}*/

	conn, err := grpc.Dial("127.0.0.1:9100", grpc.WithInsecure()) // server_address 替换为实际服务器地址
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := v1.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	resp, err := client.GetUserInfo(ctx, &v1.UserInfoRequest{UserName: "admin"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", resp)
}
