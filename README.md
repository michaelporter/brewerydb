# BreweryDB API Wrapper in Golang

## To Use:

```
    import (
    	"github.com/michaelporter/brewerydb"
    )

    func main() {
      brewerydb.Setkey("your-api-key")
      
      queryParams := map[string]string { "key" : "value" }
      res, err := brewerydb.Get("beers/beer-id", queryParams)
    }

```

## To-Do
- Pretty much everything