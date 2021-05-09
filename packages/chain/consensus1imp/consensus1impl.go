package consensus1imp

import (
	"time"

	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/chain"
	"go.uber.org/atomic"
)

type consensusImpl struct {
	isReady                   atomic.Bool
	chain                     chain.ChainCore
	committee                 chain.Committee
	mempool                   chain.Mempool
	nodeConn                  chain.NodeConnection
	log                       *logger.Logger
	eventStateTransitionMsgCh chan *chain.StateTransitionMsg
	eventTimerMsgCh           chan chain.TimerTick
	closeCh                   chan struct{}
}

var _ chain.Consensus1 = &consensusImpl{}

func New(chainCore chain.ChainCore, mempool chain.Mempool, committee chain.Committee, nodeConn chain.NodeConnection, log *logger.Logger) *consensusImpl {
	ret := &consensusImpl{
		chain:                     chainCore,
		committee:                 committee,
		mempool:                   mempool,
		nodeConn:                  nodeConn,
		log:                       log.Named("c"),
		eventStateTransitionMsgCh: make(chan *chain.StateTransitionMsg),
		eventTimerMsgCh:           make(chan chain.TimerTick),
		closeCh:                   make(chan struct{}),
	}
	go ret.recvLoop()
	return ret
}

func (c *consensusImpl) IsReady() bool {
	return c.isReady.Load()
}

func (c *consensusImpl) Close() {
	close(c.closeCh)
}

func (c *consensusImpl) recvLoop() {
	// wait at startup
	for !c.committee.IsReady() {
		select {
		case <-time.After(100 * time.Millisecond):
		case <-c.closeCh:
			return
		}
	}
	c.log.Infof("consensus object is ready")
	c.isReady.Store(true)
	for {
		select {
		case msg, ok := <-c.eventStateTransitionMsgCh:
			if ok {
				c.eventStateTransitionMsg(msg)
			}
		case msg, ok := <-c.eventTimerMsgCh:
			if ok {
				c.eventTimerMsg(msg)
			}
		case <-c.closeCh:
			return
		}
	}
}
