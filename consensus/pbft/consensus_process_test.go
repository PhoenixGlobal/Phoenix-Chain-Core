package pbft

import (
	"Phoenix-Chain-Core/consensus/pbft/validator"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"Phoenix-Chain-Core/consensus/pbft/wal"

	"Phoenix-Chain-Core/consensus/pbft/protocols"
	"Phoenix-Chain-Core/ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestViewChange(t *testing.T) {
	pk, sk, pbftnodes := GeneratePbftNode(4)
	nodes := make([]*TestPBFT, 0)
	for i := 0; i < 4; i++ {
		node := MockNode(pk[i], sk[i], pbftnodes, 1000, 1)
		assert.Nil(t, node.Start())

		nodes = append(nodes, node)
	}

	// TestTryViewChange
	testTryViewChange(t, nodes)

	// TestTryChangeViewByViewChange
	//testTryChangeViewByViewChange(t, nodes)
}

func testTryViewChange(t *testing.T, nodes []*TestPBFT) {
	tempDir, _ := ioutil.TempDir("", "wal")
	defer os.RemoveAll(tempDir)

	result := make(chan *types.Block, 1)
	complete := make(chan struct{}, 1)
	parent := nodes[0].chain.Genesis()
	//fmt.Println("222222222222222222222222222222,parent block is ",parent.NumberU64(),parent.Hash(),parent)

	nodes[0].engine.wal, _ = wal.NewWal(nil, tempDir)
	nodes[0].engine.bridge, _ = NewBridge(nodes[0].engine.nodeServiceContext, nodes[0].engine)

	for i := 0; i < 1; i++ {
		block := NewBlockWithSign(parent.Hash(), parent.NumberU64()+1, nodes[0])
		assert.True(t, nodes[0].engine.state.HighestExecutedBlock().Hash() == block.ParentHash())
		nodes[0].engine.OnSeal(block, result, nil, complete)
		<-complete

		_, qc := nodes[0].engine.blockTree.FindBlockAndQC(parent.Hash(), parent.NumberU64())
		select {
		case b := <-result:
			assert.NotNil(t, b)
			assert.Equal(t, uint32(i-1), nodes[0].engine.state.MaxQCIndex())
			for j := 1; j < 4; j++ {
				msg := &protocols.PreCommit{
					Epoch:          nodes[0].engine.state.Epoch(),
					ViewNumber:     nodes[0].engine.state.ViewNumber(),
					BlockIndex:     uint32(0),
					BlockHash:      b.Hash(),
					BlockNumber:    b.NumberU64(),
					ValidatorIndex: uint32(j),
					ParentQC:       qc,
				}
				assert.Nil(t, nodes[j].engine.signMsgByBls(msg))
				assert.Nil(t, nodes[0].engine.OnPreCommit("id", msg), fmt.Sprintf("number:%d", b.NumberU64()))
			}
			parent = b
		}
	}
	//time.Sleep(1 * time.Second)

	block := nodes[0].engine.state.HighestPreCommitQCBlock()
	//fmt.Println("HighestPreCommitQCBlock,block is ",block.NumberU64(),block.Hash(),block)
	block, qc := nodes[0].engine.blockTree.FindBlockAndQC(block.Hash(), block.NumberU64())
	//fmt.Println("FindBlockAndQC,block is ",block.NumberU64(),block.Hash(),block)
	//fmt.Println("11111111111111111111111,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber(),"blockNumber is ",block.NumberU64())
	for i := 0; i < 4; i++ {
		epoch, view := nodes[0].engine.state.Epoch(), nodes[0].engine.state.ViewNumber()
		viewchange := &protocols.ViewChange{
			Epoch:          epoch,
			ViewNumber:     view,
			BlockHash:      block.Hash(),
			BlockNumber:    block.NumberU64(),
			ValidatorIndex: uint32(i),
			PrepareQC:      qc,
		}
		assert.Nil(t, nodes[i].engine.signMsgByBls(viewchange))
		assert.Nil(t, nodes[0].engine.OnViewChanges("id", &protocols.ViewChanges{
			VCs: []*protocols.ViewChange{
				viewchange,
			},
		}))
	}
	assert.NotNil(t, nodes[0].engine.state.LastViewChangeQC())
	assert.Equal(t, uint64(1), nodes[0].engine.state.ViewNumber())
	//fmt.Println("222222222222222222,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber(),",nodes[0].engine.state.BlockNumber() is ",nodes[0].engine.state.BlockNumber(),",nodes[0].engine.state.Epoch() is ",nodes[0].engine.state.Epoch())
	lastViewChangeQC, _ := nodes[0].engine.bridge.GetViewChangeQC(nodes[0].engine.state.Epoch(), nodes[0].engine.state.BlockNumber()-1,nodes[0].engine.state.ViewNumber())
	//fmt.Println("first lastViewChangeQC is ",lastViewChangeQC)
	assert.Nil(t, lastViewChangeQC)
	lastViewChangeQC, _ = nodes[0].engine.bridge.GetViewChangeQC(nodes[0].engine.state.Epoch(), nodes[0].engine.state.BlockNumber()-1,nodes[0].engine.state.ViewNumber()-1)
	//fmt.Println("second lastViewChangeQC is ",lastViewChangeQC)
	assert.NotNil(t, lastViewChangeQC)
	epoch, viewNumber, _, _, _, blockNumber := lastViewChangeQC.MaxBlock()
	assert.Equal(t, nodes[0].engine.state.Epoch(), epoch)
	assert.Equal(t, nodes[0].engine.state.ViewNumber()-1, viewNumber)
	assert.Equal(t, uint64(1), blockNumber)
}

