package main

type DataInteractor interface {
	FilterSongs(parameters map[string]string) []*Song
}
