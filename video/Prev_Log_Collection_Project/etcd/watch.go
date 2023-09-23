package main

import (
	"context"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err:%v", err)
		return

	}
	defer cli.Close()
	watchCh := cli.Watch(context.Background(), "s4")
	for wresp := range watchCh {
		for _, evt := range wresp.Events {
			fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}

}
