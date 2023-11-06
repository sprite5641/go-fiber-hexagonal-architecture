package utils

import (
	"encoding/json"
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func Debug(data any) {
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(Red+"Error marshaling data:"+Reset, err)
		return
	}
	fmt.Println(Green + string(bytes) + Reset)
}

func Output(data any) []byte {
	bytes, _ := json.Marshal(data)
	return bytes
}
