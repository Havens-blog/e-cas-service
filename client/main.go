package main

import (
	"context"
	"fmt"

	v1 "github.com/Havens-blog/e-cas-service/api/grpc/v1"
	"github.com/Havens-blog/e-cas-service/pkg/consul"
	"github.com/Havens-blog/e-cas-service/pkg/viper"
)

func main() {
	cc, err := viper.Load()
	if err != nil {
		panic("load config failed")
	}
	conn, err := consul.NewDiscovery(cc.Consul, cc.Consul.Discovery.GoGrpcLayout)
	if err != nil {
		panic(err)
	}
	client := v1.NewGrpcClient(conn)
	resp, err := client.GetUserInfo(context.Background(), &v1.UserInfoRequest{UserName: "s2ddds"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", resp)
}
