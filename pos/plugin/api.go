package plugin

import (
	"fmt"
	"Phoenix-Chain-Core/libs/common"
	"Phoenix-Chain-Core/ethereum/core/db/snapshotdb"
)

// Provides an API interface to obtain data related to the economic model
type PublicDPOSAPI struct {
	snapshotDB snapshotdb.DB
}

func NewPublicDPOSAPI() *PublicDPOSAPI {
	return &PublicDPOSAPI{snapshotdb.Instance()}
}

// Get node list of zero-out blocks
func (p *PublicDPOSAPI) GetWaitSlashingNodeList() string {
	list, err := slash.getWaitSlashingNodeList(0, common.ZeroHash)
	if nil != err || len(list) == 0 {
		return ""
	}
	return fmt.Sprintf("%+v", list)
}
