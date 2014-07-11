package gotrack

import (
	"net/http"
	"time"
	bencode "code.google.com/p/bencode-go"
	"log"

)

type Handler struct {
	Torrents TorrentMap
	Interval time.Duration
	Timeout time.Duration
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	log.Printf("Attempt from %s",req.RemoteAddr)
	writeError := func(e error,n int) {
		http.Error(resp,e.Error(),n)
	}
	if req.Method != "GET" {
		writeError(ErrInvalidRequest,100)
		return
	}
	announce,err := ParseRequest(req)
	if err != nil && announce == nil {
		writeCode := func(code int) {
			writeError(err,code)
		}
		switch err {
		case ErrMissingInfo:
			writeCode(101)
			return
		case ErrMissingPeer:
			writeCode(102)
			return
		case ErrMissingPort:
			writeCode(103)
			return
		case ErrInvalidInfo:
			writeCode(150)
			return
		case ErrInvalidPeer:
			writeCode(151)
			return
		default:
			writeCode(900)
			return
		}
	}
	log.Printf("Announce from %s",req.RemoteAddr)
	peer := PeerFromAnnounce(*announce)
	h.Torrents.GetPeer(announce.InfoHash,announce.PeerId)
	/*
	// Check to see if the peer already exists
	if ! ok {
		if announce.Event != "started" {
			// Error, initial connection without "started"
			http.Error(resp,"First connection must have event \"started.\"",900)
			return
		}
	}
	*/
	if announce.Event == "stopped" {
		h.Torrents.DeletePeer(announce.InfoHash,peer)
		return
	}
	if announce.Event == "completed" {
		peer.Complete = true
	}
	if announce.Left == 0 {
		peer.Complete = true
	}
	peer.IP = IPFromRemote(req.RemoteAddr)
	peer.Last = time.Now()
	h.Torrents.SetPeer(announce.InfoHash,peer)

	respStruct := buildResponse(announce.InfoHash,h.Torrents,h.Interval,*announce)
	err = bencode.Marshal(resp,respStruct)
	if err != nil {
		http.Error(resp,"",500)
	}

}

func NewHandler(interval,prune string) *Handler {
	interDur,err := time.ParseDuration(interval)
	if err != nil {
		panic(err)
	}
	pruneDur,err := time.ParseDuration(prune)
	if err != nil {
		panic(err)
	}
	h := &Handler{}
	h.Torrents = NewTorrentMap()
	go func() {
		c := time.Tick(pruneDur)
		for _ = range c {
			h.Torrents.Prune(interDur)
		}
	}()
	return h
}
