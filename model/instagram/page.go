package instagram

// type TagPage struct {
// 	Graphql struct {
// 		Hashtag struct {
// 			EdgeHashtagToMedia struct {
// 				PageInfo struct {
// 					HasNextPage bool   `json:"has_next_page"`
// 					EndCursor   string `json:"end_cursor"`
// 				} `json:"page_info"`
// 				Edges []struct {
// 					Node struct {
// 						Shortcode          string `json:"shortcode"`
// 						EdgeMediaToCaption struct {
// 							Edges []struct {
// 								Node struct {
// 									Text string `json:"text"`
// 								} `json:"node"`
// 							} `json:"edges"`
// 						} `json:"edge_media_to_caption"`
// 					} `json:"node"`
// 				} `json:"edges"`
// 			} `json:"edge_hashtag_to_media"`
// 		} `json:"hashtag"`
// 	} `json:"graphql"`
// }

// func (t TagPage) Shortcodes() []string {
// 	shortcodes := make([]string, len(t.Graphql.Hashtag.EdgeHashtagToMedia.Edges))

// 	for i, edge := range t.Graphql.Hashtag.EdgeHashtagToMedia.Edges {
// 		shortcodes[i] = edge.Node.Shortcode
// 	}

// 	return shortcodes
// }

type TagPage struct {
	Graphql struct {
		Hashtag struct {
			EdgeHashtagToMedia struct {
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
				Edges []struct {
					Node struct {
						EdgeMediaToCaption struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						Shortcode string `json:"shortcode"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_hashtag_to_media"`
		} `json:"hashtag"`
	} `json:"graphql"`
}

func (t TagPage) Shortcodes() []string {
	shortcodes := make([]string, len(t.Graphql.Hashtag.EdgeHashtagToMedia.Edges))

	for i, edge := range t.Graphql.Hashtag.EdgeHashtagToMedia.Edges {
		shortcodes[i] = edge.Node.Shortcode
	}

	return shortcodes
}
