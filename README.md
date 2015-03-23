# BreweryDB API Wrapper in Golang

[![Build Status](https://travis-ci.org/hobbeswalsh/brewerydb.svg?branch=master)](https://travis-ci.org/hobbeswalsh/brewerydb)

## To Use:

```go
    package main

    import (
    	"github.com/michaelporter/brewerydb"
    )
    
    c := brewerydb.NewClient("myApiKey")

    // I want a random beer. No constraints on the query.
    beer, err := c.RandomBeer(brewerydb.BeerQuery{})

    // I want a list of breweries near some zip code
    lq := brewerydb.LocationQuery{}
    lq.PostalCode = "10101"
    breweries, err := c.GetLocation(lq)

    // I want a strong beer; it's winter time.
    bq := brewerydb.BeerQuery{}
    bq.Abv = "8,10" // Anything between 8 and 10 will do
    beers, err := c.GetBeers(bq)

```