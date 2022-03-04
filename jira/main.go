package jira

// type Client struct {
// 	url   string
// 	token string
// 	user  string
// 	cache string
// }

// func (c *Client) ClearCache() {
// 	os.Remove(c.cache)
// }

// func (c *Client) LoadFromCache() ([]Issue, error) {
// 	data, err := ioutil.ReadFile(c.cache)
// 	if err != nil {
// 		return []Issue{}, errors.Wrap(err, "cannot read cache file")
// 	}

// 	var issues IssuesResponse

// 	err = json.Unmarshal(data, &issues)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return issues.Issues, nil
// }

// func (c *Client) Issues() []Issue {
// 	ret, err := c.LoadFromCache()
// 	if err == nil {
// 		return ret
// 	}

// 	url := c.url + "/rest/api/3/search?jql=assignee%20%3D%20currentUser()%20AND%20Status%20!%3D%20CLOSED%20AND%20Status%20!%3D%20DONE%20AND%20Status%20!%3D%20REJECTED%20AND%20Status%20!%3D%20COMPLETED"

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req.SetBasicAuth(c.user, c.token)

// 	client := &http.Client{}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var issues IssuesResponse

// 	err = json.Unmarshal(body, &issues)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = os.WriteFile(c.cache, body, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return issues.Issues
// }

