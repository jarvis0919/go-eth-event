package core

import (
	"context"
	"event/global"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

func SyncEvent() {
	var startBlock = 176416
	// var  endBlock = 15204533;
	var limit = 5000

	var i = startBlock
	var ctx = context.Background()
	var contract = common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	// var topics = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	for {
		var toBlock = i + limit
		logs, err := global.EthClient.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(i)),
			ToBlock:   big.NewInt(int64(toBlock)),
			Addresses: []common.Address{
				contract,
			},

			// Topics: [][]common.Hash{
			// 	{
			// 		topics,
			// 	},
			// },
		})

		for _, log := range logs {
			fmt.Println("log", log.Topics)
		}
		fmt.Println("err", err)
		time.Sleep(time.Second * 1)

	}
}
