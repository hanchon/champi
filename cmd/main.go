package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sort"
	"time"

	"github.com/bocha-io/champi/x/storecore"
	"github.com/bocha-io/ethclient/x/ethclient"
	"github.com/bocha-io/logger"
	"github.com/cockroachdb/pebble"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	helloStoreEvent             = "event HelloStore(bytes32 indexed storeVersion)"
	storeSetRecordEvent         = "event Store_SetRecord(bytes32 indexed tableId, bytes32[] keyTuple, bytes staticData, bytes32 encodedLengths, bytes dynamicData)"
	storeSpliceStaticDataEvent  = "event Store_SpliceStaticData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, bytes data)"
	storeSpliceDynamicDataEvent = "event Store_SpliceDynamicData(bytes32 indexed tableId, bytes32[] keyTuple, uint48 start, uint40 deleteCount, bytes32 encodedLengths, bytes data)"
	storeDeleteRecordEvent      = "event Store_DeleteRecord(bytes32 indexed tableId, bytes32[] keyTuple)"
)

func createDb() *pebble.DB {
	db, err := pebble.Open("/tmp/pebble.db", &pebble.Options{})
	if err != nil {
		log.Fatal("could not open database")
	}
	return db
}
func main() {
	c := ethclient.NewClient(context.Background(), "http://localhost:8545", 5)
	Process(c)
}
func SchemaTableId() string {
	return "0x" + common.Bytes2Hex(append(RightPadId("mudstore"), RightPadId("schema")...))
}

func MetadataTableId() string {
	return "0x" + common.Bytes2Hex(append(RightPadId("mudstore"), RightPadId("StoreMetadata")...))
}

func SchemaTableName() string {
	return "mudstore__schema"
}

func MetadataTableName() string {
	return "mudstore__storemetadata"
}

func RightPadId(id string) []byte {
	return common.RightPadBytes([]byte(id), 16)
}

func PaddedTableId(id [32]byte) string {
	return "0x" + common.Bytes2Hex(id[:])
}

