package gotrack

import (
	"net/http"
	"strconv"
)

const (
	Started string = "started"
	Completed = "completed"
	Stopped = "stopped"
)

type Announce struct {
	InfoHash   string
	PeerId     string
	IP         string
	Port       int
	Uploaded   int
	Downloaded int
	Left       int
	Event      string
	Numwant    int
	NoPeerId   string
	Compact    string
}

func grabFirstOrErr(v map[string][]string,s string,err error) (string,error) {
	i,ok := v[s]
	if ok && len(i) >= 1 {
		return i[0],nil
	} else {
		return "",ErrMissingValue(s)
	}
}
func grabIntOrErr(v map[string][]string,s string,err error) (int,error) {
	i,err2 := grabFirstOrErr(v,s,err)
	if err2 != nil {
		return 0,err2
	}
	return strconv.Atoi(i)
}

func ParseRequest(r *http.Request) (*Announce,error) {
	err := r.ParseForm()
	if err != nil {
		return nil,err
	}
	a := &Announce{}
	values := r.Form
	a.InfoHash,err = grabFirstOrErr(values,"info_hash",ErrMissingInfo)
	if err != nil {
		return nil,err
	}
	if len(a.InfoHash) != 20 {
		return nil,ErrInvalidInfo
	}
	a.PeerId,err = grabFirstOrErr(values,"peer_id",ErrMissingPeer)
	if err != nil {
		return nil,err
	}
	if len(a.PeerId) != 20 {
		return nil,ErrInvalidPeer
	}
	a.IP,_ = grabFirstOrErr(values,"ip",ErrGeneric)
	a.Port,err = grabIntOrErr(values,"port",ErrMissingPort)
	if err != nil {
		return nil,err
	}
	a.Uploaded,err = grabIntOrErr(values,"uploaded",ErrGeneric)
	if err != nil {
		return nil,err
	}
	a.Downloaded,err = grabIntOrErr(values,"downloaded",ErrGeneric)
	if err != nil {
		return nil,err
	}
	a.Left,err = grabIntOrErr(values,"left",ErrGeneric)
	if err != nil {
		return nil,err
	}
	a.Event,err = grabFirstOrErr(values,"event",ErrGeneric)
	a.Numwant,err = grabIntOrErr(values,"numwant",ErrGeneric)

	return a,err
}
