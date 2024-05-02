package rabbitmq

import (
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type ConnectionManager struct {
	logger               *logrus.Logger
	url                  string
	connection           *amqp.Connection
	amqpConfig           amqp.Config
	connectionMux        *sync.RWMutex
	ReconnectInterval    time.Duration
	reconnectionCount    uint
	reconnectionCountMux *sync.Mutex
	dispatcher           *Dispatcher
}

func NewConnectionManager(url string, conf amqp.Config, log *logrus.Logger, reconnectInterval time.Duration) (*ConnectionManager, error) {
	conn, err := amqp.DialConfig(url, amqp.Config(conf))
	if err != nil {
		return nil, err
	}
	connManager := ConnectionManager{
		logger:               log,
		url:                  url,
		connection:           conn,
		amqpConfig:           conf,
		connectionMux:        &sync.RWMutex{},
		ReconnectInterval:    reconnectInterval,
		reconnectionCount:    0,
		reconnectionCountMux: &sync.Mutex{},
		dispatcher:           NewDispatcher(),
	}
	go connManager.startNotifyClose()
	return &connManager, nil
}

func (connManager *ConnectionManager) Close() error {
	connManager.logger.Infof("closing connection manager...")
	connManager.connectionMux.Lock()
	defer connManager.connectionMux.Unlock()

	err := connManager.connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (connManager *ConnectionManager) NotifyReconnect() (<-chan error, chan<- struct{}) {
	return connManager.dispatcher.AddSubscriber()
}

func (connManager *ConnectionManager) CheckoutConnection() *amqp.Connection {
	connManager.connectionMux.RLock()
	return connManager.connection
}

func (connManager *ConnectionManager) CheckinConnection() {
	connManager.connectionMux.RUnlock()
}

func (connManager *ConnectionManager) startNotifyClose() {
	notifyCloseChan := connManager.connection.NotifyClose(make(chan *amqp.Error, 1))

	err := <-notifyCloseChan
	if err != nil {
		connManager.logger.Errorf("attempting to reconnect to amqp server after connection close with error: %v", err)
		connManager.reconnectLoop()
		connManager.logger.Warnf("successfully reconnected to amqp server")
		connManager.dispatcher.Dispatch(err)
	}
	if err == nil {
		connManager.logger.Infof("amqp connection closed gracefully")
	}
}

func (connManager *ConnectionManager) GetReconnectionCount() uint {
	connManager.reconnectionCountMux.Lock()
	defer connManager.reconnectionCountMux.Unlock()
	return connManager.reconnectionCount
}

func (connManager *ConnectionManager) incrementReconnectionCount() {
	connManager.reconnectionCountMux.Lock()
	defer connManager.reconnectionCountMux.Unlock()
	connManager.reconnectionCount++
}

func (connManager *ConnectionManager) reconnectLoop() {
	for {
		connManager.logger.Infof("waiting %s seconds to attempt to reconnect to amqp server", connManager.ReconnectInterval)
		time.Sleep(connManager.ReconnectInterval)
		err := connManager.reconnect()
		if err != nil {
			connManager.logger.Errorf("error reconnecting to amqp server: %v", err)
		} else {
			connManager.incrementReconnectionCount()
			go connManager.startNotifyClose()
			return
		}
	}
}

func (connManager *ConnectionManager) reconnect() error {
	connManager.connectionMux.Lock()
	defer connManager.connectionMux.Unlock()
	newConn, err := amqp.DialConfig(connManager.url, amqp.Config(connManager.amqpConfig))
	if err != nil {
		return err
	}

	if err = connManager.connection.Close(); err != nil {
		connManager.logger.Warnf("error closing connection while reconnecting: %v", err)
	}

	connManager.connection = newConn
	return nil
}
