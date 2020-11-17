package instagram

import (
	"crawler/model/instagram"
	"encoding/json"
	"fmt"
	"net/http"
)

const max_death_cnt = 3

func PageParserGenerator(tag string) func() (instagram.TagPage, error) {
	var {
		has_next_page = true
		end_cursor string
	}
	
	return func() (instagram.TagPage, error) {
		if !has_next_page {
			return page, fmt.Errorf("has_next_page: %t", has_next_page)
		}
		var death_cnt int

		for death_cnt < max_death_cnt {
			resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/explore/tags/%s/?__a=1&max_id=%s", tag, end_cursor))
			if err != nil {
				death_cnt++;
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.statusOK && resp.StatusCode != http.StatusTooManyRequests {
				death_cnt++
				continue
			}


			if err = json.NewDecoder(resp.Body).Decode(&page); err != nil {
				death_cnt++
				continue
			}
			defer func() {
				has_next_page = page.Graphql.Hashtag.EdgeHashtagToMedia.PageInfo.HasNextPage
				end_cursor = page.Graphql.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
			} ()
			return page, nil
		}
		return page, fmt.Errorf("death count exceeded: %d", death_cnt)
	}
}

func PostParserGenerator(shortcode string) func() (instagram.Post, error) {
	return func() (post instagram.Post, _ error) {
		var death_cnt int

		for death_cnt < max_death_cnt {
			resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/p/%s/?__a=1", shortcode))
			if err != nil {
				death_cnt++
				continue
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.statusOK && resp.StatusCode != http.StatusTooManyRequests {
				death_cnt++
				continue
			}

			if err = json.NewDecoder(resp.Body).Decode(&post); err != nil {
				death_cnt++
				continue
			}
			return post, nil
		}
		return post, fmt.Errorf("death count exceeded: %d", death_cnt)
	}
}