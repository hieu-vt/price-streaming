package pubsub

import (
	"context"
	"flag"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/nats-io/nats.go"
	"time"
)

type natsPubSub struct {
	name       string
	logger     logger.Logger
	connection *nats.Conn
	url        string
}

func NewNatsPubSub(prefix string) *natsPubSub {
	return &natsPubSub{
		name: prefix,
	}
}

func (pb *natsPubSub) GetPrefix() string {
	return pb.name
}

func (pb *natsPubSub) Get() interface{} {
	return pb
}

func (pb *natsPubSub) Name() string {
	return pb.name
}

func (pb *natsPubSub) InitFlags() {
	flag.StringVar(&pb.url, pb.name+"-url", nats.DefaultURL, "Url connect nats pubsub")
}

func (pb *natsPubSub) Configure() error {
	pb.logger = logger.GetCurrent().GetLogger(pb.name)
	nc, _ := nats.Connect(pb.url, pb.setupConnOptions([]nats.Option{})...)

	pb.connection = nc
	pb.logger.Infoln("Connected to NATS service.")

	return nil
}

func (pb *natsPubSub) Run() error {
	return pb.Configure()
}

func (pb *natsPubSub) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (pb *natsPubSub) setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		pb.logger.Infof("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		pb.logger.Infof("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		pb.logger.Infof("Exiting: %v", nc.LastError())
	}))
	return opts
}

func (pb *natsPubSub) Publish(ctx context.Context, channel string, data string) error {
	err := pb.connection.Publish(channel, []byte(data))

	if err != nil {
		pb.logger.Infoln(err)
	}

	return nil
}
