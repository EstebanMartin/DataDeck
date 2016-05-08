package main

func main() {
	di := DataBaseController{}
	qr := JsonGenerator{data: di}
	qm := HttpRequestManager{responder: qr}
	qm.StartListen()
}
