package brewerydb

func _makeQueryString(queryParams map[string]string) (queryString string) {
  for k, v := range queryParams {
    queryString = queryString + "&" + k + "=" + v
  }

  return
}

func _combinePath(path, queryString string) (fullPath string) {
  return base + path + apiKey + queryString
}

