package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IamMrTeapot/LWMNAssignment/models"
)

func GetCovidData() ([]models.CovidData, error) {
	response, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	rawResponseBody, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	var responseBody models.Response

	err = json.Unmarshal(rawResponseBody, &responseBody)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return responseBody.Data, nil

}

func GetSummary() (models.Summary, error) {
	data, err := GetCovidData()

	if err != nil {
		fmt.Print(err.Error())
		return models.Summary{}, err
	}

	var summary models.Summary

	summary.Province = make(map[string]int)

	for _, covidData := range data {
		if covidData.Province == nil {
			summary.Province["N/A"]++
		} else {
			summary.Province[*covidData.Province]++
		}
		if covidData.Age == nil {
			summary.AgeGroup.NoData++
		} else if *covidData.Age <= 30 {
			summary.AgeGroup.From0To31++
		} else if *covidData.Age <= 60 {
			summary.AgeGroup.From31To60++
		} else if *covidData.Age > 60 {
			summary.AgeGroup.GTE61++
		}
	}

	return summary, nil
}
