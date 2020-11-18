package tasks

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// GetBlockNumber is the function of gettint block number and submit to the sync server
func GetBlockNumber() error {

	reqBodyToNode, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "getblockcount",
		"id":      1,
		"params":  []string{},
	})

	if err != nil {
		log.Fatalln(err)
	}

	respToNode, err := http.Post("http://192.168.1.36:10332", "application/json", bytes.NewBuffer(reqBodyToNode))

	if err != nil {
		log.Fatalln(err)
	}

	var resultFromNode map[string]interface{}

	json.NewDecoder(respToNode.Body).Decode(&resultFromNode)

	var blockNumber int = int(resultFromNode["result"].(float64))

	reqBodyToSync, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "updateBlockNumber",
		"id":      1,
		"params": map[string]interface{}{
			"blockNumber": blockNumber,
			"timestamp":   int(time.Now().Unix()),
		},
	})

	// 提交syncService
	respToSync, errToSync := http.Post("http://localhost:8000/mutation", "application/json", bytes.NewBuffer(reqBodyToSync))
	if err != nil {
		log.Fatalln(errToSync)
	}
	var result map[string]interface{}
	json.NewDecoder(respToSync.Body).Decode(&result)
	// fmt.Println(result["data"])
	// 先暂时不用kafka

	// if errToNode != nil {
	// 	log.Fatalln(err)
	// }

	// errSubmit := kafka_service.Produce(reqBodyToNode)
	// if errSubmit != nil {
	// 	log.Fatalln(errSubmit)
	// }

	return nil
}
