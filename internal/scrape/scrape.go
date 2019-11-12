package scrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
)

var (
	selectors        = []string{".result-row .result-image", "#sortable-results > ul > li:nth-child(1) > p > a"}
	emailTagSelector = "body > section > section > header > div.reply-button-row > button"
	c                = colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(false),
	)
)

//

//GetListingURLS comment
func GetListingURLS(stateCodes []string, con websocket.Conn) {

	for i := range stateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", stateCodes[i])
		clQuery := fmt.Sprintf("%s/search/sof", stateOrg)
		percentComplete := fmt.Sprintf("%f", (float64(i)/float64(len(stateCodes)))*100)
		con.WriteJSON(SocketMessage{MessageType: "listingPercentComplete", Payload: percentComplete})
		c.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
			listingURL := e.Attr("href")
			con.WriteJSON(SocketMessage{MessageType: "listingURLs", Payload: listingURL})

		})
		c.Visit(clQuery)

	}

}

//GetContactInfoURLS comment
func GetContactInfoURLS(link string) string {
	fmt.Println(link)
	var emailLink string
	c.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
		//doc.Find("form").Each(func(i int, formDoc *goquery.Selection) {
		//	if loginFormSelection != nil {
		//		return
		//	}
		//	formDoc.Find("input").Each(func(_ int, inputDoc *goquery.Selection) {
		//		if loginFormSelection != nil {
		//			return
		//		}
		//		if name, ok := inputDoc.Attr("name"); ok {
		//			if strings.Contains(strings.ToLower(name), "pass") {
		//				loginFormSelection = formDoc
		//			}
		//		}
		//	})
		//})
	})

	c.Visit(link)
	c.Wait()
	return emailLink
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
