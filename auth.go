package quickey

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (q *Response) GetAccessTokenByEmail(email string, provider string) *Auth {
	values := map[string]string{"email": email, "provider": provider}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", q.BaseUrl+"/loginCustomer", bytes.NewBuffer(json_data))
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

	auth := Auth{}
	json.Unmarshal([]byte(responseString), &auth)
	q.Auth = append(q.Auth, auth)

	return &auth
}

func (q *Response) GetAccessTokenByPhone(phone string, provider string) *Auth {
	values := map[string]string{"phone": phone, "provider": provider}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", q.BaseUrl+"/loginCustomer", bytes.NewBuffer(json_data))
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

	auth := Auth{}
	json.Unmarshal([]byte(responseString), &auth)
	q.Auth = append(q.Auth, auth)

	return &auth
}

func (q *Response) linkPhoneToEmail(phone string, token string) *Customer {
	values := map[string]string{"phone": phone, "token": token}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", q.BaseUrl+"/otp/linkToEmail", bytes.NewBuffer(json_data))
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
