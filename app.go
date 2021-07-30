package quickey

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (q *Response) GetMetadata() *App {
	values := map[string]string{"apiKey": q.ApiKey}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := http.Post(q.BaseUrl+"/auth/apiKey", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	var responseMap map[string]interface{}
	json.NewDecoder(responseJSON.Body).Decode(&responseMap)

	responseBytes, err := json.Marshal(responseMap["app"])

	responseString := string(responseBytes)

	app := App{}
	json.Unmarshal([]byte(responseString), &app)
	q.App = append(q.App, app)

	return &app
}

func (q *Response) SendSMSOTP(phone string, provider string, appId string) *Customer {
	values := map[string]string{"phone": phone, "provider": provider}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", q.BaseUrl+"/getSMSOTP/"+appId, bytes.NewBuffer(json_data))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("authorization", q.ApiKey)

	client := &http.Client{}
	responseJSON, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer responseJSON.Body.Close()

	var responseMap map[string]interface{}
	json.NewDecoder(responseJSON.Body).Decode(&responseMap)

	responseBytes, err := json.Marshal(responseMap)

	responseString := string(responseBytes)

	customer := Customer{}
	json.Unmarshal([]byte(responseString), &customer)
	q.Customer = append(q.Customer, customer)

	return &customer
}
