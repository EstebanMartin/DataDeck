package main

type Song struct {
	Artist string `json:"artist"`
	Name   string `json:"song"`
	Genre  string `json:"genre"`
	Length int    `json:"length"`
}
