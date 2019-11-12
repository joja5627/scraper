package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/websocket"
	"github.com/joja5627/scraper/internal/scrape"
	"google.golang.org/api/gmail/v1"
	"net/http"
	"net/url"
	"strings"
)

var (
	newline    = []byte{'\n'}
	space      = []byte{' '}
	stateCodes = []string{"auburn", "bham", "dothan", "shoals", "gadsden", "huntsville", "mobile", "montgomery", "tuscaloosa", "anchorage", "fairbanks", "kenai", "juneau", "flagstaff", "mohave", "phoenix", "prescott", "showlow", "sierravista", "tucson", "yuma", "fayar", "fortsmith", "jonesboro", "littlerock", "texarkana", "bakersfield", "chico", "fresno", "goldcountry", "hanford", "humboldt", "imperial", "inlandempire", "losangeles", "mendocino", "merced", "modesto", "monterey", "orangecounty", "palmsprings", "redding", "sacramento", "sandiego", "sfbay", "slo", "santabarbara", "santamaria", "siskiyou", "stockton", "susanville", "ventura", "visalia", "yubasutter", "boulder", "cosprings", "denver", "eastco", "fortcollins", "rockies", "pueblo", "westslope", "newlondon", "hartford", "newhaven", "nwct", "delaware", "washingtondc", "miami", "daytona", "keys", "fortlauderdale", "fortmyers", "gainesville", "cfl", "jacksonville", "lakeland", "miami", "lakecity", "ocala", "okaloosa", "orlando", "panamacity", "pensacola", "sarasota", "miami", "spacecoast", "staugustine", "tallahassee", "tampa", "treasure", "miami", "albanyga", "athensga", "atlanta", "augusta", "brunswick", "columbusga", "macon", "nwga", "savannah", "statesboro", "valdosta", "honolulu", "boise", "eastidaho", "lewiston", "twinfalls", "bn", "chambana", "chicago", "decatur", "lasalle", "mattoon", "peoria", "rockford", "carbondale", "springfieldil", "quincy", "bloomington", "evansville", "fortwayne", "indianapolis", "kokomo", "tippecanoe", "muncie", "richmondin", "southbend", "terrehaute", "ames", "cedarrapids", "desmoines", "dubuque", "fortdodge", "iowacity", "masoncity", "quadcities", "siouxcity", "ottumwa", "waterloo", "lawrence", "ksu", "nwks", "salina", "seks", "swks", "topeka", "wichita", "bgky", "eastky", "lexington", "louisville", "owensboro", "westky", "batonrouge", "cenla", "houma", "lafayette", "lakecharles", "monroe", "neworleans", "shreveport", "maine", "annapolis", "baltimore", "easternshore", "frederick", "smd", "westmd", "boston", "capecod", "southcoast", "westernmass", "worcester", "annarbor", "battlecreek", "centralmich", "detroit", "flint", "grandrapids", "holland", "jxn", "kalamazoo", "lansing", "monroemi", "muskegon", "nmi", "porthuron", "saginaw", "swmi", "thumb", "up", "bemidji", "brainerd", "duluth", "mankato", "minneapolis", "rmn", "marshall", "stcloud", "gulfport", "hattiesburg", "jackson", "meridian", "northmiss", "natchez", "columbiamo", "joplin", "kansascity", "kirksville", "loz", "semo", "springfield", "stjoseph", "stlouis", "billings", "bozeman", "butte", "greatfalls", "helena", "kalispell", "missoula", "montana", "grandisland", "lincoln", "northplatte", "omaha", "scottsbluff", "elko", "lasvegas", "reno", "nh", "cnj", "jerseyshore", "newjersey", "southjersey", "albuquerque", "clovis", "farmington", "lascruces", "roswell", "santafe", "albany", "binghamton", "buffalo", "catskills", "chautauqua", "elmira", "fingerlakes", "glensfalls", "hudsonvalley", "ithaca", "longisland", "newyork", "oneonta", "plattsburgh", "potsdam", "rochester", "syracuse", "twintiers", "utica", "watertown", "asheville", "boone", "charlotte", "eastnc", "fayetteville", "greensboro", "hickory", "onslow", "outerbanks", "raleigh", "wilmington", "winstonsalem", "bismarck", "fargo", "grandforks", "nd", "akroncanton", "ashtabula", "athensohio", "chillicothe", "cincinnati", "cleveland", "columbus", "dayton", "limaohio", "mansfield", "sandusky", "toledo", "tuscarawas", "youngstown", "zanesville", "lawton", "enid", "oklahomacity", "stillwater", "tulsa", "bend", "corvallis", "eastoregon", "eugene", "klamath", "medford", "oregoncoast", "portland", "roseburg", "salem", "altoona", "chambersburg", "erie", "harrisburg", "lancaster", "allentown", "meadville", "philadelphia", "pittsburgh", "poconos", "reading", "scranton", "pennstate", "williamsport", "york", "providence", "charleston", "columbia", "florencesc", "greenville", "hiltonhead", "myrtlebeach", "nesd", "csd", "rapidcity", "siouxfalls", "sd", "chattanooga", "clarksville", "cookeville", "jacksontn", "knoxville", "memphis", "nashville", "tricities", "abilene", "amarillo", "austin", "beaumont", "brownsville", "collegestation", "corpuschristi", "dallas", "nacogdoches", "delrio", "elpaso", "galveston", "houston", "killeen", "laredo", "lubbock", "mcallen", "odessa", "sanangelo", "sanantonio", "sanmarcos", "bigbend", "texoma", "easttexas", "victoriatx", "waco", "wichitafalls", "logan", "ogden", "provo", "saltlakecity", "stgeorge", "vermont", "charlottesville", "danville", "fredericksburg", "norfolk", "harrisonburg", "lynchburg", "blacksburg", "richmond", "roanoke", "swva", "winchester", "bellingham", "kpr", "moseslake", "olympic", "pullman", "seattle", "skagit", "spokane", "wenatchee", "yakima", "charlestonwv", "martinsburg", "huntington", "morgantown", "wheeling", "parkersburg", "swv", "wv", "appleton", "eauclaire", "greenbay", "janesville", "racine", "lacrosse", "madison", "milwaukee", "northernwi", "sheboygan", "wausau", "wyoming", "micronesia", "puertorico", "virgin"}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func getRequest(rawUrl string, c *colly.Context) *colly.Request {
	h := http.Header{}

	url, _ := url.Parse(rawUrl)
	request := colly.Request{URL: url, Headers: &h, Ctx: c}
	return &request

}

func scrapeCL(stateCodes []string, w http.ResponseWriter, r *http.Request) {

	//q.AddURL()
	//q.AddURL()
	//conn, _ := upgrader.Upgrade(w, r, nil)
	//var socks5ADDRS []string
	////stooges := [...]string{}
	//for i := 1200; i < 1220; i++ {
	//	sock5URL := fmt.Sprintf("socks5://127.0.0.1:%d",i)
	//	socks5ADDRS = append(socks5ADDRS,sock5URL)
	//}
	//
	////v := map[string]string{}
	////colly.ProxyFunc()
	////proxy.RoundRobinProxySwitcher()
	//rp, err := scrape.CustomProxy(socks5ADDRS)

	//if err != nil {
	//	log.Fatal(err)
	//}
	//collector.SetProxyFunc(rp)

	//q, _ := queue.New(
	//	2, // Number of consumer threads
	//	&queue2.InMemoryQueueStorage{MaxSize:10000},
	//)

	//
	//}
	//q.Run(c)

}

func buildEmail(email string, url string) gmail.Message {
	var message gmail.Message

	temp := []byte("From: 'me'\r\n" +
		fmt.Sprintf("To: %s \r\n", email) +
		"Subject: Software Position \r\n" +
		"\r\nHey!\r\n" + " My name is Joe Jackson and I'm interested in applying for the position you posted on craigs list." +
		" This is a link to my most up to date resume https://docs.google.com/document/d/1ugz6WqXaWEj2s4CLRC5ecz40RiRUmfC9XxmvW-TSXwA/edit?usp=sharing " +
		fmt.Sprintf("%s", url) +
		"\r\nBest," + "\r\nJoe Jackson")

	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)
	return message
}

