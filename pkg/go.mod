module github.com/andrew-pavlov-ua/pkg

go 1.25.7

require (
	github.com/andrew-pavlov-ua/proto v0.0.0
	github.com/rabbitmq/amqp091-go v1.10.0
	go.uber.org/zap v1.27.1
	google.golang.org/protobuf v1.36.11
)

replace github.com/andrew-pavlov-ua/proto => ../proto

require (
	github.com/stretchr/testify v1.11.1 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/grpc v1.79.1 // indirect
)
