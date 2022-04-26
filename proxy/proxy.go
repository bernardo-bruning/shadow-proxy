package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"shadowproxy/domain"
	"shadowproxy/replication"
)

type Proxy struct {
	Url     *url.URL
	proxy   *httputil.ReverseProxy
	replica replication.Replica
	filter  Filter
}

type Filter = func(*http.Request) bool

func NewProxy(replica replication.Replica, urlString string, filter Filter) (*Proxy, error) {
	u, err := url.Parse(urlString)

	proxy := httputil.NewSingleHostReverseProxy(u)
	return &Proxy{Url: u, proxy: proxy, replica: replica, filter: filter}, err
}

func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	if !p.filter(r) {
		req, _ := domain.FromHttpRequest(r)
		err := p.replica.Emit(r.Context(), req)

		if err != nil {
			log.Fatalf("failed replication: %s\n", err.Error())
		}
	}

	p.proxy.ServeHTTP(w, r)
}

func (p *Proxy) ListenAndServe(port string) {
	http.HandleFunc("/", p.Handle)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
