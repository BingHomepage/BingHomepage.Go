package BingHomepage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Get returns BingHomepage structure including data for URL, Copyright, and Copyright link
func Get(CountryCode string) (*BingHomepage, error) {
	api := "http://cdn.muzzammil.xyz/bing/bing.php?format=json&cc=" + CountryCode
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	binghomepage := &BingHomepage{}
	if json.Unmarshal(data, binghomepage) != nil {
		return nil, err
	}

	return binghomepage, nil
}

//BingHomepage structure -- includes URL, Copyright, and Copyright
type BingHomepage struct {
	URL           string `json:"url"`
	Copyright     string `json:"copyright"`
	CopyrightLink string `json:"copyrightlink"`
}
