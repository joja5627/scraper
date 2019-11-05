package scrape

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
)
// "log"
// "net/http"

// "github.com/PuerkitoBio/goquery"
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
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("div.reply-info.js-only", func(e *colly.HTMLElement) {
		e.DOM.RemoveClass("display:none;")
		fmt.Println("remove", e.DOM.Text())
		e.DOM.AddClass("display:block;")
		// ch := e.DOM.Children()
		// ch.ForEach("table tr", func(_ int, el *colly.HTMLElement) {
		// 	mail := Mail{
		// 		Title:   el.ChildText("td:nth-of-type(1)"),
		// 		Link:    el.ChildAttr("td:nth-of-type(1)", "href"),
		// 		Author:  el.ChildText("td:nth-of-type(2)"),
		// 		Date:    el.ChildText("td:nth-of-type(3)"),
		// 		Message: el.ChildText("td:nth-of-type(4)"),
		// 	}
		// 	threads[threadSubject] = append(threads[threadSubject], mail)
		// })
		// fmt.Println(e.DOM.Text())
		// fmt.Println(e.DOM.Find("a[href]").Text())
		// fmt.Println(e.DOM.Find("a.mailapp").Text())
		
	})
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	
	c.Visit(link)
	c.Wait()
	// Load the HTML document
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // document, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
	// // if err != nil {
	// // 	fmt.Errorf("Error loading HTTP response body. ", err)
	// // }
	// fmt.Printf("%s", doc.Text())
	// doc.Find(".reply-info.js-only").Each(func(i int, s *goquery.Selection) {
	// 	// band := s.Find("a").Text()
	// 	// title := s.Find("i").Text()
	// 	fmt.Printf("%s", s.Text())
	// })

	return ""
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
