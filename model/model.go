package model

type Data struct {
	Data []RawData `json:"Data"`
}

type RawData struct {
	Age      int    `json:"Age"`
	NationEn string `json:"NationEn"`
	Province string `json:"Province"`
}
