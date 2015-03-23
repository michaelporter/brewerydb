package brewerydb

type BeerQuery struct {
	Page             int    `url:"p,omitempty"`
	Ids              string `url:"ids,omitempty"`
	Name             string `url:"name,omitempty"`
	Abv              string `url:"abv,omitempty"`
	Ibu              string `url:"ibu,omitempty"`
	IncludeBreweries string `url:"withBreweries,omitempty"`
}

type LocationQuery struct {
	Page       int    `url:"p,omitempty"`
	Ids        string `url:"ids,omitempty"`
	Locality   string `url:"locality,omitempty"`
	Region     string `url:"region,omitempty"`
	PostalCode string `url:"postalCode,omitempty"`
	Closed     bool   `url:"isClosed,omitempty"`
}
