package models

import (
	"fmt"

	"github.com/prometheus/common/log"
)

type Blockcount struct {
	Model

	Blockcount int  `json:"blockcount"`
	IsRecorded bool `json:"isRecorded"`
}

// GetRecordsCount returns if there is an record in the database
func GetRecordsCount() (int, error) {
	var blockcount Blockcount
	var count int
	fmt.Println("GetRecordsCount")
	if err := db.Find(blockcount).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// CreateOrUpdateRecordsCount returns if successfully create or update the records
func CreateOrUpdateRecordsCount(updatedBlockCount int) (bool, error) {
	var blockcount Blockcount
	recordsCount, err := GetRecordsCount()
	fmt.Println(recordsCount)
	if err != nil {
		return false, err
	} else if recordsCount == 0 {
		// create blockcount
		blockcount := Blockcount{Blockcount: updatedBlockCount, IsRecorded: false}
		db.Create(&blockcount)
		return true, nil
	} else {
		err := db.Find(blockcount).Select([]string{"blockcount"}).First(&blockcount)
		if err != nil {
			log.Fatalln(err)
		}
		return true, nil
		// if result != updatedBlockCount {
		// 	blockcount.Blockcount = updatedBlockCount
		// 	blockcount.IsRecorded = false
		// 	db.Save(&blockcount)
		// 	return true, nil
		// }
	}
}
