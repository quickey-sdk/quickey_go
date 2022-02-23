package quickey

const (
	APIVersion = "v0.1.1"
	APIURL     = "https://api.getquickey.com"
)

type Response struct {
	ApiKey   string
	BaseUrl  string
	App      []App
	Auth     []Auth
	Customer []Customer
}

type App struct {
	Email          string `json:"email"`
	AppName        string `json:"appName"`
	SocialApps     string `json:"socialApps"`
	RedirectUri    string `json:"redirectUri"`
	RedirectUrlApp string `json:"redirectUrlApp"`
	ApiKey         string `json:"apiKey"`
}

type Auth struct {
	Token string `json:"access_token"`
	User  string `json:"user"`
}

type Customer struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	OTP     string `json:"otp"`
	Expires string `json:"expires"`
}


type VerifiedData struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Id     string `json:"_id"`
	JWTExpires string `json:"iat"`
}

func New(api_key string) *Response {
	return &Response{
		ApiKey:  api_key,
		BaseUrl: APIURL,
	}
}