// "tbstoreTables"
var STORE_TABLES = [32]byte{116, 98, 115, 116, 111, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 84, 97, 98, 108, 101, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// type Metadata struct {
// 	tableID              string
// 	fieldLayout          string
// 	keySchema            string
// 	valueSchema          string
// 	abiEncodedKeyNames   string
// 	abiEncodedValueNames string
// }
//
// type Key struct {
// 	world string
// 	storecore.TableName
// }
//
// type Database struct {
// 	Tables map[Key]Table
// }
//
// type Field struct {
// 	dataType storecore.SchemaType
// 	name     string
// }
//
// type Schema struct {
// 	staticFields  []Field
// 	dynamicFields []Field
// }
//
// type Table struct {
// 	Schema Schema
// 	Data   map[string][]interface{}
// }
//
// func NewDatabase() *Database {
// 	return &Database{
// 		Tables: map[Key]Table{},
// 	}
// }

// func Process(client *ethclient.EthClient, database *data.Database, quit *bool, startingHeight uint64, sleepDuration time.Duration) {
func Process(client *ethclient.EthClient) {
	// tables := Tables{Tables: []Table{}}
	// metadatas := Metadatas{tables: map[string]Metadata{}}
	// db := createDb()
	// defer db.Close()
	db := storecore.NewDatabase()

	startingHeight := uint64(0)
	sleepDuration := time.Duration(1 * time.Second)
	logger.LogInfo("indexer is starting...")
	// database.ChainID = client.ChainID().String()

	height := client.BlockNumber()

	endHeight := height
	amountOfBlocks := uint64(500)

	if height > startingHeight+amountOfBlocks {
		endHeight = startingHeight + amountOfBlocks
	}

	logs := client.FilterLogs(QueryForStoreLogs(big.NewInt(int64(startingHeight)), big.NewInt(int64(endHeight))))
	logs = OrderLogs(logs)
	// fmt.Println(MetadataTableId())
	// fmt.Println(SchemaTableId())
	// fmt.Println("-----")
	for _, v := range logs {
		// fmt.Println(v.TxHash.Hex())
		// for _, t := range v.Topics {
		// 	a, _ := hex.DecodeString(strings.ReplaceAll(t.String(), "0x", ""))
		// 	fmt.Println(string(a))
		// }

		// var logMudEvent data.MudEvent
		// fmt.Println(k)

		if v.Topics[0] == storecore.GetEventID(storecore.SetRecordEventID) {
			event, err := storecore.ParseStoreSetRecord(v)
			if err != nil {
				logger.LogError(fmt.Sprintf("[indexer] error decoding message:%s", err))
				// TODO: what should we do here?
				break
			}
			storecore.HandleStoreSetRecord(db, event)
		} else if v.Topics[0] == storecore.GetEventID(storecore.DeleteRecordEventID) {
			event, err := storecore.ParseStoreDeleteRecord(v)
			if err != nil {
				logger.LogError(fmt.Sprintf("[indexer] error decoding message:%s", err))
				// TODO: what should we do here?
				break
			}
			fmt.Println("DELETE RECORD EVENT ------")
			tableName := storecore.KeyToTableName(event.TableId)
			fmt.Println(tableName)
		} else if v.Topics[0] == storecore.GetEventID(storecore.SpliceStaticDataEventID) {
			event, err := storecore.ParseStoreSpliceStaticData(v)
			if err != nil {
				logger.LogError(fmt.Sprintf("[indexer] error decoding message:%s", err))
				// TODO: what should we do here?
				break
			}
			fmt.Println("Splice static RECORD EVENT ------")

			tableName := storecore.KeyToTableName(event.TableId)
			fmt.Println(tableName)
			storecore.HandleStoreSpliceStaticData(db, event)
		} else if v.Topics[0] == storecore.GetEventID(storecore.SpliceDynamicDataEventID) {
			event, err := storecore.ParseStoreSpliceDynamicData(v)
			if err != nil {
				logger.LogError(fmt.Sprintf("[indexer] error decoding message:%s", err))
				// TODO: what should we do here?
				break
			}
			fmt.Println("Splice dynamic RECORD EVENT ------")

			tableName := storecore.KeyToTableName(event.TableId)
			fmt.Println(tableName)
		} else {
			panic("asd")
		}

		//       else {
		// 	fmt.Println("qwe")
		// 	fmt.Println(string(event.TableId[:]))
		// 	fmt.Println(event.TableId[:])
		// 	fmt.Println(STORE_TABLES)
		// 	fmt.Println(event.TableId == STORE_TABLES)
		// 	table := PaddedTableId(event.TableId)
		// 	fmt.Println(table)
		// 	tablename, _ := hex.DecodeString(strings.ReplaceAll(table, "0x", ""))
		// 	fmt.Println("tablename")
		// 	fmt.Println(string(tablename))
		// 	for _, v := range event.KeyTuple {
		// 		fmt.Println(v)
		// 		fmt.Println(string(hex.EncodeToString(v[:])))
		// 		fmt.Println(string(v[:]))
		// 	}
		// 	if table == MetadataTableId() {
		// 		panic("metadata")
		// 	}
		// 	if table == SchemaTableId() {
		// 		panic("schema")
		// 	}
		// 	panic("stop")
		// }

		// switch mudhelpers.PaddedTableId(event.TableId) {
		// case mudhelpers.SchemaTableId():
		// 	logger.LogInfo("[indexer] processing and creating schema table")
		// 	mudhandlers.HandleSchemaTableEvent(event, db)
		// case mudhelpers.MetadataTableId():
		// 	logger.LogInfo("[indexer] processing and updating a schema with metadata")
		// 	mudhandlers.HandleMetadataTableEvent(event, db)
		// default:
		// 	logger.LogInfo("[indexer] processing a generic table event like adding a row")
		// 	logMudEvent = mudhandlers.HandleGenericTableEvent(event, db)
		// }

		// if v.Topics[0].Hex() == mudhelpers.GetStoreAbiEventID("StoreSetField").Hex() {
		// 	event, err := mudhandlers.ParseStoreSetField(v)
		// 	logger.LogInfo("[indexer] processing store set field message")
		// 	if err != nil {
		// 		logger.LogError(fmt.Sprintf("[indexer] error decoding message for store set field:%s\n", err))
		// 	} else {
		// 		logMudEvent = mudhandlers.HandleSetFieldEvent(event, db)
		// 	}
		// }
		// if v.Topics[0].Hex() == mudhelpers.GetStoreAbiEventID("StoreDeleteRecord").Hex() {
		// 	logger.LogInfo("[indexer] processing store delete record message")
		// 	event, err := mudhandlers.ParseStoreDeleteRecord(v)
		// 	if err != nil {
		// 		logger.LogError(fmt.Sprintf("[indexer] error decoding message for store delete record:%s\n", err))
		// 	} else {
		// 		logMudEvent = mudhandlers.HandleDeleteRecordEvent(event, db)
		// 	}
		// }

	}

	// eth.ProcessBlocks(client, database, big.NewInt(int64(startingHeight)), big.NewInt(int64(endHeight)))
	fmt.Println("db")
	fmt.Println(db)
	storecore.PrintDebug(db)
	quit := true

	for !quit {
		newHeight := client.BlockNumber()

		if newHeight != endHeight {
			startingHeight = endHeight
			endHeight = newHeight

			if newHeight > startingHeight+amountOfBlocks {
				endHeight = startingHeight + amountOfBlocks
			}

			logger.LogInfo(fmt.Sprintf("Heights: %d %d", startingHeight, endHeight))

			// eth.ProcessBlocks(client, database, big.NewInt(int64(startingHeight)), big.NewInt(int64(endHeight)))
		}

		// database.LastHeight = newHeight

		time.Sleep(sleepDuration)

	}
}

func OrderLogs(logs []types.Log) []types.Log {
	// Filter removed logs due to chain reorgs.
	filteredLogs := []types.Log{}
	for _, log := range logs {
		if !log.Removed {
			filteredLogs = append(filteredLogs, log)
		}
	}

	// Order logs.
	sort.SliceStable(filteredLogs, func(i, j int) bool {
		first := filteredLogs[i]
		second := filteredLogs[j]
		if first.BlockNumber < second.BlockNumber {
			return true
		}
		if second.BlockNumber < first.BlockNumber {
			return false
		}
		return first.Index < second.Index
	})

	return filteredLogs
}

func QueryForStoreLogs(initBlockHeight *big.Int, endBlockHeight *big.Int) ethereum.FilterQuery {
	if initBlockHeight == nil {
		initBlockHeight = big.NewInt(1)
	}

	// TODO: we should query the blockchain to get the latest block
	if endBlockHeight == nil {
		endBlockHeight = big.NewInt(999999999)
	}
	fmt.Println(storecore.Topics)

	return ethereum.FilterQuery{
		FromBlock: initBlockHeight,
		ToBlock:   endBlockHeight,
		Topics: [][]common.Hash{
			storecore.Topics,
		},
	}
}
