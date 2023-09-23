
# Issue

> `go run main.go`
```
send message failed, kafka server: In the middle of a leadership election, there is currently no leader for this partition and hence it is unavailable for writes.
pid:-1 offset:-1
```

- Running Kafka and ZooKeeper in Docker

- Kafka Log
```
[2023-09-11 16:21:15,824] INFO Topic creation {"version":1,"partitions":{"0":[0]}} (kafka.admin.AdminUtils$)
[2023-09-11 16:21:15,829] INFO [KafkaApi-0] Auto creation of topic shoppings with 1 partitions and replication factor 1 is successful (kafka.server.KafkaApis)
[2023-09-11 16:27:27,118] INFO [Group Metadata Manager on Broker 0]: Removed 0 expired offsets in 0 milliseconds. (kafka.coordinator.GroupMetadataManager)
[2023-09-11 16:37:27,123] INFO [Group Metadata Manager on Broker 0]: Removed 0 expired offsets in 3 milliseconds. (kafka.coordinator.GroupMetadataManager)
```


