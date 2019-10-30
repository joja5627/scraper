
package scrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
	"regexp"
	"strings"
)

var (
	selectors        = []string{".result-row .result-image", "#sortable-results > ul > li:nth-child(1) > p > a"}
	emailTagSelector = "body > section > section > header > div.reply-button-row > button"
	c = colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(false),
	)
)
//

//GetListingURLS comment
func GetListingURLS(stateCodes []string,con websocket.Conn) []Listing {

	var listings []Listing

	for i := range stateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", stateCodes[i])
		clQuery := fmt.Sprintf("%s/search/sof?employment_type=3",stateOrg)
		r, _ := regexp.Compile("https://([a-z]+).craigslist.org")
		percentComplete := fmt.Sprintf("%f", (float64(i) / float64(len(stateCodes))) * 100)
		con.WriteJSON(SocketMessage{MessageType:"listingPercentComplete",Payload:percentComplete})


		c.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
			listingURL := e.Attr("href")
			listing :=  Listing{StateOrg:r.FindString(listingURL) ,ListingUrl:listingURL,QueryUrl:clQuery}
			listings = append(listings,listing )

		})
		c.Visit(clQuery)

	}

	return listings
}

//GetContactInfoURLS comment
func GetContactInfoURLS(listings []Listing,con websocket.Conn) []Listing {

	for i := range listings {
		percentComplete := fmt.Sprintf("%f", (float64(i) / float64(len(listings))) * 100)
		con.WriteJSON(SocketMessage{MessageType:"contactInfoPercentComplete",Payload:percentComplete})
		c.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
			serviceID := e.Attr("data-href")
			contactInfoUrl := strings.Replace(serviceID, "/__SERVICE_ID__", listings[i].StateOrg, -1)
			listings[i].ContactInfoUrl = contactInfoUrl
		})
		c.Visit(listings[i].ListingUrl)

	}
	return listings
}
//func GetContactInfoURLS(listings []Listing) []Listing {
//
//
//	for i := range listings {
//		c.OnHTML("body > section > section > header > div.reply-button-row > button", func(e *colly.HTMLElement) {
//			listings[i].GetContactInfoURLS = e.Attr("data-href")
//		})
//		c.Visit(listings[i].ListingUrl)
//		c.Wait()
//	}
//	return listings
//}
