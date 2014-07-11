package gotrack

func IPFromRemote(s string) string {
	var i int
	var v rune
	for i,v = range s {
		if v == ':'{
			break
		}
	}
	return s[:i]
}
