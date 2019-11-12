package main

import (
	"github.com/armon/go-socks5"
)

type runnable func() error

//func (s *Server) ListenAndServe(network, addr string) error

func stoppableChan(url string, server *socks5.Server) {

	func(u string, s *socks5.Server) {

		if err := server.ListenAndServe("tcp", u); err != nil {
			panic(err)
		}

		//s.ListenAndServe("tcp", u)
	}(url, server)

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				go fn()
			}
		}
	}()
	// …
	close(quit)
}
