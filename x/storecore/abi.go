package storecore

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	StorecoreAbi abi.ABI
	Topics       []common.Hash
)

func init() {
	var err error
	StorecoreAbi, err = abi.JSON(strings.NewReader(StorecoreMetaData.ABI))
	if err != nil {
		panic(err)
	}

	Topics = []common.Hash{
		GetEventID(SetRecordEventID),
		GetEventID(SpliceStaticDataEventID),
		GetEventID(SpliceDynamicDataEventID),
		GetEventID(DeleteRecordEventID),
	}
}

func GetEventID(eventID string) common.Hash {
	return StorecoreAbi.Events[eventID].ID
}
