package main

type createSecretPayload struct {
	PlainText string `json:"plain_text"`
}

type createSecretResponse struct {
	Id string `json:"id"`
}

type secretData struct {
	Id     string
	Secret string
}
