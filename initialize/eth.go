package initialize

import (
	"context"
	"crypto/ecdsa"
	"event/global"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InitGlobal() {
	var err error
	global.EthClient, err = ethclient.Dial("https://bsc-mainnet.infura.io/v3/6a1e4e32ce7847db8e32a20a79ea8a63")
	if err != nil {
		panic(err)
	}
	var ctx = context.Background()
	global.ChainID, err = global.EthClient.ChainID(ctx)
	if err != nil {
		panic(err)
	}
	global.PrivateKey, err = crypto.HexToECDSA("4d5396d5e575f33d81c27e356a6b7d6b10204f2239c91acebd299bef7a7b03f5")
	if err != nil {
		panic(err)
	}
	publicKey := global.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	global.FromAddr = crypto.PubkeyToAddress(*publicKeyECDSA)

}
