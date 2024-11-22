package pubsub

import (
	"context"
	"sync"
)

type IPubSub interface {
	Publish(ctx context.Context, topic string, data any)
	Subscribe(ctx context.Context, topic string, setting SubscribeSetting) <-chan any
	SubscribeV2(ctx context.Context, topic, name string, setting SubscribeSetting) ISubscribe
	Close()
}

type ISubscribe interface {
	Close()
}

type SubscribeSetting struct {
	NumBufferChan int
}

func New() IPubSub {
	return &pubSub{
		subscriber: make(map[string][]chan any),
	}
}

type pubSub struct {
	sync.RWMutex
	subscriber   map[string][]chan any // deprecated
	subscriberV2 map[string]map[string]ISubscribe
	isClosed     bool
}

func (p *pubSub) Publish(_ context.Context, topic string, data any) {
	p.RLock()
	defer p.RUnlock()

	if p.isClosed {
		return
	}

	for _, subsChan := range p.subscriber[topic] {
		subsChan <- data
	}
}

func (p *pubSub) Subscribe(_ context.Context, topic string, setting SubscribeSetting) <-chan any {
	if p.isClosed {
		return nil
	}

	dataChan := make(chan any, setting.NumBufferChan)

	p.Lock()
	p.subscriber[topic] = append(p.subscriber[topic], dataChan)
	p.Unlock()

	return dataChan
}

func (p *pubSub) SubscribeV2(_ context.Context, topic, name string, setting SubscribeSetting) ISubscribe {
	//if p.isClosed {
	//	return nil
	//}
	//
	//dataChan := make(chan any, setting.NumBufferChan)
	//
	//p.Lock()
	//p.subscriber[topic] = append(p.subscriber[topic], dataChan)
	//p.Unlock()

	return nil // TODO implement this
}

func (p *pubSub) Close() {
	p.isClosed = true

	p.Lock()
	for _, channels := range p.subscriber {
		for _, channel := range channels {
			close(channel)
		}
	}
	p.Unlock()
}

type subscriber struct {
	dataChan chan any
}

func (s *subscriber) Close() {
	close(s.dataChan)
}
