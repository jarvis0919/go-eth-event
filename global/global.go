package global

import (
	"crypto/ecdsa"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

var (
	LOG *zap.Logger

	EthClient  *ethclient.Client
	ChainID    *big.Int
	FromAddr   common.Address
	PrivateKey *ecdsa.PrivateKey
)

// func SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
// 	return types.SignTx(tx, types.NewEIP155Signer(chainID), privatekey)
// }
