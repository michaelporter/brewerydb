package brewerydb

import(
  "net/http"
)

var base string = "https://api.brewerydb.com/"
var version string = "v2"

var apiKey string

func Setkey(k string) {
  apiKey = "?key=" + k
  return
}

func Get(path string, queryParams map[string]string) (res *http.Response, err error) {
  queryString := _makeQueryString(queryParams)
  fullPath    := _combinePath(path, queryString)

  res, err = http.Get(fullPath)

  return
}

