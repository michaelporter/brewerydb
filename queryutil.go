package brewerydb

func _makeQueryString(queryParams map[string]string) (queryString string) {
  for k, v := range queryParams {
    queryString = queryString + "&" + k + "=" + v
  }

  return
}

func _combinePath(path, queryString string) (fullPath string) {
  fullPath = base + version + path + apiKey + queryString

  return
}

