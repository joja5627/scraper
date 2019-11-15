package socks5

import (
	"fmt"
	"github.com/phayes/freeport"
)

type Service struct {
	Socks5Servers map[string]*Server
}

func getLocalURL() string {
	port, err := freeport.GetFreePort()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("0.0.0.0:%d", port)
}

func startServer(server *Server, url string) {
	server.ListenAndServe("tcp", url)
}
func (s *Service) addAndStartServer(url string) {
	conf := &Config{}
	server, err := New(conf)
	if err != nil {
		panic(err)
	}
	s.Socks5Servers[url] = server
	go startServer(server, url)

}

func (s *Service) RemoveServer(url string) {
	s.Socks5Servers[url].Kill()
}
func (s *Service) RotateServers(newServerCount int) []string {
	var serverURLS = []string{}
	for k, _ := range s.Socks5Servers {
		s.Socks5Servers[k].Kill()
	}
	s.Socks5Servers = map[string]*Server{}

	for i := 0; i < newServerCount; i++ {
		localURL := getLocalURL()
		serverURLS = append(serverURLS, localURL)
		s.addAndStartServer(localURL)
	}
	return serverURLS

}
