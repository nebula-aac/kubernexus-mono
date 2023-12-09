package domain

import (
	"sync"

	"github.com/nebula-aac/kubernexus-mono/pkg/utils"
)

type Channel interface {
	Subscribe(ch chan struct{})
	Publish()
}

type ConfigurationChannel struct {
	ApplicationsChannel []Channel
	PatternsChannel     []Channel
	FiltersChannel      []Channel
	mx                  sync.Mutex
}

func NewConfigurationChannel() *ConfigurationChannel {
	return &ConfigurationChannel{
		ApplicationsChannel: make([]Channel, 10),
		PatternsChannel:     make([]Channel, 10),
		FiltersChannel:      make([]Channel, 10),
	}
}

func (c *ConfigurationChannel) Subscribe(channel Channel, ch chan struct{}) {
	c.mx.Lock()
	defer c.mx.Unlock()
	channel.Subscribe(ch)
}

func (c *ConfigurationChannel) Publish(channels []Channel) {
	for _, channel := range channels {
		channel.Publish()
	}
}

type ExtendedChannel struct {
	ch []chan struct{}
}

func NewExtendedChannel() *ExtendedChannel {
	return &ExtendedChannel{
		ch: make([]chan struct{}, 10),
	}
}

func (cc *ExtendedChannel) Subscribe(ch chan struct{}) {
	cc.ch = append(cc.ch, ch)
}

func (cc *ExtendedChannel) Publish() {
	for _, ch := range cc.ch {
		if !utils.IsClosed(ch) {
			ch <- struct{}{}
		}
	}
}
