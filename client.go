package brewerydb

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

var (
	fetcher Fetcher
	base    string = "https://api.brewerydb.com/v2/"
)

type client struct {
	apiKey string
	base   string
}

type Fetcher interface {
	Get(string) ([]byte, error)
}

type HttpFetcher struct{}

func (f *HttpFetcher) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func init() {
	fetcher = &HttpFetcher{}
}

func getJson(url string) ([]byte, error) {
	return fetcher.Get(url)
}

func decodeJson(j []byte, i interface{}) {
	json.Unmarshal(j, &i)
}

func NewClient(key string) *client {
	return &client{apiKey: key, base: base}
}

func (c *client) RandomBeer(q BeerQuery) (Beer, error) {
	b := BeerResult{}
	v, _ := query.Values(q)
	url := c.base + "beer/random?withBreweries=Y&key=" + c.apiKey + "&" + v.Encode()
	j, err := getJson(url)
	decodeJson(j, b)
	if err != nil {
		return Beer{}, err
	}
	err = json.Unmarshal(j, &b)
	if err != nil {
		return Beer{}, err
	}
	return b.Beer, nil
}

func (c *client) GetBeers(q BeerQuery) ([]Beer, error) {
	b := BeerResults{}
	v, _ := query.Values(q)
	url := c.base + "beers?key=" + c.apiKey + "&" + v.Encode()

	j, err := getJson(url)
	if err != nil {
		return []Beer{}, err
	}
	err = json.Unmarshal(j, &b)
	if err != nil {
		return []Beer{}, err
	}
	return b.Beers, nil
}

func (c *client) GetLocation(q LocationQuery) ([]BreweryLocation, error) {
	l := LocationResults{}
	v, _ := query.Values(q)
	url := c.base + "locations?key=" + c.apiKey + "&" + v.Encode()
	j, err := getJson(url)
	if err != nil {
		return []BreweryLocation{}, err
	}
	err = json.Unmarshal(j, &l)
	if err != nil {
		return []BreweryLocation{}, err
	}
	return l.BreweryLocations, nil
}
