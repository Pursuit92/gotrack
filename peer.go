package gotrack

import (
	"time"
)

type Peer struct {
	PeerBencoder
	Uploaded int
	Downloaded int
	Left int
	Last time.Time
	Complete bool
}

type PeerBencoder struct {
	Id   string "peer id"
	IP   string "ip"
	Port int "port"
}

func PeerFromAnnounce(a Announce) Peer {
	peer := Peer{PeerBencoder: PeerBencoder{Id: a.PeerId,
		IP: a.IP,
		Port: a.Port,
	},
		Uploaded: a.Uploaded,
		Downloaded: a.Downloaded,
		Left: a.Left,
	}
	return peer
}
