package element

import (
	"errors"
	"time"

	"github.com/hashicorp/memberlist"
)

const (
	defaultInterval      = time.Second * 10
	nodeReconcileTimeout = defaultInterval * 3
	nodeUpdateTimeout    = defaultInterval / 2
)

var (
	ErrUnknownConnectionType = errors.New("unknown connection type")
)

// Agent represents the node agent
type Agent struct {
	config             *Config
	members            *memberlist.Memberlist
	peerUpdateChan     chan bool
	nodeEventChan      chan *NodeEvent
	registeredServices map[string]struct{}
	memberConfig       *memberlist.Config
	state              *State
}

// NewAgent returns a new node agent
func NewAgent(info *Peer, cfg *Config) (*Agent, error) {
	var (
		updateCh    = make(chan bool, 64)
		nodeEventCh = make(chan *NodeEvent, 64)
	)
	a := &Agent{
		config:         cfg,
		peerUpdateChan: updateCh,
		nodeEventChan:  nodeEventCh,
		state: &State{
			Self:  info,
			Peers: make(map[string]*Peer),
		},
	}
	mc, err := cfg.memberListConfig(a)
	if err != nil {
		return nil, err
	}
	ml, err := memberlist.Create(mc)
	if err != nil {
		return nil, err
	}
	a.members = ml
	a.memberConfig = mc

	return a, nil
}

// SyncInterval returns the cluster sync interval
func (a *Agent) SyncInterval() time.Duration {
	return a.memberConfig.PushPullInterval
}

// Subscribe subscribes to the node event channel
func (a *Agent) Subscribe(ch chan *NodeEvent) {
	go func() {
		for {
			select {
			case evt := <-a.nodeEventChan:
				ch <- evt
			}
		}
	}()
}
