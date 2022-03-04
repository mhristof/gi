package keychain

// func Item(name string) string {
// 	env := os.Getenv(name)
// 	if env != "" {
// 		return env
// 	}

// 	query := kchain.NewItem()
// 	query.SetSecClass(kchain.SecClassGenericPassword)
// 	query.SetService("germ")
// 	query.SetAccount(name)
// 	query.SetAccessGroup("germ")
// 	// query.SetMatchLimit(kchain.MatchLimitOne)
// 	query.SetReturnData(true)
// 	results, err := kchain.QueryItem(query)
// 	if err != nil {
// 		log.WithFields(log.Fields{
// 			"err": err,
// 		}).Debug("cannot query keychain")

// 		return ""
// 	} else if len(results) != 1 {
// 		log.WithFields(log.Fields{
// 			"results": results,
// 		}).Debug("item not found")

// 		return ""
// 	}

// 	// data will be in the form of
// 	// export GITLAB_TOKEN='token'
// 	return strings.Split(string(results[0].Data), "'")[1]
// }
