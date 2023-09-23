package main

import (
	"context"
	"fmt"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	pr, err := cli.Put(ctx, "s4", "Really Good")
	if err != nil {
		fmt.Println("put etcd failed,err:%v", err)
		return
	}
	fmt.Println("Revision:", pr.Header.Revision)
	cancel()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gr, err := cli.Get(ctx, "s4")
	if err != nil {
		fmt.Println("get etcd failed,err:%v", err)
		return
	}
	for _, ev := range gr.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	cancel()

}
