package staturl

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/log"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

const name = "staturl"

type Staturl struct {
	Next plugin.Handler
}

func (su Staturl) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}

	qname := state.Name()

	log.Info(qname)

	return dns.RcodeSuccess, nil

}
func (su Staturl) Name() string { return name }
