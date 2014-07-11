package gotrack

import (
	"time"
)

type Response struct {
	Interval int "interval"
	Complete int "complete"
	Incomplete int "incomplete"
	Peers    []PeerBencoder "peers"
}

type ResponseFail struct {
	Reason string "failure reason"
}

func buildResponse(hash string,tm TorrentMap,inter time.Duration,ann Announce) Response {
	pm,_ := tm.GetPeerMap(hash)
	peers := pm.Map()
	respPeers := make([]PeerBencoder,len(peers) - 1)
	var j,complete,incomplete int
	for i,v := range peers {
		if i != ann.PeerId && j < ann.Numwant {
			respPeers[j] = v.PeerBencoder
			j++
		}
		if v.Complete {
			complete++
		} else {
			incomplete++
		}
	}

	return Response{Interval:int(inter.Seconds()),
		Complete: complete,
		Incomplete:incomplete,
		Peers: respPeers,
	}
}
