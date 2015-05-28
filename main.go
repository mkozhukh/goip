//Package goip provides API for GeoIP2 Precision web services
//http://maxmind.com
package goip

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//GoIP is a core object of this package
//You can use Connect function to get the instance
type GoIP struct {
	HttpAuth string
}

//ResponseNames contains names of object in different languages
type ResponseNames struct {
	De string
	En string
	Es string
	Fr string
	Ja string
	Ru string
	Pt string `json:"pt-BR"`
	Zh string `json:"zh-CN"`
}

//ResponseContinent contains information about a single continent
type ResponseContinent struct {
	Names     ResponseNames
	Code      string
	GeonameId int `json:"geoname_id"`
}

//ResponseCountry contains information about a single country
type ResponseCountry struct {
	Confidence int    `json:"confidence"`
	IsoCode    string `json:"iso_code"`
	GeonameId  int    `json:"geoname_id"`
	Names      ResponseNames
}

//ResponseInfo contains technical information about service
type ResponseInfo struct {
	QueriesRemaining int `json:"queries_remaining"`
}

//Response contains all info from the GeoIp service
type Response struct {
	Error             string
	Country           ResponseCountry
	RegisteredCountry ResponseCountry
	Continent         ResponseContinent
	MaxMind           ResponseInfo
}

//Connect takes name and license of GeoIp user
//It returns GoIP object that can be used for further information retrieving
func Connect(user string, pass string) *GoIP {
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	return &GoIP{HttpAuth: auth}
}

//countryURL points to the URL of GeoIP service
const countryURL = "https://geoip.maxmind.com/geoip/v2.1/country/"

//countryInfo method sends request to the GeoIP service and
//returns all info about the IP address
func (m *GoIP) countryInfo(ip string) (Response, error) {
	req, err := http.NewRequest("GET", countryURL+ip, nil)
	response := Response{}

	if err == nil {
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept-Charset", "UTF-8")
		req.Header.Set("Authorization", "Basic "+m.HttpAuth)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err == nil {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			if err == nil {
				err = json.Unmarshal(body, &response)
			}
		}
	}

	if response.Error != "" {
		err = errors.New(response.Error)
	}

	return response, err

}

//Country method retrieves info about IP address
//and returns the Country structure with information about the related Country
func (m *GoIP) Country(ip string) (ResponseCountry, error) {
	info, err := m.countryInfo(ip)
	return info.Country, err
}

//Continent method retrieves info about IP address
//and returns the Continent structure with information about the related continent
func (m *GoIP) Continent(ip string) (ResponseContinent, error) {
	info, err := m.countryInfo(ip)
	return info.Continent, err
}

//ContinentName method returns English name of continent by IP
func (m *GoIP) ContinentName(ip string) (string, error) {
	country, err := m.Continent(ip)
	if err == nil {
		return country.Names.En, nil
	}

	return "", err
}

//ContinentName method returns English name of country by IP
func (m *GoIP) CountryName(ip string) (string, error) {
	country, err := m.Country(ip)
	if err == nil {
		return country.Names.En, nil
	}

	return "", err
}
