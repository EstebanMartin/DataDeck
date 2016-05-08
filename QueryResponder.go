package main

type QueryResponder interface {
	RespondSearchSongs(parameters map[string]string) []byte
}
