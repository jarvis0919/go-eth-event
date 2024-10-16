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
	startBlock := big.NewInt(43157760) // 176416
	new, _ := global.EthClient.BlockNumber(context.Background())
	newblock := big.NewInt(int64(new))
	fmt.Println(newblock)
	// var endBlock = 15204533
	limit := 10

	// var i = startBlock
	ctx := context.Background()
	contract := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	contract2 := common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d")
	toBlock := big.NewInt(0).Add(startBlock, big.NewInt(int64(limit)))
	// var topics = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	for {
		if toBlock != nil {
			i := toBlock.Cmp(newblock)
			if i == 1 || i == 0 {
				toBlock = nil
			} else {
				toBlock = big.NewInt(0).Add(startBlock, big.NewInt(int64(limit)))
			}
		}

		logs, err := global.EthClient.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: startBlock,
			ToBlock:   toBlock,
			Addresses: []common.Address{
				contract,
				contract2,
			},

			// Topics: [][]common.Hash{
			// 	{
			// 		topics,
			// 	},
			// },
		})
		if err != nil {
			fmt.Println("err", err)
			continue
		}

		fmt.Printf("startBlock: %d, endBlock: %d, len: %d\n", startBlock, toBlock, len(logs))
		for i := 0; i < len(logs); i++ {
			fmt.Println("log", logs[i])
		}
		if len(logs) == 0 {
			if toBlock != nil {
				startBlock = big.NewInt(0).Add(startBlock, big.NewInt(int64(limit)))
				continue
			} else {
				continue
			}
		}
		startBlock = big.NewInt(int64(logs[len(logs)-1].BlockNumber))

		// for _, log := range logs {
		// 	fmt.Println("log", log, len(log.Topics))
		// }
		// fmt.Println("err", err)
		time.Sleep(time.Second * 1)

	}
}
