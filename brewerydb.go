package brewerydb

import(
  "net/http"
)

var c *client

func NewClient(k string) (c *client) {
  return &client{apiKey:apiKey, base: base}
}

func SetApiKey(k string) (bool){
  apiKey = "?key=" + k
  c = NewClient(apiKey)
  return c != nil
}

func Get(path string, queryParams map[string]string) (res *http.Response, err error) {
  queryString := _makeQueryString(queryParams)
  fullPath    := _combinePath(path, queryString)

  return c._get(fullPath)
}

