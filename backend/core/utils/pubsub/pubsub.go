package pubsub

import (
	"context"
	"sync"
)

type IPubSub interface {
	Publish(ctx context.Context, topic string, data any)
	Subscribe(ctx context.Context, topic, name string, setting SubscribeSetting) ISubscribe
	Close()
}

type ISubscribe interface {
	GetData() <-chan any
	Close()
}

type SubscribeSetting struct {
	NumBufferChan int
}

func New() IPubSub {
	return &pubSub{
		subscribers: make(map[string]map[string]*subscriber),
	}
}

type pubSub struct {
	sync.RWMutex
	subscribers map[string]map[string]*subscriber
	isClosed    bool
}

func (p *pubSub) Publish(_ context.Context, topic string, data any) {
	p.RLock()
	defer p.RUnlock()

	if p.isClosed {
		return
	}

	for _, subscribe := range p.subscribers[topic] {
		if subscribe.isClosed {
			continue
		}
		subscribe.dataChan <- data
	}
}

func (p *pubSub) Subscribe(_ context.Context, topic, name string, setting SubscribeSetting) ISubscribe {
	if p.isClosed {
		return nil
	}

	newSubscriber := &subscriber{
		dataChan: make(chan any, setting.NumBufferChan),
	}

	p.Lock()
	if p.subscribers[topic] == nil {
		p.subscribers[topic] = make(map[string]*subscriber)
	}
	p.subscribers[topic][name] = newSubscriber
	p.Unlock()

	return newSubscriber
}

func (p *pubSub) Close() {
	p.isClosed = true

	p.Lock()
	for _, subscribersPerTopic := range p.subscribers {
		for _, subscribe := range subscribersPerTopic {
			subscribe.Close()
		}
	}
	p.Unlock()
}

type subscriber struct {
	dataChan chan any
	isClosed bool
}

func (s *subscriber) GetData() <-chan any {
	return s.dataChan
}

func (s *subscriber) Close() {
	s.isClosed = true
	close(s.dataChan)
}
