package git

// func userCacheDump(cache map[string]string) {
// 	cacheJSON, err := json.MarshalIndent(cache, "", "    ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	cachePath := CacheLocation()

// 	err = os.WriteFile(cachePath, cacheJSON, 0644)
// 	if err != nil {
// 		log.WithFields(log.Fields{
// 			"err":       err,
// 			"cachePath": cachePath,
// 		}).Error("cannot write cache file")
// 	}
// }

// func userCacheLoad() map[string]string {
// 	cachePath := CacheLocation()

// 	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
// 		return map[string]string{}
// 	}

// 	data, err := ioutil.ReadFile(cachePath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var cache map[string]string

// 	err = json.Unmarshal(data, &cache)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return cache
// }

// // CacheLocation Provides a cache location for user emails.
// func CacheLocation() string {
// 	cachePath, err := xdg.CacheFile("gi-reviewers/emails.json")
// 	if err != nil {
// 		log.WithFields(log.Fields{
// 			"err": err,
// 		}).Panic("cannot get a cache path")
// 	}

// 	return cachePath
// }
