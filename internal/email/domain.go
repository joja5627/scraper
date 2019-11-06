package email

type GmailCreds struct {
	ServerAddr string
	Password   string
	EmailAddr  string
	PortNumber int
	UserName   string
}
type Template struct {
	To             []string
	From           string
	AttachmentPath string
}
