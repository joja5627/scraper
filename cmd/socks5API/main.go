package main

import (
	"fmt"
	"time"
)

//func serve(url string,server *socks5.Server){
//	if err := server.ListenAndServe("tcp",url ); err != nil {
//		panic(err)
//	}
//}
//type Client struct {
//	//Conn *websocket.Conn
//	//MessageChannel chan SocketMessage
//	TerminateChannel chan bool
//	[string]
//	//Collector *colly.Collector
//}
//func Start(url string, server *socks5.Server){
//
//}

func main(){
	//socks5.ListenAndServe()
	//conf := &socks5.Config{}
	//runtime.GOMAXPROCS(1)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//
	//	killChan := map[string] *socks5.Server{}
	//
	//	for i := 1200; i < 1220; i++ {
	//		soxUrl := fmt.Sprintf("0.0.0.0:%d",i)
	//		fmt.Println(soxUrl)
	//		server, err := socks5.New(conf)
	//		if err != nil {
	//			panic(err)
	//		}
	//		killChan[soxUrl] = server
	//		go server.ListenAndServe("tcp",soxUrl)
	//
	//	}
	//	for i := 1200; i < 1220; i++ {
	//		soxUrl := fmt.Sprintf("0.0.0.0:%d",i)
	//		time.Sleep(10*time.Second)
	//		fmt.Println("killing ", soxUrl)
	//		killChan[soxUrl].Kill()
	//
	//	}
	//
	//}()
	//wg.Wait()

	//func Kill(){
	//	<-sock5KILL
	//} curl --socks5-hostname localhost:1080 https://www.google.com/
	shutdown := make(chan bool,1)
	go waitingFunc(shutdown)
	time.Sleep(time.Second*5)
	<-shutdown



}

func waitingFunc(shutdown chan bool){
	 func() {
		defer func() {
			shutdown <- true
		}()
		for {
			select {
				case <- shutdown :
					fmt.Println("closing")
					default:
						time.Sleep(time.Second*2)
						fmt.Println("running")

				}



		}
	}()
}





