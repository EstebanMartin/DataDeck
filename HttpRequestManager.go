package main

import "goji.io"
import "net/http"
import "goji.io/pat"
import "golang.org/x/net/context"

type HttpRequestManager struct {
	responder QueryResponder
}

func (rm HttpRequestManager) songSearchHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var m = make(map[string]string)
	m["name"] = r.URL.Query().Get("name")
	m["genre"] = r.URL.Query().Get("genre")
	m["artist"] = r.URL.Query().Get("artist")
	response := rm.responder.RespondSearchSongs(m)
	w.Write(response)
}

func (rm HttpRequestManager) StartListen() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/search/songs.json"), rm.songSearchHandler)

	http.ListenAndServe("localhost:8000", mux)
}
