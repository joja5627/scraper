package scrape

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
	"github.com/joja5627/scraper/internal/utils"
	"regexp"
	"strings"
)

var (
	collector = colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(false),
	)
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Conn *websocket.Conn
	TerminateChannel chan bool
}
// WriteSocket
func (c *Client) WriteSocket(statesQueue utils.Queue) {
	completionCounter := 0
	Loop:
		for{
			select {
				case <- c.TerminateChannel:
					break Loop
				default:
					if(statesQueue.Len() > 0){
						completionCounter = completionCounter + 1
						percentComplete := fmt.Sprintf("%f", (float64(completionCounter) / float64(statesQueue.Len())) * 100)
						c.Conn.WriteJSON(SocketMessage{MessageType:"listingPercentComplete",Payload:percentComplete})
						c.scrapeState(fmt.Sprintf("%v", statesQueue.Remove().Value))
					}else{
						c.TerminateChannel <- true
					}
				}
		}


}


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
func (c *Client) scrapeState(state string){
	stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
	clQuery := fmt.Sprintf("%s/search/sof?employment_type=3",stateOrg)
	collector.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
		listingURL := e.Attr("href")
		splitURL := strings.Split(listingURL, "/")
		title := splitURL[len(splitURL) - 2]
		listing :=  Listing{StateCode: stateOrg , Url:listingURL, Title:title}

		b, err := json.Marshal(listing)
		if err != nil {
			fmt.Println("error:", err)
		}
		c.Conn.WriteJSON(SocketMessage{MessageType:"listing",Payload:string(b)})

	})
	collector.Visit(clQuery)
}
func GetContactInfoURL(listing Listing) string  {
	info := ""
	collector.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
		info = e.Attr("data-href")
		r, _ := regexp.Compile("https://([a-z]+).craigslist.org")
		contactInfoTop := fmt.Sprintf("%s/contactinfo/",r.FindString(listing.Url))
		info = strings.Replace(info, "/__SERVICE_ID__/", contactInfoTop, 1)

	})
	collector.Visit(listing.Url)
	collector.Wait()
	return info
}
//3e2212cac1fb397d8ea0e7681e5a9db7@job.craigslist.org
//https://atlanta.craigslist.org/atl/sof/d/marietta-software-project-manager/7004489537.html
//https://auburn.craigslist.org/contactinfo/atl/sof/7004489537