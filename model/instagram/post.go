package instagram

type Post struct {
	Graphql struct {
		ShortcodeMedia struct {
			DisplayURL           string `json:"display_url"`
			EdgeMediaPreviewLike struct {
				Count int `json:"count"`
			} `json:"edge_media_preview_like"`
			Location struct {
				Name        string `json:"name"`
				AddressJSON string `json:"address_json"`
			} `json:"location"`
		} `json:"shortcode_media"`
	} `json:"graphql"`
}
