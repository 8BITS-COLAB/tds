# TDS API

## Install

Run `go get -u github.com/ElioenaiFerrari/tds`

## Examples

- PreAuth

  ```go
    acctNumber := "9001100911111111"

    res, err := preauth.Call(&preauth.Request{
      AccountNumber: acctNumber,
    })

    if err != nil {
      log.Fatal(err)
    }
  ```

- Auth

  ```go
    aRes, err := auth.Call(&auth.Request{
      ServerTransactionID:        res.ServerTransactionID,
      CardNumber:                 acctNumber,
      DeviceChannel:              "02",
      MessageCategory:            "02",
      MessageType:                "AReq",
      MessageVersion:             "2.2.0",
      CompleteIndicator:          "Y",
      BrowserAcceptHeader:        "Accept,Content-Type",
      BrowserJavascriptEnabled:   false,
      BrowserLanguage:            "pt-BR",
      BrowserUserAgent:           "Firefox",
      NotificationURL:            "http://localhost:4000/api/v1/callback",
      RequestorAuthenticationInd: "01",
      RequestorURL:               "http://localhost:4000/api/v1/callback",
    })

    if err != nil {
      log.Fatal(err)
    }
  ```

- PostAuth

  ```go
    rRes, err := postauth.Call(&postauth.Request{
      ServerTransactionID: aRes.ServerTransactionID,
    })

    if err != nil {
      log.Fatal(err)
    }
  ```
