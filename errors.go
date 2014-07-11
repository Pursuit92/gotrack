package gotrack

import "errors"
/*
100 	Invalid request type: client request was not a HTTP GET.
101 	Missing info_hash.
102 	Missing peer_id.
103 	Missing port.
150 	Invalid infohash: infohash is not 20 bytes long.
151 	Invalid peerid: peerid is not 20 bytes long.
152 	Invalid numwant. Client requested more peers than allowed by tracker.
200 	info_hash not found in the database. Sent only by trackers that do not automatically include new hashes into the database.
500 	Client sent an eventless request before the specified time.
900 	Generic error. 
*/

func ErrMissingValue(s string) error {
	return  errors.New("Missing parameter: " + s)
}

func ErrInvalidValue(s string) error {
	return  errors.New("Invalid parameter: " + s)
}

var (
	ErrMissingInfo = ErrMissingValue("info_hash")
	ErrMissingPeer = ErrMissingValue("peer_id")
	ErrMissingPort = ErrMissingValue("port")
	ErrInvalidInfo = ErrInvalidValue("info_hash")
	ErrInvalidPeer = ErrInvalidValue("peer_id")
	ErrInfoNotFound = errors.New("info_hash not found in database")
	ErrInvalidRequest = errors.New("Invalid Request")
	ErrEventless = errors.New("Eventless request")
	ErrGeneric = errors.New("Generic Error")
)
