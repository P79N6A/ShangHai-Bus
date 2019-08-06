package bus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
)

const (
	QUERY_LINEID_URL = "http://apps.eshimin.com/traffic/gjc/getBusBase?name=%s"
	QUERY_DETAIL_URL = "http://apps.eshimin.com/traffic/gjc/getArriveBase?name=%s&lineid=%s&direction=%d&stopid=%d"
)

var httpClient http.Client

type LineResponse struct {
	LineName string `json:"line_name"`
	LineID   string `json:"line_id"`
}

func GetBusArrivalDetail(name string, stopID, direction int) (*Car, error) {
	id, err := getBusLineID(name)
	if err != nil {
		return nil, err
	}
	return getBusDetail(name, id, stopID, direction)
}
func getBusLineID(name string) (string, error) {
	url := fmt.Sprintf(QUERY_LINEID_URL, url2.QueryEscape(name))
	req, err := http.NewRequest("GET", url, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	lineResp := new(LineResponse)
	json.Unmarshal(body, lineResp)
	return lineResp.LineID, nil
}

type ArriveBaseResponse struct {
	Cars []Car `json:"cars"`
}

type Car struct {
	Time     string `json:"time"`
	Distance string `json:"distance"`
	Terminal string `json:"terminal"`
	Stopdis  string `json:"stopdis"`
}

func getBusDetail(name, lineID string, stopID, direction int) (*Car, error) {
	url := fmt.Sprintf(QUERY_DETAIL_URL, url2.QueryEscape(name), lineID, direction, stopID)
	req, err := http.NewRequest("GET", url, nil)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := new(ArriveBaseResponse)
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if len(result.Cars) == 0 {
		return nil, fmt.Errorf("cannot find car")
	}
	return &result.Cars[0], nil
}
