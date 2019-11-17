package socks5

import (
	"fmt"
	"github.com/phayes/freeport"
)

type Service struct {
	Socks5Servers   map[string]*Server
	RotatingServers bool
}

func getLocalURL() string {
	port, err := freeport.GetFreePort()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("socks5://localhost:%d", port)
}

func (s *Service) addAndStartServer(url string) {
	conf := &Config{}
	server, err := New(conf)
	if err != nil {
		panic(err)
	}
	s.Socks5Servers[url] = server
	server.ListenAndServe("tcp", url)

}

func (s *Service) RemoveServer(url string) {
	s.Socks5Servers[url].Kill()
}
func (s *Service) RotateServers(newServerCount int) []string {
	s.RotatingServers = true
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
	s.RotatingServers = false
	return serverURLS

}
