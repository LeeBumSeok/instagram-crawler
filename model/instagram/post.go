package instagram

type Post struct {
	Graphql struct {
		ShortcodeMedia struct {
			Owner struct {
				Username                 string `json:"username"`
				Fullname                 string `json:"full_name"`
				EdgeOwnerToTimelineMedia struct {
					Count int `json:"count"`
				} `json:"edge_owner_to_timeline_media"`
				EdgeFollowBy struct {
					Count int `json:"count"`
				} `json:"edge_followed_by"`
			} `json:"owner"`
			Location struct {
				Name        string `json:"name"`
				AddressJson string `json:"address_json"`
			} `json:"location"`
		} `json:"shortcode_media"`
	} `json:"graphql"`
}
