package tool

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

func Contain(arr []string, element string) bool {
	for _, v := range arr {
		if v == element {
			return true
		}
	}
	return false
}

func MarshalJson(date interface{}) []byte {
	res, err := json.Marshal(date)

	if err != nil {
		fmt.Println(err)
	}

	return res
}

func UnmarshalJson(date string) map[string]interface{} {
	var res map[string]interface{}

	_ = json.Unmarshal([]byte(date), &res)

	return res
}

func GetRandmod(length int) int64 {
	result := int64(0)
	res, err := rand.Int(rand.Reader, big.NewInt(int64(length)))

	if err != nil {
		return result
	}

	return res.Int64()
}

func IntToString(value int64) string {
	v := strconv.FormatInt(value, 10)

	return v
}

func StringToInt(value string) int64 {
	res, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		res = 0
	}

	return res
}

func InterfaceToString(inter interface{}) (result string) {
	switch inter.(type) {
	case string:
		result = inter.(string)
	}

	return result
}
