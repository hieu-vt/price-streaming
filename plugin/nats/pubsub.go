package nats

import "context"

type NatsPubsub interface {
	Publish(ctx context.Context, channel string, data string) error
}
