# QuickeyRubySdk

A Passwordless Authentication Platform and Login Management System for Application

## Installation

    $ go get github.com/quickey-sdk/quickey_go

## Usage

```
import quickey 'github.com/quickey-sdk/quickey_go'

q := quickey.New("YOUR API KEY")
```
### Get App Data
```
appData := q.GetMetadata()
fmt.Println(appData)
```
### Get Access Token
```
accessToken := q.GetAccessToken(appData.Email).Token

or 

accessToken := q.GetAccessToken("YOUR APP EMAIL").Token
fmt.Println(accessToken)
```

