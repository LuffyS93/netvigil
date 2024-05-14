package tix

import (
	"net"
	"time"

	"github.com/cakturk/go-netstat/netstat"
	"github.com/saurlax/netvigil/netvigil"
)

type Local struct {
	Blacklist []net.IP
}

func (t *Local) Check(netstats []netstat.SockTabEntry) []netvigil.Record {
	if len(t.Blacklist) == 0 {
		return nil
	}
	records := make([]netvigil.Record, 0)
	for _, e := range netstats {
	Loop:
		for _, banned := range t.Blacklist {
			if e.RemoteAddr.IP.Equal(banned) {
				records = append(records, netvigil.Record{
					Time:       time.Now().UnixMilli(),
					LocalAddr:  e.LocalAddr.String(),
					RemoteAddr: e.RemoteAddr.String(),
					TIX:        "Local",
					Reason:     "Blacklisted",
					Executable: e.Process.Name,
					Risk:       netvigil.Malicious,
					Confidence: netvigil.High,
				})
				break Loop
			}
		}
	}
	return records
}
