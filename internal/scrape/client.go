package scrape

import (
	"fmt"
	"github.com/corpix/uarand"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Conn             *websocket.Conn
	MessageChannel   chan SocketMessage
	TerminateChannel chan bool
	Collector        *colly.Collector
}

func getCollector() *colly.Collector {
	return colly.NewCollector(
		colly.MaxDepth(1),
		colly.UserAgent(uarand.GetRandom()),
		colly.Async(false),
	)
}

//func getCollectorAsync() *colly.Collector {
//	 colly.NewCollector(
//		colly.MaxDepth(1),
//		colly.UserAgent(uarand.GetRandom()),
//		colly.Async(true),
//	)
//	colly.LimitRule{&colly.LimitRule{DomainGlob: "*", Parallelism: 12}}
//
//	colly.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 12})
//	return colly
//
//}

// WriteSocket
func (c *Client) WriteSocket(stateCodes []string) {
	completionCounter := 0
	go func() {

		for {
			select {
			case socketMessage := <-c.MessageChannel:
				if socketMessage.MessageType == "state" {
					completionCounter = completionCounter + 1
					divisor := float64(len(stateCodes))
					percentComplete := float64(completionCounter) / divisor
					scaledPercent := percentComplete * 100
					c.Conn.WriteJSON(SocketMessage{MessageType: "percentComplete", Payload: fmt.Sprintf("%f", scaledPercent)})
				}

				c.Conn.WriteJSON(socketMessage)
			case <-c.TerminateChannel:
				if err := c.Conn.Close(); err != nil {
					fmt.Println(err)
				}

			default:
				{
					//if()
					//
					//c.Collector.Wait()
				}
			}
		}
	}()
}

//
//writer, _ := c.Conn.NextWriter(websocket.TextMessage)
//socketMessage, _ := json.Marshal(})
//
//writer.Write(socketMessage)
//writer.Close()
//completionCounter := 0
//Loop:
//	for{
//		select {
//			case <- c.TerminateChannel:
//				break Loop
//			default:
//
//
//			}
//	}

// ReadSocket
func (c *Client) ReadSocket() {
	for {
		messageType, _, err := c.Conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			fmt.Println("closed")
			c.TerminateChannel <- true
			break
		}

	}
}
func (c *Client) sendState(message SocketMessage) {
	c.MessageChannel <- message
}

//func (c *Client) scrapeState(state string) {
//	stateMessage  := SocketMessage{MessageType:"state",Payload:state}
//	defer c.sendState(stateMessage)
//
//	collector := getCollector()
//	stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
//
//
//
//}

//3e2212cac1fb397d8ea0e7681e5a9db7@job.craigslist.org
//https://atlanta.craigslist.org/atl/sof/d/marietta-software-project-manager/7004489537.html
//https://auburn.craigslist.org/contactinfo/atl/sof/7004489537
