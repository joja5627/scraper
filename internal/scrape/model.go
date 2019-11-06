package scrape


//Listing comment
type Listing struct {
	//ID             bson.ObjectId "_id,omitempty"
	//ScrapeID       bson.ObjectId "_id,omitempty"
	//Date           time.Time     `json:"created"`
	StateCode      string        `json:"stateOrg"`
	Url            string        `json:"listingUrl"`
	ContactInfoUrl string        `json:"contactInfoUrl"`
	EmailResponse  string        `json:"emailResponse"`
	Email          string        `json:"email"`
}
//SocketMessage
type SocketMessage struct{
	MessageType string `json:"messageType"`
	Payload string `json:"payload"`
}