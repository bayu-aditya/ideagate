package pubsub

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func Test_pubSub_Publish(t *testing.T) {
	type args struct {
		ctx   context.Context
		topic string
		data  any
	}
	tests := []struct {
		name     string
		args     args
		funcTest func(*testing.T, IPubSub, args)
	}{
		{
			name: "without subscriber",
			args: args{
				ctx:   context.TODO(),
				topic: "topic_1",
				data:  "mock_data",
			},
			funcTest: func(t *testing.T, pubSub IPubSub, args args) {
				pubSub.Publish(args.ctx, args.topic, args.data)
			},
		},
		{
			name: "with 1 subscriber",
			args: args{
				ctx:   context.TODO(),
				topic: "topic_1",
				data:  "mock_data",
			},
			funcTest: func(t *testing.T, pubSub IPubSub, args args) {
				wg := sync.WaitGroup{}

				var gotData any
				go func() {
					defer wg.Done()
					wg.Add(1)
					gotData = <-pubSub.Subscribe(args.ctx, args.topic, SubscribeSetting{
						NumBufferChan: 1,
					})
				}()

				time.Sleep(1 * time.Second)
				pubSub.Publish(args.ctx, args.topic, args.data)

				wg.Wait()
				assert.Equal(t, "mock_data", gotData)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New()

			tt.funcTest(t, p, tt.args)
		})
	}
}
