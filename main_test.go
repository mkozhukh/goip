package goip

import "testing"

//you need to set the username and license to run the tests
var userName string = "------"
var license string = "----------"

func TestConnect(t *testing.T) {
	Connect(userName, license)
}

func TestCountryName(t *testing.T) {
	check := Connect(userName, license)
	country, err := check.CountryName("80.249.82.222")
	if err != nil {
		t.Error("Can't resolve IP to Country\n", err)
	} else if country != "Belarus" {
		t.Error("Incorrect IP to Country (expected Belarus)", country)
	}
}

func TestContinentName(t *testing.T) {
	check := Connect(userName, license)
	continent, err := check.ContinentName("80.249.82.222")
	if err != nil {
		t.Error("Can't resolve IP to Continent\n", err)
	} else if continent != "Europe" {
		t.Error("Incorrect IP to Continent (expected Europe)", continent)
	}
}
