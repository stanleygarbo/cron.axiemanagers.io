package util

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyPrint(v interface{}) {
	empJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}