# Learning Notes

## The difference between "make" and "new"

- initial memory

### new

- initial basic data types `bool`, `string`, `int`, and return pointer

### make

- initial `slice`, `map`, `channel`, and return specific type

## The difference between "Kafka" and "MQ"

### Kafka

- Kafka is a heavyweight application that combines storage and message queuing.

### MQ

- MQ is used for message queuing.

## contex

- Controlling Goroutines and Tracing Across Goroutines

Two root nodes： 
- `contex.Background()`
- `contex.TODO()`

Four functions：
- `contex.WithCancel()`
- `contex.WithDeadline()`
- `contex.WithTimeout()`
- `contex.WithValue()`：Using custom types for keys

## pprof

- The two directions of performance optimization： CPU、Memory

> go tool pprof cpu.pprof

> go tool pprof mem.pprof

You can generate a FlameGraph