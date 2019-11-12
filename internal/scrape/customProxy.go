package scrape

import (
	"context"
	"github.com/gocolly/colly"
	"net/http"
	"net/url"
	"sync/atomic"
)

type roundRobinSwitcher struct {
	proxyURLs []*url.URL
	index     uint32
}

func (r *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	u := r.proxyURLs[r.index%uint32(len(r.proxyURLs))]
	atomic.AddUint32(&r.index, 1)
	ctx := context.WithValue(pr.Context(), colly.ProxyURLKey, u.String())
	*pr = *pr.WithContext(ctx)
	return u, nil
}

// RoundRobinProxySwitcher creates a proxy switcher function which rotates
// ProxyURLs on every request.
// The proxy type is determined by the URL scheme. "http", "https"
// and "socks5" are supported. If the scheme is empty,
// "http" is assumed.
func CustomProxy(urls []string) (colly.ProxyFunc, error) {
	proxyURLS := make([]*url.URL, len(urls))
	for i, u := range urls {
		parsedU, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		proxyURLS[i] = parsedU
	}
	return (&roundRobinSwitcher{proxyURLS, 0}).GetProxy, nil
}
