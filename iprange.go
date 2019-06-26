package iputils

import (
	"errors"
	"github.com/apcera/util/iprange"
)

type IPRange string

func (this IPRange) Range() (ipRange *iprange.IPRange, err error) {
	if ipRange, err = iprange.ParseIPRange(string(this)); err == nil && ipRange.Start == nil || ipRange.End == nil {
		return nil, errors.New("bad ip address")
	}
	return
}
func (this IPRange) MustRange() (ipRange *iprange.IPRange) {
	ipRange, _ = iprange.ParseIPRange(string(this))
	return
}
