package core

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types/pbfttypes"
	"math/big"
	"testing"
	"time"

	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/common"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/db/snapshotdb"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/core/types"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/event"
)

func TestBlockChainReactorClose(t *testing.T) {
	t.Run("close after commit", func(t *testing.T) {
		eventmux := new(event.TypeMux)
		reacter := NewBlockChainReactor(eventmux, big.NewInt(100))
		reacter.Start(common.DPOS_VALIDATOR_MODE)
		var parenthash common.Hash
		pbftress := make(chan pbfttypes.PbftResult, 5)
		go func() {
			for i := 1; i < 11; i++ {
				header := new(types.Header)
				header.Number = big.NewInt(int64(i))
				header.Time = uint64(i)
				header.ParentHash = parenthash
				block := types.NewBlock(header, nil, nil)
				snapshotdb.Instance().NewBlock(header.Number, header.ParentHash, block.Hash())
				parenthash = block.Hash()
				pbftress <- pbfttypes.PbftResult{Block: block}
			}
			close(pbftress)
		}()

		for value := range pbftress {
			eventmux.Post(value)
		}

		reacter.Close()

		time.Sleep(time.Second)
		snapshotdb.Instance().Clear()
	})
}
