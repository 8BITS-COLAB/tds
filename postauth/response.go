package postauth

type Response struct {
	ACSEndProtocolVersion   string `json:"acsEndProtocolVersion"`
	ACSStartProtocolVersion string `json:"acsStartProtocolVersion"`
	DSEndProtocolVersion    string `json:"dsEndProtocolVersion"`
	DSStartProtocolVersion  string `json:"dsStartProtocolVersion"`
	MessageType             string `json:"messageType"`
	Scheme                  string `json:"scheme"`
	MethodURL               string `json:"threeDSMethodURL"`
	ServerTransactionID     string `json:"threeDSServerTransID"`
}
