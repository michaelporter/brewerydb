package brewerydb

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestFetcher struct {
	retval string
	url    string
}

func (t *TestFetcher) Get(url string) ([]byte, error) {
	t.url = url
	var json string
	if t.retval != "" {
		json = t.retval
	} else {
		json = `
		{"beer": "none"}
	`

	}
	return []byte(json), nil
}

type ErrorFetcher struct{}

func (t *ErrorFetcher) Get(url string) ([]byte, error) {
	return []byte{}, errors.New("Oh noes")
}

func TestGetRandomBeer(t *testing.T) {
	Convey("Given we fetch a random beer", t, func() {
		Convey("We should see a beer and no errors", func() {
			tf := &TestFetcher{}
			tf.retval = `{"status": "ok", "data": {"description": "some_beer"}}`
			fetcher = tf
			c := NewClient("123")
			q := BeerQuery{Abv: "6"}
			beer, err := c.RandomBeer(q)
			So(err, ShouldEqual, nil)
			So(beer.Description, ShouldEqual, "some_beer")
		})

	})
	Convey("Given we fetch a random beer and get bad JSON back", t, func() {
		Convey("We should see a an error and no beer", func() {
			tf := &TestFetcher{}
			tf.retval = `}THIS_AINT_JSON{`
			fetcher = tf
			c := NewClient("123")
			q := BeerQuery{Abv: "6"}
			beer, err := c.RandomBeer(q)
			So(err, ShouldNotBeNil)
			So(beer, ShouldResemble, Beer{})

		})
	})
	Convey("Given we fetch a random beer and the client encounters an error", t, func() {
		Convey("We should see an error", func() {
			tf := &ErrorFetcher{}
			fetcher = tf
			c := NewClient("123")
			q := BeerQuery{Abv: "6"}
			beer, err := c.RandomBeer(q)
			So(err, ShouldNotBeNil)
			So(beer, ShouldResemble, Beer{})

		})
	})

}

func TestGetBeers(t *testing.T) {
	Convey("Given we fetch a list of beers", t, func() {
		Convey("We should get a list of beers and no errors", func() {
			tf := &TestFetcher{}
			tf.retval = `{"status": "ok", "data": [{"description": "some_beer1"}, {"description": "some_beer2"}]}`
			fetcher = tf
			c := NewClient("123")
			q := BeerQuery{Abv: "6"}
			beers, err := c.GetBeers(q)
			So(err, ShouldEqual, nil)
			So(len(beers), ShouldEqual, 2)
			So(beers[0].Description, ShouldEqual, "some_beer1")
			So(beers[1].Description, ShouldEqual, "some_beer2")
		})

	})
	Convey("Given we fetch a list of beers and get bad JSON", t, func() {
		Convey("We should see a an error and no beers", func() {
			tf := &TestFetcher{}
			tf.retval = `}THIS_AINT_JSON{`
			fetcher = tf
			c := NewClient("123")
			q := BeerQuery{Abv: "6"}
			beers, err := c.GetBeers(q)
			So(err, ShouldNotBeNil)
			So(beers, ShouldResemble, []Beer{})

		})
	})
	Convey("Given we fetch a list of beers and the client encounters and error", t, func() {
		Convey("We should see an error", func() {
			tf := &ErrorFetcher{}
			fetcher = tf
			c := NewClient("123")
			q := BeerQuery{Abv: "6"}
			beers, err := c.GetBeers(q)
			So(err, ShouldNotBeNil)
			So(beers, ShouldResemble, []Beer{})

		})
	})

}

func TestGetLocation(t *testing.T) {
	Convey("Given we fetch a location", t, func() {
		Convey("We should get a list of brewery locations and no error", func() {
			tf := &TestFetcher{}
			tf.retval = `{"status": "ok", "data": [{"breweryId": "123"}, {"breweryId": "456"}]}`
			fetcher = tf
			c := NewClient("123")
			q := LocationQuery{PostalCode: "10101"}
			locations, err := c.GetLocation(q)
			So(err, ShouldEqual, nil)
			So(len(locations), ShouldEqual, 2)
		})

	})
	Convey("Given we fetch a location and get bad JSON", t, func() {
		Convey("We should see a an error and no locations", func() {
			tf := &TestFetcher{}
			tf.retval = `}THIS_AINT_JSON{`
			fetcher = tf
			c := NewClient("123")
			q := LocationQuery{PostalCode: "10101"}
			locations, err := c.GetLocation(q)
			So(err, ShouldNotBeNil)
			So(locations, ShouldResemble, []BreweryLocation{})

		})
	})
	Convey("Given we fetch a locations and the client sees and error", t, func() {
		Convey("We should see an error", func() {
			tf := &ErrorFetcher{}
			fetcher = tf
			c := NewClient("123")
			q := LocationQuery{PostalCode: "10101"}
			locations, err := c.GetLocation(q)
			So(err, ShouldNotBeNil)
			So(locations, ShouldResemble, []BreweryLocation{})

		})
	})

}
