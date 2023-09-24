package etherscan_test

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/joshklop/etherscan"
	"github.com/joshklop/etherscan/etherscantest"
	"github.com/stretchr/testify/require"
)

func TestABI(t *testing.T) {
	client := etherscan.New("0xdeadbeef", etherscan.WithURL(etherscantest.New(t)))
	got, err := client.ABI(context.Background(), "0x2B3B750f1f10c85c8A6D476Fc209A8DC7E4Ca3F8")
	require.NoError(t, err)
	// TODO: this is too big to inline. Figure out a better way.
	want := []byte(`[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"fromAddress","type":"uint256"},{"indexed":true,"internalType":"address","name":"toAddress","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"payload","type":"uint256[]"}],"name":"ConsumedMessageToL1","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"fromAddress","type":"address"},{"indexed":true,"internalType":"uint256","name":"toAddress","type":"uint256"},{"indexed":true,"internalType":"uint256","name":"selector","type":"uint256"},{"indexed":false,"internalType":"uint256[]","name":"payload","type":"uint256[]"},{"indexed":false,"internalType":"uint256","name":"nonce","type":"uint256"}],"name":"ConsumedMessageToL2","type":"event"},{"anonymous":false,"inputs":[],"name":"Finalized","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"fromAddress","type":"uint256"},{"indexed":true,"internalType":"address","name":"toAddress","type":"address"},{"indexed":false,"internalType":"uint256[]","name":"payload","type":"uint256[]"}],"name":"LogMessageToL1","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"fromAddress","type":"address"},{"indexed":true,"internalType":"uint256","name":"toAddress","type":"uint256"},{"indexed":true,"internalType":"uint256","name":"selector","type":"uint256"},{"indexed":false,"internalType":"uint256[]","name":"payload","type":"uint256[]"},{"indexed":false,"internalType":"uint256","name":"nonce","type":"uint256"}],"name":"LogMessageToL2","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"acceptedGovernor","type":"address"}],"name":"LogNewGovernorAccepted","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"nominatedGovernor","type":"address"}],"name":"LogNominatedGovernor","type":"event"},{"anonymous":false,"inputs":[],"name":"LogNominationCancelled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"operator","type":"address"}],"name":"LogOperatorAdded","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"operator","type":"address"}],"name":"LogOperatorRemoved","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"removedGovernor","type":"address"}],"name":"LogRemovedGovernor","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"stateTransitionFact","type":"bytes32"}],"name":"LogStateTransitionFact","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"globalRoot","type":"uint256"},{"indexed":false,"internalType":"int256","name":"blockNumber","type":"int256"}],"name":"LogStateUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"fromAddress","type":"address"},{"indexed":true,"internalType":"uint256","name":"toAddress","type":"uint256"},{"indexed":true,"internalType":"uint256","name":"selector","type":"uint256"},{"indexed":false,"internalType":"uint256[]","name":"payload","type":"uint256[]"},{"indexed":false,"internalType":"uint256","name":"nonce","type":"uint256"}],"name":"MessageToL2Canceled","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"fromAddress","type":"address"},{"indexed":true,"internalType":"uint256","name":"toAddress","type":"uint256"},{"indexed":true,"internalType":"uint256","name":"selector","type":"uint256"},{"indexed":false,"internalType":"uint256[]","name":"payload","type":"uint256[]"},{"indexed":false,"internalType":"uint256","name":"nonce","type":"uint256"}],"name":"MessageToL2CancellationStarted","type":"event"},{"inputs":[{"internalType":"uint256","name":"toAddress","type":"uint256"},{"internalType":"uint256","name":"selector","type":"uint256"},{"internalType":"uint256[]","name":"payload","type":"uint256[]"},{"internalType":"uint256","name":"nonce","type":"uint256"}],"name":"cancelL1ToL2Message","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"configHash","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"fromAddress","type":"uint256"},{"internalType":"uint256[]","name":"payload","type":"uint256[]"}],"name":"consumeMessageFromL2","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"finalize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"identify","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"bytes","name":"data","type":"bytes"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"isFinalized","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isFrozen","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"name":"isOperator","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"msgHash","type":"bytes32"}],"name":"l1ToL2MessageCancellations","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"l1ToL2MessageNonce","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"msgHash","type":"bytes32"}],"name":"l1ToL2Messages","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"msgHash","type":"bytes32"}],"name":"l2ToL1Messages","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"messageCancellationDelay","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"programHash","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOperator","type":"address"}],"name":"registerOperator","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"toAddress","type":"uint256"},{"internalType":"uint256","name":"selector","type":"uint256"},{"internalType":"uint256[]","name":"payload","type":"uint256[]"}],"name":"sendMessageToL2","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"},{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"newConfigHash","type":"uint256"}],"name":"setConfigHash","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"delayInSeconds","type":"uint256"}],"name":"setMessageCancellationDelay","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"newProgramHash","type":"uint256"}],"name":"setProgramHash","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"starknetAcceptGovernance","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"starknetCancelNomination","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"user","type":"address"}],"name":"starknetIsGovernor","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newGovernor","type":"address"}],"name":"starknetNominateNewGovernor","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"governorForRemoval","type":"address"}],"name":"starknetRemoveGovernor","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"toAddress","type":"uint256"},{"internalType":"uint256","name":"selector","type":"uint256"},{"internalType":"uint256[]","name":"payload","type":"uint256[]"},{"internalType":"uint256","name":"nonce","type":"uint256"}],"name":"startL1ToL2MessageCancellation","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"stateBlockNumber","outputs":[{"internalType":"int256","name":"","type":"int256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"stateRoot","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"removedOperator","type":"address"}],"name":"unregisterOperator","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256[]","name":"programOutput","type":"uint256[]"},{"internalType":"uint256","name":"onchainDataHash","type":"uint256"},{"internalType":"uint256","name":"onchainDataSize","type":"uint256"}],"name":"updateState","outputs":[],"stateMutability":"nonpayable","type":"function"}]`)
	require.Equal(t, json.RawMessage(want), got)
}
