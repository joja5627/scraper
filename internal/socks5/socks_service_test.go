package socks5

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestSoxService(t *testing.T) {

	socks5Service := &Service{}
	var urls = []string{}
	for {
		fmt.Println("rotating proxies....")
		urls = socks5Service.RotateServers(10)
		time.Sleep(10 * time.Second)
	}

	resp, err := http.Get("http://example.com/")

}
