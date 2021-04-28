package listener

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridgev2/chains/evm/client"
	"github.com/ChainSafe/chainbridgev2/example/keystore"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"

	"github.com/centrifuge/go-substrate-rpc-client/types"
)

func Test_byteconv(t *testing.T) {
	recS := br()
	b := []byte(recS)
	fmt.Printf("%b", b)
}

func br() types.Bytes {
	return types.Bytes{1, 1, 1, 1, 1}
}

var AliceKp = keystore.TestKeyRing.EthereumKeys[keystore.AliceKey]
var BobKp = keystore.TestKeyRing.EthereumKeys[keystore.BobKey]
var EveKp = keystore.TestKeyRing.EthereumKeys[keystore.EveKey]

var (
	DefaultRelayerAddresses = []common.Address{
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.AliceKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.BobKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.CharlieKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.DaveKey].Address()),
		common.HexToAddress(keystore.TestKeyRing.EthereumKeys[keystore.EveKey].Address()),
	}
)

func Test_gettx(t *testing.T) {
	ethClient, err := client.NewEVMClient("ws://localhost:8545", false, AliceKp)
	if err != nil {
		log.Error().Err(err).Msg("errors creating client")
		t.Fail()
	}
	_, err = Simulate(ethClient.Client, big.NewInt(420), common.HexToHash("0x23afeec87f5400749c22acd39b0d0a95f13dd832619fdbc7310fc5ab0e3a627d"), AliceKp.CommonAddress())
	//receipt, err := ethClient.Client.TransactionReceipt(context.TODO(), common.HexToHash("0x23afeec87f5400749c22acd39b0d0a95f13dd832619fdbc7310fc5ab0e3a627d"))
	if err != nil {
		log.Error().Err(err).Msg("errors simulating")
		t.Fail()
	}
	//log.Debug().Msgf("%+v", receipt)
	//log.Debug().Msg(string(res))
}

// Simulate function gets transaction info by hash and then executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain. Execution happens against provided block.
func Simulate(client *ethclient.Client, block *big.Int, txHash common.Hash, from common.Address) ([]byte, error) {
	tx, _, err := client.TransactionByHash(context.TODO(), txHash)
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}
	res, err := client.CallContract(context.TODO(), msg, block)
	if err != nil {
		return nil, err
	}
	bs, err := hex.DecodeString(common.Bytes2Hex(res))
	if err != nil {
		return nil, err
	}
	log.Debug().Msg(string(bs))
	return bs, nil
}
