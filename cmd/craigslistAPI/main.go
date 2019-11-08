package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joja5627/scraper/internal/email"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"regexp"
	"strings"

	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
	"github.com/joja5627/scraper/internal/scrape"
	"github.com/joja5627/scraper/internal/utils"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	stateCodes = []string{"auburn", "bham", "dothan", "shoals", "gadsden", "huntsville", "mobile", "montgomery", "tuscaloosa", "anchorage", "fairbanks", "kenai", "juneau", "flagstaff", "mohave", "phoenix", "prescott", "showlow", "sierravista", "tucson", "yuma", "fayar", "fortsmith", "jonesboro", "littlerock", "texarkana", "bakersfield", "chico", "fresno", "goldcountry", "hanford", "humboldt", "imperial", "inlandempire", "losangeles", "mendocino", "merced", "modesto", "monterey", "orangecounty", "palmsprings", "redding", "sacramento", "sandiego", "sfbay", "slo", "santabarbara", "santamaria", "siskiyou", "stockton", "susanville", "ventura", "visalia", "yubasutter", "boulder", "cosprings", "denver", "eastco", "fortcollins", "rockies", "pueblo", "westslope", "newlondon", "hartford", "newhaven", "nwct", "delaware", "washingtondc", "miami", "daytona", "keys", "fortlauderdale", "fortmyers", "gainesville", "cfl", "jacksonville", "lakeland", "miami", "lakecity", "ocala", "okaloosa", "orlando", "panamacity", "pensacola", "sarasota", "miami", "spacecoast", "staugustine", "tallahassee", "tampa", "treasure", "miami", "albanyga", "athensga", "atlanta", "augusta", "brunswick", "columbusga", "macon", "nwga", "savannah", "statesboro", "valdosta", "honolulu", "boise", "eastidaho", "lewiston", "twinfalls", "bn", "chambana", "chicago", "decatur", "lasalle", "mattoon", "peoria", "rockford", "carbondale", "springfieldil", "quincy", "bloomington", "evansville", "fortwayne", "indianapolis", "kokomo", "tippecanoe", "muncie", "richmondin", "southbend", "terrehaute", "ames", "cedarrapids", "desmoines", "dubuque", "fortdodge", "iowacity", "masoncity", "quadcities", "siouxcity", "ottumwa", "waterloo", "lawrence", "ksu", "nwks", "salina", "seks", "swks", "topeka", "wichita", "bgky", "eastky", "lexington", "louisville", "owensboro", "westky", "batonrouge", "cenla", "houma", "lafayette", "lakecharles", "monroe", "neworleans", "shreveport", "maine", "annapolis", "baltimore", "easternshore", "frederick", "smd", "westmd", "boston", "capecod", "southcoast", "westernmass", "worcester", "annarbor", "battlecreek", "centralmich", "detroit", "flint", "grandrapids", "holland", "jxn", "kalamazoo", "lansing", "monroemi", "muskegon", "nmi", "porthuron", "saginaw", "swmi", "thumb", "up", "bemidji", "brainerd", "duluth", "mankato", "minneapolis", "rmn", "marshall", "stcloud", "gulfport", "hattiesburg", "jackson", "meridian", "northmiss", "natchez", "columbiamo", "joplin", "kansascity", "kirksville", "loz", "semo", "springfield", "stjoseph", "stlouis", "billings", "bozeman", "butte", "greatfalls", "helena", "kalispell", "missoula", "montana", "grandisland", "lincoln", "northplatte", "omaha", "scottsbluff", "elko", "lasvegas", "reno", "nh", "cnj", "jerseyshore", "newjersey", "southjersey", "albuquerque", "clovis", "farmington", "lascruces", "roswell", "santafe", "albany", "binghamton", "buffalo", "catskills", "chautauqua", "elmira", "fingerlakes", "glensfalls", "hudsonvalley", "ithaca", "longisland", "newyork", "oneonta", "plattsburgh", "potsdam", "rochester", "syracuse", "twintiers", "utica", "watertown", "asheville", "boone", "charlotte", "eastnc", "fayetteville", "greensboro", "hickory", "onslow", "outerbanks", "raleigh", "wilmington", "winstonsalem", "bismarck", "fargo", "grandforks", "nd", "akroncanton", "ashtabula", "athensohio", "chillicothe", "cincinnati", "cleveland", "columbus", "dayton", "limaohio", "mansfield", "sandusky", "toledo", "tuscarawas", "youngstown", "zanesville", "lawton", "enid", "oklahomacity", "stillwater", "tulsa", "bend", "corvallis", "eastoregon", "eugene", "klamath", "medford", "oregoncoast", "portland", "roseburg", "salem", "altoona", "chambersburg", "erie", "harrisburg", "lancaster", "allentown", "meadville", "philadelphia", "pittsburgh", "poconos", "reading", "scranton", "pennstate", "williamsport", "york", "providence", "charleston", "columbia", "florencesc", "greenville", "hiltonhead", "myrtlebeach", "nesd", "csd", "rapidcity", "siouxfalls", "sd", "chattanooga", "clarksville", "cookeville", "jacksontn", "knoxville", "memphis", "nashville", "tricities", "abilene", "amarillo", "austin", "beaumont", "brownsville", "collegestation", "corpuschristi", "dallas", "nacogdoches", "delrio", "elpaso", "galveston", "houston", "killeen", "laredo", "lubbock", "mcallen", "odessa", "sanangelo", "sanantonio", "sanmarcos", "bigbend", "texoma", "easttexas", "victoriatx", "waco", "wichitafalls", "logan", "ogden", "provo", "saltlakecity", "stgeorge", "vermont", "charlottesville", "danville", "fredericksburg", "norfolk", "harrisonburg", "lynchburg", "blacksburg", "richmond", "roanoke", "swva", "winchester", "bellingham", "kpr", "moseslake", "olympic", "pullman", "seattle", "skagit", "spokane", "wenatchee", "yakima", "charlestonwv", "martinsburg", "huntington", "morgantown", "wheeling", "parkersburg", "swv", "wv", "appleton", "eauclaire", "greenbay", "janesville", "racine", "lacrosse", "madison", "milwaukee", "northernwi", "sheboygan", "wausau", "wyoming", "micronesia", "puertorico", "virgin"}
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}



