package snapshotdb

import (
	"math/big"

	"Phoenix-Chain-Core/libs/common"
)

const WalKeyPrefix = "journal-"

func EncodeWalKey(blockNum *big.Int) []byte {
	return append([]byte(WalKeyPrefix), blockNum.Bytes()...)
}

func DecodeWalKey(key []byte) *big.Int {
	return new(big.Int).SetBytes(key[len([]byte(WalKeyPrefix)):])
}

type journalData struct {
	Key, Value []byte
}

type blockWal struct {
	ParentHash  common.Hash
	BlockHash   common.Hash
	BlockNumber *big.Int `rlp:"nil"`
	KvHash      common.Hash
	Data        []journalData
}

func (s *snapshotDB) loopWriteWal() {
	for {
		select {
		case block := <-s.walCh:
			if err := s.writeWal(block); err != nil {
				logger.Error("asynchronous write Journal fail", "err", err, "block", block.Number, "hash", block.BlockHash.String())
				s.dbError = err
				s.walSync.Done()
				continue
			}
			nc := newCurrent(block.Number, nil, block.BlockHash)
			if err := nc.saveCurrentToBaseDB(CurrentHighestBlock, s.baseDB, false); err != nil {
				logger.Error("asynchronous update current highest fail", "err", err, "block", block.Number, "hash", block.BlockHash.String())
				s.dbError = err
				s.walSync.Done()
				continue
			}
			s.walSync.Done()
		case <-s.walExitCh:
			logger.Info("loopWriteWal exist")
			close(s.walCh)
			return
		}
	}
}

func (s *snapshotDB) writeBlockToWalAsynchronous(block *blockData) {
	s.walSync.Add(1)
	s.walCh <- block
}

func (s *snapshotDB) writeWal(block *blockData) error {
	return s.baseDB.Put(block.BlockKey(), block.BlockVal(), nil)
}
