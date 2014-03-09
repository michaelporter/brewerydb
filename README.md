# BreweryDB API Wrapper in Golang

## To Use:

```
    import (
    	"github.com/michaelporter/brewerydb"
    )

    func getBeers() {
      brewerydb.SetApiKey("your-api-key")
      
      queryParams := map[string]string{
      	'abv': '6',
      }
      res, err := brewerydb.Get("/beers", queryParams)
    }

```

## To-Do
- Pretty much everything
