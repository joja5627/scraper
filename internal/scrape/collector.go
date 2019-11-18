package scrape

import (
	"fmt"
	"github.com/corpix/uarand"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"github.com/joja5627/scraper/internal/socks5"
	"regexp"
	"strings"
	"time"
)

var errors []string
var links []string
var contactInfos []string

func GetContactInfos() []string {
	return contactInfos
}

func getProxyFunc(proxyService *socks5.Service) (colly.ProxyFunc, error) {
	newServers := proxyService.RotateServers(30)
	fmt.Printf("rotating servers")
	proxyFunc, err := proxy.RoundRobinProxySwitcher(newServers...)
	if err != nil {
		return nil, err
	}
	return proxyFunc, nil
}

func BuildCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.IgnoreRobotsTxt())

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*.craigslist.org.*",
		Parallelism: 1,
		RandomDelay: 30 * time.Second,
	})
	//socks5Service := &socks5.Service{RotatingServers: false}
	//
	//proxyFunc, err := getProxyFunc(socks5Service)
	//if err != nil {
	//	fmt.Println("can't set proxy")
	//}
	//collector.SetProxyFunc(proxyFunc)

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("error: ", err.Error())
		fmt.Println("errors: ", len(errors))
		errors = append(errors, err.Error())
		time.Sleep(5 * time.Second)
		c.Visit(r.Request.URL.String())
	})
	collector.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {
		listingURL := e.Attr("href")
		c.Visit(listingURL)
	})
	collector.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
		info := e.Attr("data-href")
		r, _ := regexp.Compile("https://([a-z]+).craigslist.org")
		contactInfoTop := fmt.Sprintf("%s/contactinfo/", r.FindString(e.Request.URL.String()))
		info = strings.Replace(info, "/__SERVICE_ID__/", contactInfoTop, 1)
		contactInfos = append(contactInfos, info)

	})
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Del("User-Agent")
		r.Headers.Add("User-Agent", uarand.GetRandom())

	})

	collector.OnResponse(func(r *colly.Response) {

	})

	return collector
}
