package service

import (
	"errors"
	"sync"

	"github.com/iagoholekdev/blockchain-go/domains/models"
)

const (
	ErrCouldNotSendMsg = "ErrCouldNotSendMsg"
)

type LocalTransport struct {
	addr      models.NetAddr
	consumeCh chan models.RPC
	lock      sync.RWMutex
	peers     map[models.NetAddr]*LocalTransport
}

func NewLocalTransport(addr models.NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan models.RPC, 1024),
		peers:     make(map[models.NetAddr]*LocalTransport),
	}
}

func (l *LocalTransport) Consume() <-chan models.RPC {
	return l.consumeCh
}

func (l *LocalTransport) Connect(tr models.Transport) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.peers[tr.Addr()] = tr.(*LocalTransport)
	return nil
}

func (l *LocalTransport) Addr() models.NetAddr {
	return l.addr
}

func (l *LocalTransport) SendMessage(to models.NetAddr, payload []byte) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	peers, ok := l.peers[to]
	if !ok {
		return errors.New(ErrCouldNotSendMsg)
	}
	peers.consumeCh <- models.RPC{
		From:    l.addr,
		Payload: payload,
	}
	return nil
}
