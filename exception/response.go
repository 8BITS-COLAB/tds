package exception

type Response struct {
	MessageType      string `json:"messageType"`
	MessageVersion   string `json:"messageVersion"`
	ErrorCode        string `json:"errorCode"`
	ErrorComponent   string `json:"errorComponent"`
	ErrorDescription string `json:"errorDescription"`
	ErrorDetail      string `json:"errorDetail"`
	ErrorMessageType string `json:"errorMessageType"`
}
