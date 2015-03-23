package brewerydb

type Beer struct {
	Abv       string `json:"abv"`
	Available struct {
		Description string `json:"description"`
		Id          int64  `json:"id"`
		Name        string `json:"name"`
	} `json:"available"`
	AvailableId int64 `json:"availableId"`
	Breweries   []struct {
		CreateDate  string `json:"createDate"`
		Description string `json:"description"`
		Established string `json:"established"`
		Id          string `json:"id"`
		Images      struct {
			Icon   string `json:"icon"`
			Large  string `json:"large"`
			Medium string `json:"medium"`
		} `json:"images"`
		IsOrganic string `json:"isOrganic"`
		Locations []struct {
			Country struct {
				CreateDate  string `json:"createDate"`
				DisplayName string `json:"displayName"`
				IsoCode     string `json:"isoCode"`
				IsoThree    string `json:"isoThree"`
				Name        string `json:"name"`
				NumberCode  int64  `json:"numberCode"`
			} `json:"country"`
			CountryIsoCode      string  `json:"countryIsoCode"`
			CreateDate          string  `json:"createDate"`
			HoursOfOperation    string  `json:"hoursOfOperation"`
			Id                  string  `json:"id"`
			InPlanning          string  `json:"inPlanning"`
			IsClosed            string  `json:"isClosed"`
			IsPrimary           string  `json:"isPrimary"`
			Latitude            float64 `json:"latitude"`
			Locality            string  `json:"locality"`
			LocationType        string  `json:"locationType"`
			LocationTypeDisplay string  `json:"locationTypeDisplay"`
			Longitude           float64 `json:"longitude"`
			Name                string  `json:"name"`
			OpenToPublic        string  `json:"openToPublic"`
			Phone               string  `json:"phone"`
			PostalCode          string  `json:"postalCode"`
			Region              string  `json:"region"`
			Status              string  `json:"status"`
			StatusDisplay       string  `json:"statusDisplay"`
			StreetAddress       string  `json:"streetAddress"`
			UpdateDate          string  `json:"updateDate"`
			Website             string  `json:"website"`
		} `json:"locations"`
		Name          string `json:"name"`
		Status        string `json:"status"`
		StatusDisplay string `json:"statusDisplay"`
		UpdateDate    string `json:"updateDate"`
		Website       string `json:"website"`
	} `json:"breweries"`
	CreateDate   string `json:"createDate"`
	Description  string `json:"description"`
	FoodPairings string `json:"foodPairings"`
	Glass        struct {
		CreateDate string `json:"createDate"`
		Id         int64  `json:"id"`
		Name       string `json:"name"`
	} `json:"glass"`
	GlasswareId int64  `json:"glasswareId"`
	Ibu         string `json:"ibu"`
	Id          string `json:"id"`
	IsOrganic   string `json:"isOrganic"`
	Labels      struct {
		Icon   string `json:"icon"`
		Large  string `json:"large"`
		Medium string `json:"medium"`
	} `json:"labels"`
	Name                      string `json:"name"`
	OriginalGravity           string `json:"originalGravity"`
	ServingTemperature        string `json:"servingTemperature"`
	ServingTemperatureDisplay string `json:"servingTemperatureDisplay"`
	Srm                       struct {
		Hex  string `json:"hex"`
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"srm"`
	SrmId         int64  `json:"srmId"`
	Status        string `json:"status"`
	StatusDisplay string `json:"statusDisplay"`
	Style         Style  `json:"style"`
	StyleId       int64  `json:"styleId"`
	UpdateDate    string `json:"updateDate"`
	Year          int64  `json:"year"`
}

type Style struct {
	AbvMax   string `json:"abvMax"`
	AbvMin   string `json:"abvMin"`
	Category struct {
		CreateDate string `json:"createDate"`
		Id         int64  `json:"id"`
		Name       string `json:"name"`
	} `json:"category"`
	CategoryId  int64  `json:"categoryId"`
	CreateDate  string `json:"createDate"`
	Description string `json:"description"`
	FgMax       string `json:"fgMax"`
	FgMin       string `json:"fgMin"`
	IbuMax      string `json:"ibuMax"`
	IbuMin      string `json:"ibuMin"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	OgMin       string `json:"ogMin"`
	SrmMax      string `json:"srmMax"`
	SrmMin      string `json:"srmMin"`
}

type BeerResults struct {
	CurrentPage   int64  `json:"currentPage"`
	Beers         []Beer `json:"data"`
	NumberOfPages int64  `json:"numberOfPages"`
	Status        string `json:"status"`
	TotalResults  int64  `json:"totalResults"`
}

type BeerResult struct {
	Status string `json:"status"`
	Beer   Beer   `json:"data"`
}

type LocationResults struct {
	CurrentPage      int64             `json:"currentPage"`
	BreweryLocations []BreweryLocation `json:"data"`
	NumberOfPages    int64             `json:"numberOfPages"`
	Status           string            `json:"status"`
	TotalResults     int64             `json:"totalResults"`
}

type BreweryLocation struct {
	Brewery   Brewery
	BreweryId string `json:"breweryId"`
	Country   struct {
		CreateDate  string `json:"createDate"`
		DisplayName string `json:"displayName"`
		IsoCode     string `json:"isoCode"`
		IsoThree    string `json:"isoThree"`
		Name        string `json:"name"`
		NumberCode  int64  `json:"numberCode"`
	} `json:"country"`
	CountryIsoCode      string  `json:"countryIsoCode"`
	CreateDate          string  `json:"createDate"`
	ExtendedAddress     string  `json:"extendedAddress"`
	HoursOfOperation    string  `json:"hoursOfOperation"`
	Id                  string  `json:"id"`
	InPlanning          string  `json:"inPlanning"`
	IsClosed            string  `json:"isClosed"`
	IsPrimary           string  `json:"isPrimary"`
	Latitude            float64 `json:"latitude"`
	Locality            string  `json:"locality"`
	LocationType        string  `json:"locationType"`
	LocationTypeDisplay string  `json:"locationTypeDisplay"`
	Longitude           float64 `json:"longitude"`
	Name                string  `json:"name"`
	OpenToPublic        string  `json:"openToPublic"`
	Phone               string  `json:"phone"`
	PostalCode          string  `json:"postalCode"`
	Region              string  `json:"region"`
	Status              string  `json:"status"`
	StatusDisplay       string  `json:"statusDisplay"`
	StreetAddress       string  `json:"streetAddress"`
	UpdateDate          string  `json:"updateDate"`
	Website             string  `json:"website"`
	YearOpened          string  `json:"yearOpened"`
}

type Brewery struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Established string `json:"established"`
	Organic     string `json:"isOrganic"`
}
