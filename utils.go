package main

import "encoding/json"

func jsonToGoStruct[T any](jsonStr string) (T, error) {
	var result T
	marshalErr := json.Unmarshal([]byte(jsonStr), &result)
	return result, marshalErr
}
