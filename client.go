package brewerydb

import(
  "net/http"
)

var base string = "https://api.brewerydb.com/v2/"
var apiKey string

type client struct {
  apiKey string
  base string
}


func (*client) _get(fullPath string) (res *http.Response, err error) {
  return http.Get(fullPath)
}

