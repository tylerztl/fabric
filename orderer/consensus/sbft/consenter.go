/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sbft

import (
	"github.com/hyperledger/fabric/orderer/consensus"
	"github.com/hyperledger/fabric/orderer/consensus/migration"
	"github.com/hyperledger/fabric/orderer/consensus/sbft/backend"
	"github.com/hyperledger/fabric/orderer/consensus/sbft/connection"
	"github.com/hyperledger/fabric/orderer/consensus/sbft/persist"
	"github.com/hyperledger/fabric/orderer/consensus/sbft/simplebft"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/op/go-logging"
)

type consensusStack struct {
	persist *persist.Persist
	backend *backend.Backend
}

var logger = logging.MustGetLogger("orderer/consensus/sbft")

// Consenter interface implementation for new main application
type consenter struct {
	config          *ConsensusConfig
	consensusStack  *consensusStack
	sbftStackConfig *backend.StackConfig
	sbftPeers       map[string]*simplebft.SBFT
}

type chain struct {
	chainID         string
	exitChan        chan struct{}
	consensusStack  *consensusStack
	migrationStatus migration.Status
}

// New creates a new consenter for the SBFT consensus scheme.
// It accepts messages being delivered via Enqueue, orders them, and then uses the blockcutter to form the messages
// into blocks before writing to the given ledger.
func New(c *ConsensusConfig, sc *backend.StackConfig) consensus.Consenter {
	return &consenter{config: c, sbftStackConfig: sc}
}

func (sbft *consenter) HandleChain(support consensus.ConsenterSupport, metadata *cb.Metadata) (consensus.Chain, error) {
	return newChain(sbft, support), nil
}

func newChain(sbft *consenter, support consensus.ConsenterSupport) *chain {
	logger.Infof("Starting a chain: %d", support.ChainID())

	if sbft.sbftPeers == nil {
		sbft.consensusStack = createConsensusStack(sbft)
		sbft.sbftPeers = make(map[string]*simplebft.SBFT)
	}
	sbft.sbftPeers[support.ChainID()] = initSbftPeer(support.ChainID(), sbft, support)

	return &chain{
		chainID:         support.ChainID(),
		exitChan:        make(chan struct{}),
		consensusStack:  sbft.consensusStack,
		migrationStatus: migration.NewStatusStepper(support.IsSystemChannel(), support.ChainID()), // Needed by consensus-type migration
	}
}

func createConsensusStack(sbft *consenter) *consensusStack {
	logger.Infof("%v    %v      %v", sbft.sbftStackConfig.ListenAddr, sbft.sbftStackConfig.CertFile, sbft.sbftStackConfig.KeyFile)
	conn, err := connection.New(sbft.sbftStackConfig.ListenAddr, sbft.sbftStackConfig.CertFile, sbft.sbftStackConfig.KeyFile)
	if err != nil {
		logger.Errorf("Error when trying to connect: %s", err)
		panic(err)
	}
	persist := persist.New(sbft.sbftStackConfig.DataDir)
	backend, err := backend.NewBackend(sbft.config.Peers, conn, persist)
	if err != nil {
		logger.Errorf("Backend instantiation error.")
		panic(err)
	}

	go conn.Server.Serve(conn.Listener)

	return &consensusStack{
		backend: backend,
		persist: persist,
	}
}

func initSbftPeer(chainID string, sbft *consenter, support consensus.ConsenterSupport) *simplebft.SBFT {
	sbftPeer, err := sbft.consensusStack.backend.AddSbftPeer(support.ChainID(), support, sbft.config.Consensus)
	if err != nil {
		logger.Errorf("SBFT peer instantiation error.")
		panic(err)
	}
	return sbftPeer
}

// Chain interface implementation:

// Start allocates the necessary resources for staying up to date with this Chain.
// It implements the multichain.Chain interface. It is called by multichain.NewManagerImpl()
// which is invoked when the ordering process is launched, before the call to NewServer().
func (ch *chain) Start() {

}

// Halt frees the resources which were allocated for this Chain
func (ch *chain) Halt() {
	// panic("There is no way to halt SBFT")
	select {
	case <-ch.exitChan:
		// Allow multiple halts without panic
	default:
		close(ch.exitChan)
	}
}

// Enqueue accepts a message and returns true on acceptance, or false on shutdown
//func (ch *chain) Enqueue(env *cb.Envelope) bool {
//	return ch.consensusStack.backend.Enqueue(ch.chainID, env)
//}

func (ch *chain) WaitReady() error {
	return nil
}

// Order accepts normal messages for ordering
func (ch *chain) Order(env *cb.Envelope, configSeq uint64) error {
	return ch.consensusStack.backend.Enqueue(ch.chainID, env)
}

// Configure accepts configuration update messages for ordering
func (ch *chain) Configure(config *cb.Envelope, configSeq uint64) error {
	return ch.consensusStack.backend.Enqueue(ch.chainID, config)
}

// Errored only closes on exit
func (ch *chain) Errored() <-chan struct{} {
	return ch.exitChan
}

func (ch *chain) MigrationStatus() migration.Status {
	return ch.migrationStatus
}
