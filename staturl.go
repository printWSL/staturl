package staturl

import (
	"context"
	"strings"
	"sync"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const name = "staturl"

// done in 1.0
type Staturl struct {
	Next     plugin.Handler
	mu       sync.Mutex
	counters map[string]uint64
}

var (
	// prometheus metrics
	domainQueries = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "domain_queries_total",
		Help: "Total number of queries per domain.",
	}, []string{"domain"})
)

func (s *Staturl) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	name := state.Name()

	if strings.HasSuffix(name, ".cluster.local") {
		return plugin.NextOrFailure(s.Name(), s.Next, ctx, w, r)
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counters[name]++
	domainQueries.WithLabelValues(name).Inc()

	return plugin.NextOrFailure(s.Name(), s.Next, ctx, w, r)

}
func (su *Staturl) Name() string { return name }
