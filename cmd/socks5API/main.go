package main

import (
	"fmt"
	"github.com/joja5627/scraper/internal/socks5"
	"sync"
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
type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) Run() {
	for {
		select {
		case <-t.closed:
			return

		default:
			fmt.Print("#")
			time.Sleep(time.Millisecond * 200)
			//case <-t.ticker.C:
			//	handle()
			//}
		}
	}
}

func (t *Task) Stop() {
	close(t.closed)
	t.wg.Wait()
}

func handle() {
	for i := 0; i < 5; i++ {
		fmt.Print("#")
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println()
}
func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	socks5Service := &socks5.Service{}

	for {
		fmt.Println("rotating proxies....")
		urls := socks5Service.RotateServers(10)
		for url := range urls {
			fmt.Println(urls[url])
		}
		time.Sleep(10 * time.Second)
	}

	wg.Wait()

	//for i := 0; i < 10; i++ {
	//	socks5Service.
	//}
	//go func() {
	//	defer wg.Done()
	//
	//
	//
	//
	//}()
	//wg.Wait()

	//func Kill(){
	//	<-sock5KILL
	//} curl --socks5-hostname localhost:1080 https://www.google.com/
	//

	//exit := make(chan bool)
	//go server.ListenAndServe("tcp","0.0.0.0:1220",exit)
	//time.Sleep(time.Second*5)
	//exit <- true
	//task := &Task{
	//	closed: make(chan struct{}),
	//	ticker: time.NewTicker(time.Second * 2),
	//}
	//
	//c := make(chan os.Signal)
	//signal.Notify(c, os.Interrupt)

	//task.wg.Add(1)
	//go func() { defer task.wg.Done(); task.Run() }()
	//time.Sleep(time.Second * 5)
	//task.Stop()

	//server := waitingFunc(shutdown,"0.0.0.0:1220")

}

//func waitingFunc(shutdown <- chan bool,url string) *socks5.Server{

//	 func() {
//		defer func() {
//			fmt.Println("closing")
//		}()
//		for {
//			err := server.ListenAndServe("tcp",url) ;if err != nil {
//				panic(err)
//			}
//		}
//	}()
//	return server
//}
