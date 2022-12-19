package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{
		port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Request")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(int(ws.Port())), nil))
}
