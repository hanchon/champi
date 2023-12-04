package storecore

import (
	"fmt"

	"github.com/bocha-io/logger"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func ParseStoreSetRecord(log types.Log) (*StorecoreStoreSetRecord, error) {
	event := new(StorecoreStoreSetRecord)
	if err := UnpackLog(event, SetRecordEventID, log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func ParseStoreSpliceStaticData(log types.Log) (*StorecoreStoreSpliceStaticData, error) {
	event := new(StorecoreStoreSpliceStaticData)
	if err := UnpackLog(event, SpliceStaticDataEventID, log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func ParseStoreSpliceDynamicData(log types.Log) (*StorecoreStoreSpliceDynamicData, error) {
	event := new(StorecoreStoreSpliceDynamicData)
	if err := UnpackLog(event, SpliceDynamicDataEventID, log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func ParseStoreDeleteRecord(log types.Log) (*StorecoreStoreDeleteRecord, error) {
	event := new(StorecoreStoreDeleteRecord)
	if err := UnpackLog(event, DeleteRecordEventID, log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func UnpackLog(out interface{}, eventName string, log types.Log) error {
	if log.Topics[0] != StorecoreAbi.Events[eventName].ID {
		return fmt.Errorf("event signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := StorecoreAbi.UnpackIntoInterface(out, eventName, log.Data); err != nil {
			logger.LogError(fmt.Sprintf("failed to unpack into interface %s", err))
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range StorecoreAbi.Events[eventName].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func GetWorld(log types.Log) string {
	return log.Address.Hex()
}
