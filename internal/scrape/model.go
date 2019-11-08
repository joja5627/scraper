package scrape


//Listing comment
type Listing struct {
	//ID             bson.ObjectId "_id,omitempty"
	//ScrapeID       bson.ObjectId "_id,omitempty"
	//Date           time.Time     `json:"created"`
	StateCode      string        `json:"stateOrg"`
	Title      	   string        `json:"title"`
	Url            string        `json:"url"`
	ContactInfoUrl string        `json:"contactInfoUrl"`
	ListingInfoResponse string    `json:"listingInfoResponse"`
	EmailResponse  string        `json:"emailResponse"`
	Email          string        `json:"email"`
}
//SocketMessage
type SocketMessage struct{
	MessageType string `json:"messageType"`
	Payload string `json:"payload"`
}