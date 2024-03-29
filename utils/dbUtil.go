package utils

import (
	"encoding/json"
	"log"
)

func ConvertBsonToStruct(from interface{}, to interface{}) (error) {
	records, err := json.Marshal(from)
	if err != nil {
		log.Println("error while marshalling json", err)
		return err
	}

	err = json.Unmarshal(records, to)
	if err != nil {
		log.Println("error while unmarshalling json", err)
		return err
	}
	
	return nil
}