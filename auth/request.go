package auth

type Request struct {
	CardNumber                 string `json:"acctNumber"`
	BrowserAcceptHeader        string `json:"browserAcceptHeader"`
	BrowserAcceptDepth         string `json:"browserAcceptDepth"`
	BrowserJavaEnabled         bool   `json:"browserJavaEnabled"`
	BrowserJavascriptEnabled   bool   `json:"browserJavascriptEnabled"`
	BrowserLanguage            string `json:"browserLanguage"`
	BrowserScreenHeight        string `json:"browserScreenHeight"`
	BrowserScreenWidth         string `json:"browserScreenWidth"`
	BrowserTimeZone            string `json:"browserTZ"`
	BrowserUserAgent           string `json:"browserUserAgent"`
	ServerTransactionID        string `json:"threeDSServerTransID"`
	CompleteIndicator          string `json:"threeDSCompInd"`
	DeviceChannel              string `json:"deviceChannel,omitempty"`
	MessageCategory            string `json:"messageCategory,omitempty"`
	MessageType                string `json:"messageType,omitempty"`
	MessageVersion             string `json:"messageVersion,omitempty"`
	NotificationURL            string `json:"notificationUrl,omitempty"`
	RequestorAuthenticationInd string `json:"threeDSRequestorAuthenticationInd,omitempty"`
	RequestorURL               string `json:"threeDSRequestorURL,omitempty"`
}