//https://sandiego.craigslist.org/nsd/sof/d/carlsbad-full-stack-web-developer/6955927244.html
//"https://sandiego.craigslist.org/contactinfo/sdo/sof/6955927244

func main() {
	c := scrape.BuildCollector()
	for _, state := range stateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
		scrape.VisitWithRetry(c, fmt.Sprintf("%s/d/software-qa-dba-etc/search/sof", stateOrg), 30)
		scrape.VisitWithRetry(c, fmt.Sprintf("%s/search/sof?employment_type=3", stateOrg), 30)
	}
	//ctx := context.Background()
	//
	//b, err := ioutil.ReadFile("/Users/joejackson/GolandProjects/scraper/cmd/craigslistAPI/credentials.json")
	//if err != nil {
	//	log.Fatalf("Unable to read client secret file: %v", err)
	//}
	//
	//config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	//if err != nil {
	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
	//}
	//client := email.GetClient(ctx, config)
	//
	//srv, err := gmail.New(client)
	//if err != nil {
	//	log.Fatalf("Unable to retrieve gmail Client %v", err)
	//}

	//gin.SetMode(gin.DebugMode)
	//r := gin.New()
	//r.Use(cors.Default())
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//
	//r.GET("/scrape", func(c *gin.Context) {
	//
	//	scrapeCL(stateCodes,c.Writer, c.Request)
	//
	//})
	//r.POST("/sendEmail", func(c *gin.Context) {
	//
	//
	//	listing := scrape.Listing{}
	//	err := json.NewDecoder(c.Request.Body).Decode(&listing)
	//	if err != nil {
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),"statusText":http.StatusText(http.StatusBadRequest),"error":err.Error()}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//
	//	contactInfo := scrape.GetContactInfoURL(listing)
	//	if contactInfo == "" {
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),"statusText":http.StatusText(http.StatusBadRequest),"error": "no contact info for link"}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//	listing.ContactInfoUrl = contactInfo
	//	r, _ := regexp.Compile(":([a-zA-Z0-9])+@job.craigslist.org")
	//	infoRESP, err := http.Get(listing.ContactInfoUrl)
	//	if infoRESP == nil {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
	//			"statusText":http.StatusText(http.StatusBadRequest),
	//			"body": string(jsonListing)}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//
	//	if err != nil {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
	//			"statusText":http.StatusText(http.StatusBadRequest),
	//			"body": string(jsonListing)}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//	htmlData, err := ioutil.ReadAll(infoRESP.Body)
	//	if err != nil {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status":fmt.Sprintf("%d",http.StatusBadRequest),
	//			"statusText":http.StatusText(http.StatusBadRequest),
	//			"body": string(jsonListing)}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//	if htmlData == nil {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
	//			"statusText":http.StatusText(http.StatusBadRequest),
	//			"body": string(jsonListing)}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//
	//	if err != nil {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
	//			"statusText":http.StatusText(http.StatusBadRequest),
	//			"body": string(jsonListing)}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//
	//	emailString := r.FindString(string(htmlData))
	//
	//	listing.Email = strings.Trim(emailString, ":")
	//
	//
	//	if listing.Email  != ""{
	//
	//		queryString := fmt.Sprintf("in:sent %s ",listing.Url)
	//
	//		messages,err := srv.Users.Messages.List("me").Q(queryString).MaxResults(10000).Do()
	//
	//		if len(messages.Messages) > 0 {
	//			jsonListing, _ := json.Marshal(listing)
	//			body := map[string]string{
	//				"status": fmt.Sprintf("%d",http.StatusAccepted),
	//				"statusText":http.StatusText(http.StatusAccepted),
	//				"body": string(jsonListing)}
	//			c.JSON(http.StatusAccepted,body)
	//			return
	//		}
	//
	//		clEmail := buildEmail(listing.Email,listing.Url)
	//
	//		emailResponse, err := srv.Users.Messages.Send("me",&clEmail).Do()
	//
	//		if err != nil {
	//			body := map[string]string{
	//				"status":fmt.Sprintf("%d",http.StatusInternalServerError),
	//				"statusText":http.StatusText(http.StatusInternalServerError),
	//				"error": err.Error()}
	//			c.JSON(http.StatusInternalServerError,body)
	//			return
	//
	//		}else {
	//			listing.EmailResponse =  emailResponse.Raw
	//			jsonListing, _ := json.Marshal(listing)
	//			body := map[string]string{"status":fmt.Sprintf("%d", http.StatusCreated),
	//				"statusText":http.StatusText(http.StatusCreated),
	//				"body": string(jsonListing)}
	//			c.JSON(http.StatusCreated,body)
	//			return
	//
	//		}
	//	}else {
	//		jsonListing, _ := json.Marshal(listing)
	//		body := map[string]string{"status": fmt.Sprintf("%d",http.StatusBadRequest),
	//			"statusText":http.StatusText(http.StatusBadRequest),
	//			"body": string(jsonListing)}
	//		c.JSON(http.StatusBadRequest,body)
	//		return
	//	}
	//})
	//r.Run()
}
