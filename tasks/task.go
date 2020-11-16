package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/service/kafka_service"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

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

	//  提交到kafka
	reqBodyToNode, errToNode := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "submitData",
		"id":      1,
		"params": map[string]interface{}{
			"key":   "getblockcount",
			"value": blockNumber,
		},
	})

	if errToNode != nil {
		log.Fatalln(err)
	}

	errSubmit := kafka_service.Produce(reqBodyToNode)
	if errSubmit != nil {
		log.Fatalln(errSubmit)
	}

	// currently we don't need http request to submit data(instead we use kafka)
	// respToSync, errToSync := http.Post("http://192.168.99.99:3001", "application/json", bytes.NewBuffer(reqBodyToNode))
	// if err != nil {
	// 	log.Fatalln(errToSync)
	// }
	// var result map[string]interface{}
	// json.NewDecoder(respToSync.Body).Decode(&result)

	return nil
}
