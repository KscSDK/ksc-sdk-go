package nosql

import (
	"encoding/json"
	"fmt"
)

const (
	AK = "ak"
	SK = "sk"
)

var Region = "region"

func ProcessResult(resp *map[string]interface{}, err error) {
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		return
	}

	if resp != nil {
		result, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("json marshal error", err)
			return
		}

		fmt.Println(string(result))
	}
}
