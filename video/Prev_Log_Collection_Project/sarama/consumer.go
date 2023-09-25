package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log")
	fmt.Println(partitionList)
	var wg sync.WaitGroup
	if err != nil {
		fmt.Printf("fail to get list of partition, err:%v", err)
		return
	}
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d,err:%v", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
	wg.Wait()

}
