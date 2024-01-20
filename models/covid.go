package models

type CovidData struct {
	ConfirmDate    *string `json:"ConfirmDate"`
	No             *int    `json:"No"`
	Age            *int    `json:"Age"`
	Gender         *string `json:"Gender"`
	GenderEn       *string `json:"GenderEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       *string `json:"Province"`
	ProvinceId     *int    `json:"ProvinceId"`
	District       *string `json:"District"`
	ProvinceEn     *string `json:"ProvinceEn"`
	StatQuarantine *int    `json:"StatQuarantine"`
}

type Response struct {
	Data []CovidData `json:"Data"`
}

type Summary struct {
	Province map[string]int `json:"Province"`
	AgeGroup struct {
		From0To31  int `json:"0-30"`
		From31To60 int `json:"31-60"`
		GTE61      int `json:"61+"`
		NoData     int `json:"N/A"`
	} `json:"AgeGroup"`
}
