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

### Get Access Token By Email
```
accessToken := q.GetAccessTokenByEmail("YOUR EMAIL", "YOUR PROVIDER").Token
fmt.Println(accessToken)
```

### Get Access Token By Phone
```
accessToken := q.GetAccessTokenByPhone("YOUR PHONE", "YOUR PROVIDER").Token
fmt.Println(accessToken)
```

### Send SMS OTP
```
response := q.SendSMSOTP("YOUR PHONE", "YOUR PROVIDER")
fmt.Println(response)
```

### Link Phone To Email
```
response := q.LinkPhoneToEmail("YOUR PHONE", "YOUR TOKEN")
fmt.Println(response)
```
