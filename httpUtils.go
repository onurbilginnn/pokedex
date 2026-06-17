package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/onurbilginnn/pokecache"
)

func getFromURL[T any](url string, cache *pokecache.Cache) (T, error) {
	var zeroValue T
	entry, isEntryExists := cache.Get(url)
	if isEntryExists == true {
		data, marshalErr := jsonToGoStruct[T](string(entry))
		if marshalErr != nil {
			return zeroValue, marshalErr
		}
		return data, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return zeroValue, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zeroValue, errors.New("failed to fetch data")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zeroValue, err
	}
	data, marshalErr := jsonToGoStruct[T](string(body))
	if marshalErr != nil {
		return zeroValue, marshalErr
	}
	cache.Add(url, body)

	return data, nil
}
