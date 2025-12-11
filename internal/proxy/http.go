package proxy

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"steam-lancache/internal/config"
	"steam-lancache/internal/stats"
)

func Start(state *config.AppState) {
	targetURL, _ := url.Parse("http://" + config.CacheDomain)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// force connection to cache server ip to prevent dns loop
	proxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: func(ctx context.Context, network, _ string) (net.Conn, error) {
			target := net.JoinHostPort(state.CacheServerIP, config.HTTPPort)
			return net.Dial("tcp", target)
		},
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Header.Set("Real-Host", req.Host)
		req.Host = config.CacheDomain
		req.URL.Host = config.CacheDomain
		req.URL.Scheme = "http"
		req.Header.Set("Auth-Token", state.UserToken)

		log.Printf("[HTTP] Proxying %s -> %s", req.Header.Get("Real-Host"), config.CacheDomain)
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			if resp.Request.Header.Get("User-Agent") != "GamingServices" {
				host := resp.Request.Header.Get("Real-Host")
				category := stats.DetectCategory(host)
				if resp.ContentLength > 0 {
					stats.Add(category, resp.ContentLength)
				}
			}
		}
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("[HTTP] %v", err)
		w.WriteHeader(http.StatusBadGateway)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Printf("[HTTP] Proxy listening on -> %s", config.HTTPPort)
	if err := http.ListenAndServe(":"+config.HTTPPort, nil); err != nil {
		log.Fatalf("Failed to start HTTP server -> %v", err)
	}
}
