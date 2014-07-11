package gotrack

import (
	"github.com/Pursuit92/syncmap"
	"time"
)

type PeerMap struct {
	internal syncmap.Map
}

type TorrentMap struct {
	internal syncmap.Map
}

func NewPeerMap() PeerMap {
	return PeerMap{syncmap.New()}
}

func NewTorrentMap() TorrentMap {
	return TorrentMap{syncmap.New()}
}

func (t TorrentMap) CreateTorrent(i string) (PeerMap) {
	pm := NewPeerMap()
	t.internal.Set(i,pm)
	return pm
}

func (t TorrentMap) DeleteTorrent(i string) {
	t.internal.Delete(i)
}

func (t TorrentMap) SetPeer(i string,p Peer) {
	pm,ok := t.internal.Get(i)
	if ok {
		pm.(PeerMap).Set(p.Id,p)
	} else {
		pm := t.CreateTorrent(i)
		pm.Set(p.Id,p)
	}
}

func (t TorrentMap) DeletePeer(i string,p Peer) {
	pm,ok := t.GetPeerMap(i)
	if ok {
		pm.Delete(p.Id)
		if pm.Size() == 0 {
			t.DeleteTorrent(i)
		}
	}
}

func (t TorrentMap) GetPeerMap(i string) (PeerMap,bool) {
	pm,ok := t.internal.Get(i)
	if ok {
		return pm.(PeerMap),ok
	} else {
		return PeerMap{},ok
	}
}

func (t TorrentMap) GetPeer(i,p string) (Peer,bool) {
	pm,ok := t.GetPeerMap(i)
	if ok {
		peer,ok := pm.Get(p)
		if ok {
			return peer,ok
		}
	}
	return Peer{},false
}

func (t TorrentMap) GetPeers(i string) ([]Peer,bool) {
	pm,ok := t.GetPeerMap(i)
	if ok {
		peerMap := pm.Map()
		peers := make([]Peer,len(peerMap))
		i := 0
		for _,v := range peerMap {
			peers[i] = v
			i++
		}
		return peers,true
	}
	return nil,false
}

func (m PeerMap) Get(s string) (Peer,bool) {
	ret,ok := m.internal.Get(s)
	if ok {
		return ret.(Peer),ok
	} else {
		return Peer{},ok
	}
}

func (m PeerMap) Size() int {
	return m.internal.Size()
}

func (m PeerMap) Set(s string,p Peer) {
	m.internal.Set(s,p)
}

func (m PeerMap) Delete(s string) {
	m.internal.Delete(s)
}

func (m PeerMap) Map() map[string]Peer {
	m.internal.Lock()
	interMap := m.internal.Map()
	m.internal.Unlock()
	pm := make(map[string]Peer)
	for i,v := range interMap {
		pm[i.(string)] = v.(Peer)
	}

	return pm
}

func (t TorrentMap) Map() map[string]PeerMap {
	t.internal.Lock()
	interMap := t.internal.Map()
	t.internal.Unlock()
	tm := make(map[string]PeerMap)
	for i,v := range interMap {
		tm[i.(string)] = v.(PeerMap)
	}
	return tm
}

func (t TorrentMap) Prune(interval time.Duration) {
	tm := t.Map()
	for i,v := range tm {
		pm := v.Map()
		for _,w := range pm {
			if time.Since(w.Last) > interval {
				t.DeletePeer(i,w)
			}
		}
	}
}