func scrapeCL(statesQueue utils.Queue,w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &scrape.Client{Conn: conn, TerminateChannel: make(chan bool)}

	go client.WriteSocket(statesQueue)
	go client.ReadSocket()


}
// 		"To: joja5627@gmail.com \r\n" +

func buildEmail(email string) gmail.Message{
	var message gmail.Message

	temp := []byte("From: 'me'\r\n" +
		fmt.Sprintf( "To: %s \r\n", email)+
		"Subject: Software Position \r\n" +
		"\r\nHey! My name is Joe Jackson and I'm interested in applying for the position you posted on craigs list." +
		" This is a link to my most up to date resume https://docs.google.com/document/d/1ugz6WqXaWEj2s4CLRC5ecz40RiRUmfC9XxmvW-TSXwA/edit?usp=sharing "+
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
	ctx := context.Background()

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := email.GetClient(ctx, config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}
	stateQueue := utils.New()

	for _, state := range stateCodes {
		stateQueue.Add(state)
	}

	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/scrape", func(c *gin.Context) {

		scrapeCL(stateQueue,c.Writer, c.Request)

	})
	r.POST("/sendEmail", func(c *gin.Context) {

		listing := scrape.Listing{}
		err := json.NewDecoder(c.Request.Body).Decode(&listing)
		if err != nil {
			body := map[string]string{"error": "bad request body"}
			c.JSON(500,body)
		}
		listing.ContactInfoUrl = scrape.GetContactInfoURL(listing)
		r, _ := regexp.Compile(":([a-zA-Z0-9])+@job.craigslist.org")
		infoRESP, err := http.Get(listing.ContactInfoUrl)
		htmlData, err := ioutil.ReadAll(infoRESP.Body)
		listing.ListingInfoResponse = string(htmlData)
		emailString := r.FindString(listing.ListingInfoResponse)
		listing.Email = strings.Trim(emailString, ":")


		if listing.Email  != ""{
			clEmail := buildEmail(listing.Email )
			emailResponse, err := srv.Users.Messages.Send("me",&clEmail).Do()

			if err != nil {
				body := map[string]string{"error": err.Error()}
				utils.RespondJSON(c.Writer,http.StatusInternalServerError,body)

			}else {
				listing.EmailResponse =  emailResponse.Raw
				jsonListing, _ := json.Marshal(listing)
				body := map[string]string{"body":string(jsonListing) }
				utils.RespondJSON(c.Writer,http.StatusOK,body)
			}
		}else {
			jsonListing, _ := json.Marshal(listing)
			body := map[string]string{"body":string(jsonListing) }
			utils.RespondJSON(c.Writer,http.StatusBadRequest,body)
		}
	})
	r.Run()
}
