package main

import "encoding/json"

//Implements QueryResponder

type JsonGenerator struct {
	data DataInteractor
}

func (jg JsonGenerator) RespondSearchSongs(parameters map[string]string) []byte {
	songs := jg.data.FilterSongs(parameters)
	response, err := json.Marshal(songs)
	checkError(err)
	return response
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