func testTryChangeViewByViewChange(t *testing.T, nodes []*TestPBFT) {
	// note: node-0 has been successfully switched to view-1, HighestQC blockNumber = 4
	// build a duplicate block-4
	fmt.Println("-------------------------------------------------")
	number, hash := nodes[0].engine.HighestQCBlockBn()
	block, _ := nodes[0].engine.blockTree.FindBlockAndQC(hash, number)
	dulBlock := NewBlock(block.ParentHash(), block.NumberU64())
	_, preQC := nodes[0].engine.blockTree.FindBlockAndQC(block.ParentHash(), block.NumberU64()-1)
	// Vote and generate prepareQC for dulBlock
	votes := make(map[uint32]*protocols.PrepareVote)
	for j := 1; j < 3; j++ {
		vote := &protocols.PrepareVote{
			Epoch:          nodes[0].engine.state.Epoch(),
			ViewNumber:     nodes[0].engine.state.ViewNumber(),
			BlockIndex:     uint32(0),
			BlockHash:      dulBlock.Hash(),
			BlockNumber:    dulBlock.NumberU64(),
			ValidatorIndex: uint32(j),
			ParentQC:       preQC,
		}
		assert.Nil(t, nodes[j].engine.signMsgByBls(vote))
		votes[uint32(j)] = vote
	}
	dulQC := nodes[0].engine.generatePrepareQC(votes)
	// build new viewChange
	viewChanges := make(map[uint32]*protocols.ViewChange)
	for j := 1; j < 3; j++ {
		viewchange := &protocols.ViewChange{
			Epoch:          nodes[0].engine.state.Epoch(),
			ViewNumber:     nodes[0].engine.state.ViewNumber(),
			BlockHash:      dulBlock.Hash(),
			BlockNumber:    dulBlock.NumberU64(),
			ValidatorIndex: uint32(j),
			PrepareQC:      dulQC,
		}
		assert.Nil(t, nodes[j].engine.signMsgByBls(viewchange))
		viewChanges[uint32(j)] = viewchange
	}
	viewChangeQC := nodes[0].engine.generateViewChangeQC(viewChanges)

	// Case1: local highestqc is behind other validators and not exist viewChangeQC.maxBlock, sync qc block
	fmt.Println("11111,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber())
	nodes[0].engine.tryChangeViewByViewChange(viewChangeQC)
	fmt.Println("22222,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber())
	assert.Equal(t, uint64(1), nodes[0].engine.state.ViewNumber())
	assert.Equal(t, hash, nodes[0].engine.state.HighestQCBlock().Hash())

	// Case2: local highestqc is equal other validators and exist viewChangeQC.maxBlock, change the view
	nodes[0].engine.insertQCBlock(dulBlock, dulQC)
	fmt.Println("333333,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber())
	nodes[0].engine.tryChangeViewByViewChange(viewChangeQC)
	fmt.Println("444444,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber())
	assert.Equal(t, uint64(2), nodes[0].engine.state.ViewNumber())
	assert.Equal(t, dulQC.BlockHash, nodes[0].engine.state.HighestQCBlock().Hash())

	// based on the view-2 build a duplicate block-4
	dulBlock = NewBlock(block.ParentHash(), block.NumberU64())
	// Vote and generate prepareQC for dulBlock
	votes = make(map[uint32]*protocols.PrepareVote)
	for j := 1; j < 3; j++ {
		vote := &protocols.PrepareVote{
			Epoch:          nodes[0].engine.state.Epoch(),
			ViewNumber:     nodes[0].engine.state.ViewNumber(),
			BlockIndex:     uint32(0),
			BlockHash:      dulBlock.Hash(),
			BlockNumber:    dulBlock.NumberU64(),
			ValidatorIndex: uint32(j),
			ParentQC:       preQC,
		}
		assert.Nil(t, nodes[j].engine.signMsgByBls(vote))
		votes[uint32(j)] = vote
	}
	dulQC = nodes[0].engine.generatePrepareQC(votes)
	fmt.Println("1111111111555555,before nodes[0].engine.blockTree.InsertQCBlock(dulBlock, dulQC)")
	nodes[0].engine.blockTree.InsertQCBlock(dulBlock, dulQC)
	fmt.Println("2222222222555555,after nodes[0].engine.blockTree.InsertQCBlock(dulBlock, dulQC)")
	nodes[0].engine.state.SetHighestQCBlock(dulBlock)
	// Case3: local highestqc is ahead other validators, and not send viewChange, generate new viewChange quorumCert and change the view
	fmt.Println("555555,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber())
	nodes[0].engine.tryChangeViewByViewChange(viewChangeQC)
	fmt.Println("666666,nodes[0].engine.state.ViewNumber() is ",nodes[0].engine.state.ViewNumber())
	assert.Equal(t, uint64(3), nodes[0].engine.state.ViewNumber())
	assert.Equal(t, dulQC.BlockHash, nodes[0].engine.state.HighestQCBlock().Hash())
	_, _, _, blockView, _, _ := viewChangeQC.MaxBlock()
	assert.Equal(t, uint64(0), blockView)
}

type testCase struct {
	hadViewTimeout bool
}

func TestRichViewChangeQC(t *testing.T) {
	tests := []testCase{
		{true},
		{false},
	}
	for _, c := range tests {
		testRichViewChangeQCCase(t, c)
	}
}

func testRichViewChangeQCCase(t *testing.T, c testCase) {
	pk, sk, pbftnodes := GeneratePbftNode(4)
	nodes := make([]*TestPBFT, 0)
	for i := 0; i < 4; i++ {
		node := MockNode(pk[i], sk[i], pbftnodes, 1000, 1)
		assert.Nil(t, node.Start())

		nodes = append(nodes, node)
	}

	result := make(chan *types.Block, 1)
	complete := make(chan struct{}, 1)
	parent := nodes[0].chain.Genesis()
	for i := 0; i < 1; i++ {
		block := NewBlockWithSign(parent.Hash(), parent.NumberU64()+1, nodes[0])
		assert.True(t, nodes[0].engine.state.HighestExecutedBlock().Hash() == block.ParentHash())
		nodes[0].engine.OnSeal(block, result, nil, complete)
		<-complete

		_, qc := nodes[0].engine.blockTree.FindBlockAndQC(parent.Hash(), parent.NumberU64())
		select {
		case b := <-result:
			assert.NotNil(t, b)
			assert.Equal(t, uint32(i-1), nodes[0].engine.state.MaxQCIndex())
			for j := 1; j < 3; j++ {
				msg := &protocols.PrepareVote{
					Epoch:          nodes[0].engine.state.Epoch(),
					ViewNumber:     nodes[0].engine.state.ViewNumber(),
					BlockIndex:     uint32(i),
					BlockHash:      b.Hash(),
					BlockNumber:    b.NumberU64(),
					ValidatorIndex: uint32(j),
					ParentQC:       qc,
				}
				assert.Nil(t, nodes[j].engine.signMsgByBls(msg))
				assert.Nil(t, nodes[0].engine.OnPrepareVote("id", msg), fmt.Sprintf("number:%d", b.NumberU64()))
			}
			parent = b
		}
	}
	if c.hadViewTimeout {
		time.Sleep(1 * time.Second)
	}

	hadSend := nodes[0].engine.state.ViewChangeByIndex(0)
	if c.hadViewTimeout {
		assert.NotNil(t, hadSend)
	}
	if !c.hadViewTimeout {
		assert.Nil(t, hadSend)
	}

	qcBlock := nodes[0].engine.state.HighestQCBlock()
	//fmt.Println(1111111111111,"qcBlock is ",qcBlock)
	_, qc := nodes[0].engine.blockTree.FindBlockAndQC(qcBlock.Hash(), qcBlock.NumberU64())
	//fmt.Println("qc is ",qc)
	lockBlock := nodes[0].engine.state.HighestLockBlock()
	//fmt.Println(222222222222,"lockBlock is ",lockBlock)
	_, lockqc := nodes[0].engine.blockTree.FindBlockAndQC(lockBlock.Hash(), lockBlock.NumberU64())
	//fmt.Println("lockqc is ",lockqc)

	viewChanges := make(map[uint32]*protocols.ViewChange, 0)

	for i := 1; i < 4; i++ {
		epoch, view := nodes[0].engine.state.Epoch(), nodes[0].engine.state.ViewNumber()
		v := &protocols.ViewChange{
			Epoch:          epoch,
			ViewNumber:     view,
			BlockHash:      lockBlock.Hash(), // base lock qc
			BlockNumber:    lockBlock.NumberU64(),
			ValidatorIndex: uint32(i),
			PrepareQC:      lockqc,
		}
		assert.Nil(t, nodes[i].engine.signMsgByBls(v))
		viewChanges[v.ValidatorIndex] = v
	}

	viewChangeQC := nodes[0].engine.generateViewChangeQC(viewChanges)
	//fmt.Println(333333333333,"viewChangeQC is ",viewChangeQC)
	nodes[0].engine.richViewChangeQC(viewChangeQC)
	//fmt.Println(444444444444)

	epoch, viewNumber, blockEpoch, blockViewNumber, blockHash, blockNumber := viewChangeQC.MaxBlock()
	assert.Equal(t, qc.Epoch, epoch)
	assert.Equal(t, qc.ViewNumber, viewNumber)
	assert.Equal(t, qc.Epoch, blockEpoch)
	assert.Equal(t, qc.ViewNumber, blockViewNumber)
	assert.Equal(t, qc.BlockHash, blockHash)
	assert.Equal(t, qc.BlockNumber, blockNumber)
}

func TestViewChangeBySwitchPoint(t *testing.T) {
	pk, sk, pbftnodes := GeneratePbftNode(4)
	nodes := make([]*TestPBFT, 0)
	for i := 0; i < 4; i++ {
		node := MockNode(pk[i], sk[i], pbftnodes, 1000, 1)
		node.agency = validator.NewMockAgency(pbftnodes, 1)
		assert.Nil(t, node.Start())
		node.engine.validatorPool.MockSwitchPoint(1)
		nodes = append(nodes, node)
	}

	result := make(chan *types.Block, 1)
	complete := make(chan struct{}, 1)
	parent := nodes[0].chain.Genesis()
	for i := 0; i < 1; i++ {
		block := NewBlockWithSign(parent.Hash(), parent.NumberU64()+1, nodes[0])
		assert.True(t, nodes[0].engine.state.HighestExecutedBlock().Hash() == block.ParentHash())
		nodes[0].engine.OnSeal(block, result, nil, complete)
		<-complete

		_, qc := nodes[0].engine.blockTree.FindBlockAndQC(parent.Hash(), parent.NumberU64())
		select {
		case b := <-result:
			assert.NotNil(t, b)
			assert.Equal(t, uint32(i-1), nodes[0].engine.state.MaxQCIndex())
			pb := &protocols.PrepareBlock{
				Epoch:         nodes[0].engine.state.Epoch(),
				ViewNumber:    nodes[0].engine.state.ViewNumber(),
				Block:         b,
				BlockIndex:    uint32(i),
				ProposalIndex: uint32(0),
			}
			nodes[0].engine.signMsgByBls(pb)
			nodes[1].engine.OnPrepareBlock("id", pb)
			for j := 1; j < 4; j++ {
				msg := &protocols.PrepareVote{
					Epoch:          nodes[0].engine.state.Epoch(),
					ViewNumber:     nodes[0].engine.state.ViewNumber(),
					BlockIndex:     uint32(i),
					BlockHash:      b.Hash(),
					BlockNumber:    b.NumberU64(),
					ValidatorIndex: uint32(j),
					ParentQC:       qc,
				}
				//if j == 1 {
				//	nodes[1].engine.state.HadSendPrepareVote().Push(msg)
				//}
				assert.Nil(t, nodes[j].engine.signMsgByBls(msg))
				nodes[0].engine.OnPrepareVote("id", msg)
				if i < 9 {
					assert.Nil(t, nodes[1].engine.OnPrepareVote("id", msg), fmt.Sprintf("number:%d", b.NumberU64()))
				}
			}
			parent = b
		}
	}
	// node-0 enough 10 block qc,change the epoch
	assert.Equal(t, uint64(2), nodes[0].engine.state.Epoch())
	assert.Equal(t, uint64(0), nodes[0].engine.state.ViewNumber())

	// node-1 change the view base lock block
	lockBlock := nodes[0].engine.state.HighestQCBlock()
	_, lockQC := nodes[0].engine.blockTree.FindBlockAndQC(lockBlock.Hash(), lockBlock.NumberU64())
	for i := 0; i < 4; i++ {
		epoch, view := nodes[1].engine.state.Epoch(), nodes[1].engine.state.ViewNumber()
		viewchange := &protocols.ViewChange{
			Epoch:          epoch,
			ViewNumber:     view,
			BlockHash:      lockBlock.Hash(),
			BlockNumber:    lockBlock.NumberU64(),
			ValidatorIndex: uint32(i),
			PrepareQC:      lockQC,
		}
		assert.Nil(t, nodes[i].engine.signMsgByBls(viewchange))
		assert.Nil(t, nodes[1].engine.OnViewChanges("id", &protocols.ViewChanges{
			VCs: []*protocols.ViewChange{
				viewchange,
			},
		}))
	}
	assert.NotNil(t, nodes[1].engine.state.LastViewChangeQC())
	assert.Equal(t, uint64(1), nodes[1].engine.state.ViewNumber())

	//qcBlock := nodes[0].engine.state.HighestQCBlock()
	//_, qcQC := nodes[0].engine.blockTree.FindBlockAndQC(qcBlock.Hash(), qcBlock.NumberU64())
	//// change view by switchPoint
	//nodes[1].engine.insertQCBlock(qcBlock, qcQC)
	assert.Equal(t, uint64(2), nodes[1].engine.state.Epoch())
	assert.Equal(t, uint64(0), nodes[1].engine.state.ViewNumber())
}