// func New(url, user, token string) *Client {
// 	cache, err := xdg.CacheFile("gfeat")
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &Client{
// 		url:   url,
// 		user:  user,
// 		token: token,
// 		cache: cache,
// 	}
// }

// type IssuesResponse struct {
// 	Expand     string  `json:"expand"`
// 	Issues     []Issue `json:"issues"`
// 	MaxResults int64   `json:"maxResults"`
// 	StartAt    int64   `json:"startAt"`
// 	Total      int64   `json:"total"`
// }

// type Issue struct {
// 	Expand string `json:"expand"`
// 	Fields struct {
// 		Aggregateprogress struct {
// 			Progress int64 `json:"progress"`
// 			Total    int64 `json:"total"`
// 		} `json:"aggregateprogress"`
// 		Aggregatetimeestimate         interface{} `json:"aggregatetimeestimate"`
// 		Aggregatetimeoriginalestimate interface{} `json:"aggregatetimeoriginalestimate"`
// 		Aggregatetimespent            interface{} `json:"aggregatetimespent"`
// 		Assignee                      struct {
// 			AccountID   string `json:"accountId"`
// 			AccountType string `json:"accountType"`
// 			Active      bool   `json:"active"`
// 			AvatarUrls  struct {
// 				One6x16   string `json:"16x16"`
// 				Two4x24   string `json:"24x24"`
// 				Three2x32 string `json:"32x32"`
// 				Four8x48  string `json:"48x48"`
// 			} `json:"avatarUrls"`
// 			DisplayName  string `json:"displayName"`
// 			EmailAddress string `json:"emailAddress"`
// 			Self         string `json:"self"`
// 			TimeZone     string `json:"timeZone"`
// 		} `json:"assignee"`
// 		Components []struct {
// 			Description string `json:"description"`
// 			ID          string `json:"id"`
// 			Name        string `json:"name"`
// 			Self        string `json:"self"`
// 		} `json:"components"`
// 		Created string `json:"created"`
// 		Creator struct {
// 			AccountID   string `json:"accountId"`
// 			AccountType string `json:"accountType"`
// 			Active      bool   `json:"active"`
// 			AvatarUrls  struct {
// 				One6x16   string `json:"16x16"`
// 				Two4x24   string `json:"24x24"`
// 				Three2x32 string `json:"32x32"`
// 				Four8x48  string `json:"48x48"`
// 			} `json:"avatarUrls"`
// 			DisplayName  string `json:"displayName"`
// 			EmailAddress string `json:"emailAddress"`
// 			Self         string `json:"self"`
// 			TimeZone     string `json:"timeZone"`
// 		} `json:"creator"`
// 		Description struct {
// 			Content []struct {
// 				Attrs struct {
// 					Layout string `json:"layout"`
// 				} `json:"attrs"`
// 				Content []struct {
// 					Attrs struct {
// 						AccessLevel string `json:"accessLevel"`
// 						Collection  string `json:"collection"`
// 						Height      int64  `json:"height"`
// 						ID          string `json:"id"`
// 						Text        string `json:"text"`
// 						Type        string `json:"type"`
// 						URL         string `json:"url"`
// 						Width       int64  `json:"width"`
// 					} `json:"attrs"`
// 					Marks []struct {
// 						Attrs struct {
// 							Href string `json:"href"`
// 						} `json:"attrs"`
// 						Type string `json:"type"`
// 					} `json:"marks"`
// 					Text string `json:"text"`
// 					Type string `json:"type"`
// 				} `json:"content"`
// 				Type string `json:"type"`
// 			} `json:"content"`
// 			Type    string `json:"type"`
// 			Version int64  `json:"version"`
// 		} `json:"description"`
// 		Duedate     interface{} `json:"duedate"`
// 		Environment struct {
// 			Content []struct {
// 				Content []struct {
// 					Text string `json:"text"`
// 					Type string `json:"type"`
// 				} `json:"content"`
// 				Type string `json:"type"`
// 			} `json:"content"`
// 			Type    string `json:"type"`
// 			Version int64  `json:"version"`
// 		} `json:"environment"`
// 		FixVersions []interface{} `json:"fixVersions"`
// 		Issuelinks  []struct {
// 			ID          string `json:"id"`
// 			InwardIssue struct {
// 				Fields struct {
// 					Issuetype struct {
// 						AvatarID       int64  `json:"avatarId"`
// 						Description    string `json:"description"`
// 						EntityID       string `json:"entityId"`
// 						HierarchyLevel int64  `json:"hierarchyLevel"`
// 						IconURL        string `json:"iconUrl"`
// 						ID             string `json:"id"`
// 						Name           string `json:"name"`
// 						Self           string `json:"self"`
// 						Subtask        bool   `json:"subtask"`
// 					} `json:"issuetype"`
// 					Priority struct {
// 						IconURL string `json:"iconUrl"`
// 						ID      string `json:"id"`
// 						Name    string `json:"name"`
// 						Self    string `json:"self"`
// 					} `json:"priority"`
// 					Status struct {
// 						Description    string `json:"description"`
// 						IconURL        string `json:"iconUrl"`
// 						ID             string `json:"id"`
// 						Name           string `json:"name"`
// 						Self           string `json:"self"`
// 						StatusCategory struct {
// 							ColorName string `json:"colorName"`
// 							ID        int64  `json:"id"`
// 							Key       string `json:"key"`
// 							Name      string `json:"name"`
// 							Self      string `json:"self"`
// 						} `json:"statusCategory"`
// 					} `json:"status"`
// 					Summary string `json:"summary"`
// 				} `json:"fields"`
// 				ID   string `json:"id"`
// 				Key  string `json:"key"`
// 				Self string `json:"self"`
// 			} `json:"inwardIssue"`
// 			OutwardIssue struct {
// 				Fields struct {
// 					Issuetype struct {
// 						AvatarID       int64  `json:"avatarId"`
// 						Description    string `json:"description"`
// 						EntityID       string `json:"entityId"`
// 						HierarchyLevel int64  `json:"hierarchyLevel"`
// 						IconURL        string `json:"iconUrl"`
// 						ID             string `json:"id"`
// 						Name           string `json:"name"`
// 						Self           string `json:"self"`
// 						Subtask        bool   `json:"subtask"`
// 					} `json:"issuetype"`
// 					Priority struct {
// 						IconURL string `json:"iconUrl"`
// 						ID      string `json:"id"`
// 						Name    string `json:"name"`
// 						Self    string `json:"self"`
// 					} `json:"priority"`
// 					Status struct {
// 						Description    string `json:"description"`
// 						IconURL        string `json:"iconUrl"`
// 						ID             string `json:"id"`
// 						Name           string `json:"name"`
// 						Self           string `json:"self"`
// 						StatusCategory struct {
// 							ColorName string `json:"colorName"`
// 							ID        int64  `json:"id"`
// 							Key       string `json:"key"`
// 							Name      string `json:"name"`
// 							Self      string `json:"self"`
// 						} `json:"statusCategory"`
// 					} `json:"status"`
// 					Summary string `json:"summary"`
// 				} `json:"fields"`
// 				ID   string `json:"id"`
// 				Key  string `json:"key"`
// 				Self string `json:"self"`
// 			} `json:"outwardIssue"`
// 			Self string `json:"self"`
// 			Type struct {
// 				ID      string `json:"id"`
// 				Inward  string `json:"inward"`
// 				Name    string `json:"name"`
// 				Outward string `json:"outward"`
// 				Self    string `json:"self"`
// 			} `json:"type"`
// 		} `json:"issuelinks"`
// 		Issuetype struct {
// 			AvatarID       int64  `json:"avatarId"`
// 			Description    string `json:"description"`
// 			EntityID       string `json:"entityId"`
// 			HierarchyLevel int64  `json:"hierarchyLevel"`
// 			IconURL        string `json:"iconUrl"`
// 			ID             string `json:"id"`
// 			Name           string `json:"name"`
// 			Self           string `json:"self"`
// 			Subtask        bool   `json:"subtask"`
// 		} `json:"issuetype"`
// 		Labels     []string `json:"labels"`
// 		LastViewed string   `json:"lastViewed"`
// 		Parent     struct {
// 			Fields struct {
// 				Issuetype struct {
// 					AvatarID       int64  `json:"avatarId"`
// 					Description    string `json:"description"`
// 					EntityID       string `json:"entityId"`
// 					HierarchyLevel int64  `json:"hierarchyLevel"`
// 					IconURL        string `json:"iconUrl"`
// 					ID             string `json:"id"`
// 					Name           string `json:"name"`
// 					Self           string `json:"self"`
// 					Subtask        bool   `json:"subtask"`
// 				} `json:"issuetype"`
// 				Priority struct {
// 					IconURL string `json:"iconUrl"`
// 					ID      string `json:"id"`
// 					Name    string `json:"name"`
// 					Self    string `json:"self"`
// 				} `json:"priority"`
// 				Status struct {
// 					Description    string `json:"description"`
// 					IconURL        string `json:"iconUrl"`
// 					ID             string `json:"id"`
// 					Name           string `json:"name"`
// 					Self           string `json:"self"`
// 					StatusCategory struct {
// 						ColorName string `json:"colorName"`
// 						ID        int64  `json:"id"`
// 						Key       string `json:"key"`
// 						Name      string `json:"name"`
// 						Self      string `json:"self"`
// 					} `json:"statusCategory"`
// 				} `json:"status"`
// 				Summary string `json:"summary"`
// 			} `json:"fields"`
// 			ID   string `json:"id"`
// 			Key  string `json:"key"`
// 			Self string `json:"self"`
// 		} `json:"parent"`
// 		Priority struct {
// 			IconURL string `json:"iconUrl"`
// 			ID      string `json:"id"`
// 			Name    string `json:"name"`
// 			Self    string `json:"self"`
// 		} `json:"priority"`
// 		Progress struct {
// 			Progress int64 `json:"progress"`
// 			Total    int64 `json:"total"`
// 		} `json:"progress"`
// 		Project struct {
// 			AvatarUrls struct {
// 				One6x16   string `json:"16x16"`
// 				Two4x24   string `json:"24x24"`
// 				Three2x32 string `json:"32x32"`
// 				Four8x48  string `json:"48x48"`
// 			} `json:"avatarUrls"`
// 			ID             string `json:"id"`
// 			Key            string `json:"key"`
// 			Name           string `json:"name"`
// 			ProjectTypeKey string `json:"projectTypeKey"`
// 			Self           string `json:"self"`
// 			Simplified     bool   `json:"simplified"`
// 		} `json:"project"`
// 		Reporter struct {
// 			AccountID   string `json:"accountId"`
// 			AccountType string `json:"accountType"`
// 			Active      bool   `json:"active"`
// 			AvatarUrls  struct {
// 				One6x16   string `json:"16x16"`
// 				Two4x24   string `json:"24x24"`
// 				Three2x32 string `json:"32x32"`
// 				Four8x48  string `json:"48x48"`
// 			} `json:"avatarUrls"`
// 			DisplayName  string `json:"displayName"`
// 			EmailAddress string `json:"emailAddress"`
// 			Self         string `json:"self"`
// 			TimeZone     string `json:"timeZone"`
// 		} `json:"reporter"`
// 		Resolution struct {
// 			Description string `json:"description"`
// 			ID          string `json:"id"`
// 			Name        string `json:"name"`
// 			Self        string `json:"self"`
// 		} `json:"resolution"`
// 		Resolutiondate string      `json:"resolutiondate"`
// 		Security       interface{} `json:"security"`
// 		Status         struct {
// 			Description    string `json:"description"`
// 			IconURL        string `json:"iconUrl"`
// 			ID             string `json:"id"`
// 			Name           string `json:"name"`
// 			Self           string `json:"self"`
// 			StatusCategory struct {
// 				ColorName string `json:"colorName"`
// 				ID        int64  `json:"id"`
// 				Key       string `json:"key"`
// 				Name      string `json:"name"`
// 				Self      string `json:"self"`
// 			} `json:"statusCategory"`
// 		} `json:"status"`
// 		Statuscategorychangedate string `json:"statuscategorychangedate"`
// 		Subtasks                 []struct {
// 			Fields struct {
// 				Issuetype struct {
// 					AvatarID       int64  `json:"avatarId"`
// 					Description    string `json:"description"`
// 					EntityID       string `json:"entityId"`
// 					HierarchyLevel int64  `json:"hierarchyLevel"`
// 					IconURL        string `json:"iconUrl"`
// 					ID             string `json:"id"`
// 					Name           string `json:"name"`
// 					Self           string `json:"self"`
// 					Subtask        bool   `json:"subtask"`
// 				} `json:"issuetype"`
// 				Priority struct {
// 					IconURL string `json:"iconUrl"`
// 					ID      string `json:"id"`
// 					Name    string `json:"name"`
// 					Self    string `json:"self"`
// 				} `json:"priority"`
// 				Status struct {
// 					Description    string `json:"description"`
// 					IconURL        string `json:"iconUrl"`
// 					ID             string `json:"id"`
// 					Name           string `json:"name"`
// 					Self           string `json:"self"`
// 					StatusCategory struct {
// 						ColorName string `json:"colorName"`
// 						ID        int64  `json:"id"`
// 						Key       string `json:"key"`
// 						Name      string `json:"name"`
// 						Self      string `json:"self"`
// 					} `json:"statusCategory"`
// 				} `json:"status"`
// 				Summary string `json:"summary"`
// 			} `json:"fields"`
// 			ID   string `json:"id"`
// 			Key  string `json:"key"`
// 			Self string `json:"self"`
// 		} `json:"subtasks"`
// 		Summary              string        `json:"summary"`
// 		Timeestimate         interface{}   `json:"timeestimate"`
// 		Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
// 		Timespent            interface{}   `json:"timespent"`
// 		Updated              string        `json:"updated"`
// 		Versions             []interface{} `json:"versions"`
// 		Votes                struct {
// 			HasVoted bool   `json:"hasVoted"`
// 			Self     string `json:"self"`
// 			Votes    int64  `json:"votes"`
// 		} `json:"votes"`
// 		Watches struct {
// 			IsWatching bool   `json:"isWatching"`
// 			Self       string `json:"self"`
// 			WatchCount int64  `json:"watchCount"`
// 		} `json:"watches"`
// 		Workratio int64 `json:"workratio"`
// 	} `json:"fields"`
// 	ID   string `json:"id"`
// 	Key  string `json:"key"`
// 	Self string `json:"self"`
// }
