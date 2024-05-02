package rabbitmq

import (
	"errors"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type ChannelManager struct {
	logger               *logrus.Logger
	channel              *amqp.Channel
	connManager          *ConnectionManager
	channelMux           *sync.RWMutex
	reconnectInterval    time.Duration
	reconnectionCount    uint
	reconnectionCountMux *sync.Mutex
	dispatcher           *Dispatcher
}

func NewChannelManager(connManager *ConnectionManager, log *logrus.Logger, reconnectInterval time.Duration) (*ChannelManager, error) {
	ch, err := getNewChannel(connManager)
	if err != nil {
		return nil, err
	}

	chanManager := ChannelManager{
		logger:               log,
		connManager:          connManager,
		channel:              ch,
		channelMux:           &sync.RWMutex{},
		reconnectInterval:    reconnectInterval,
		reconnectionCount:    0,
		reconnectionCountMux: &sync.Mutex{},
		dispatcher:           NewDispatcher(),
	}
	go chanManager.startNotifyCancelOrClosed()
	return &chanManager, nil
}

func getNewChannel(connManager *ConnectionManager) (*amqp.Channel, error) {
	conn := connManager.CheckoutConnection()
	defer connManager.CheckinConnection()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (chanManager *ChannelManager) startNotifyCancelOrClosed() {
	notifyCloseChan := chanManager.channel.NotifyClose(make(chan *amqp.Error, 1))
	notifyCancelChan := chanManager.channel.NotifyCancel(make(chan string, 1))

	select {
	case err := <-notifyCloseChan:
		if err != nil {
			chanManager.logger.Errorf("attempting to reconnect to amqp server after close with error: %v", err)
			chanManager.reconnectLoop()
			chanManager.logger.Warnf("successfully reconnected to amqp server")
			chanManager.dispatcher.Dispatch(err)
		}
		if err == nil {
			chanManager.logger.Infof("amqp channel closed gracefully")
		}
	case err := <-notifyCancelChan:
		chanManager.logger.Errorf("attempting to reconnect to amqp server after cancel with error: %s", err)
		chanManager.reconnectLoop()
		chanManager.logger.Warnf("successfully reconnected to amqp server after cancel")
		chanManager.dispatcher.Dispatch(errors.New(err))
	}
}

func (chanManager *ChannelManager) GetReconnectionCount() uint {
	chanManager.reconnectionCountMux.Lock()
	defer chanManager.reconnectionCountMux.Unlock()
	return chanManager.reconnectionCount
}

func (chanManager *ChannelManager) incrementReconnectionCount() {
	chanManager.reconnectionCountMux.Lock()
	defer chanManager.reconnectionCountMux.Unlock()
	chanManager.reconnectionCount++
}

func (chanManager *ChannelManager) reconnectLoop() {
	for {
		chanManager.logger.Infof("waiting %s seconds to attempt to reconnect to amqp server", chanManager.reconnectInterval)
		time.Sleep(chanManager.reconnectInterval)
		err := chanManager.reconnect()
		if err != nil {
			chanManager.logger.Errorf("error reconnecting to amqp server: %v", err)
		} else {
			chanManager.incrementReconnectionCount()
			go chanManager.startNotifyCancelOrClosed()
			return
		}
	}
}

func (chanManager *ChannelManager) reconnect() error {
	chanManager.channelMux.Lock()
	defer chanManager.channelMux.Unlock()
	newChannel, err := getNewChannel(chanManager.connManager)
	if err != nil {
		return err
	}

	if err = chanManager.channel.Close(); err != nil {
		chanManager.logger.Warnf("error closing channel while reconnecting: %v", err)
	}

	chanManager.channel = newChannel
	return nil
}

func (chanManager *ChannelManager) Close() error {
	chanManager.logger.Infof("closing channel manager...")
	chanManager.channelMux.Lock()
	defer chanManager.channelMux.Unlock()

	err := chanManager.channel.Close()
	if err != nil {
		return err
	}

	return nil
}

func (chanManager *ChannelManager) NotifyReconnect() (<-chan error, chan<- struct{}) {
	return chanManager.dispatcher.AddSubscriber()
}
