package manager 

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Dispatcher struct {
	subscribers    map[int]dispatchSubscriber
	subscribersMux *sync.Mutex
}

type dispatchSubscriber struct {
	notifyCancelOrCloseChan chan error
	closeCh                 <-chan struct{}
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		subscribers:    make(map[int]dispatchSubscriber),
		subscribersMux: &sync.Mutex{},
	}
}

func (d *Dispatcher) Dispatch(err error) error {
	d.subscribersMux.Lock()
	defer d.subscribersMux.Unlock()
	for _, subscriber := range d.subscribers {
		select {
		case <-time.After(time.Second * 5):
			log.Println("Unexpected rabbitmq error: timeout in dispatch")
		case subscriber.notifyCancelOrCloseChan <- err:
		}
	}
	return nil
}

func (d *Dispatcher) AddSubscriber() (<-chan error, chan<- struct{}) {
	const maxRand = math.MaxInt
	const minRand = 0
	id := rand.Intn(maxRand-minRand) + minRand

	closeCh := make(chan struct{})
	notifyCancelOrCloseChan := make(chan error)

	d.subscribersMux.Lock()
	d.subscribers[id] = dispatchSubscriber{
		notifyCancelOrCloseChan: notifyCancelOrCloseChan,
		closeCh:                 closeCh,
	}
	d.subscribersMux.Unlock()

	go func(id int) {
		<-closeCh
		d.subscribersMux.Lock()
		defer d.subscribersMux.Unlock()
		sub, ok := d.subscribers[id]
		if !ok {
			return
		}
		close(sub.notifyCancelOrCloseChan)
		delete(d.subscribers, id)
	}(id)
	return notifyCancelOrCloseChan, closeCh
}
