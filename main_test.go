package goip

import "testing"

//you need to set the username and license to run the tests
var userName = "101479"
var license = "LLrtcBQcPkkT"

func TestConnect(t *testing.T) {
	New(userName, license)
}

func TestCountryName(t *testing.T) {
	check := New(userName, license)
	country, err := check.CountryName("80.249.82.222")
	if err != nil {
		t.Error("Can't resolve IP to Country\n", err)
	} else if country != "Belarus" {
		t.Error("Incorrect IP to Country (expected Belarus)", country)
	}
}

func TestCountryNameFail(t *testing.T) {
	check := New(userName, license)
	_, err := check.CountryName("not-exists")
	if err == nil {
		t.Error("Error not produced for the invalid input")
	}
}

func TestContinentName(t *testing.T) {
	check := New(userName, license)
	continent, err := check.ContinentName("80.249.82.222")
	if err != nil {
		t.Error("Can't resolve IP to Continent\n", err)
	} else if continent != "Europe" {
		t.Error("Incorrect IP to Continent (expected Europe)", continent)
	}
}

func TestContinentNameFail(t *testing.T) {
	check := New(userName, license)
	_, err := check.ContinentName("not-exists")
	if err == nil {
		t.Error("Error not produced for the invalid input")
	}
}
