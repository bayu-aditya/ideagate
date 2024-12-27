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
	return &pubSub{}
}

type pubSub struct {
	subscribersV2 sync.Map // map[topic]map[subscriberName]*subscriber
	isClosed      bool
}

func (p *pubSub) Publish(_ context.Context, topic string, data any) {
	if p.isClosed {
		return
	}

	subscribersPerTopic, ok := p.subscribersV2.Load(topic)
	if !ok {
		// if not found, return
		return
	}

	subscribersPerTopicMap := subscribersPerTopic.(*sync.Map)
	subscribersPerTopicMap.Range(func(subscribeName, subscribe interface{}) bool {
		subs := subscribe.(*subscriber)
		if subs.isClosed {
			return true
		}
		subs.dataChan <- data
		return true
	})
}

func (p *pubSub) Subscribe(_ context.Context, topic, name string, setting SubscribeSetting) ISubscribe {
	if p.isClosed {
		return nil
	}

	newSubscriber := &subscriber{
		dataChan: make(chan any, setting.NumBufferChan),
	}

	subscribersPerTopic, _ := p.subscribersV2.LoadOrStore(topic, &sync.Map{})
	subscribersPerTopicMap := subscribersPerTopic.(*sync.Map)
	subscribersPerTopicMap.Store(name, newSubscriber)

	return newSubscriber
}

func (p *pubSub) Close() {
	p.isClosed = true

	p.subscribersV2.Range(func(topic, subscribers interface{}) bool {
		subscribersPerTopicMap := subscribers.(*sync.Map)
		subscribersPerTopicMap.Range(func(_, subscribe interface{}) bool {
			subscribe.(*subscriber).Close()
			return true
		})
		return true
	})
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
