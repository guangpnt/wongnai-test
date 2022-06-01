package services

import (
	"crypto/tls"
	"encoding/json"
	"example/wongnai-test/model"
	"io/ioutil"
	"log"
	"net/http"
)

func GetData() model.Data {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	rawData := model.Data{}

	resp, _ := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")

	body, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(body, &rawData)
	if err != nil {
		log.Fatal(err)

	}

	return rawData
}
