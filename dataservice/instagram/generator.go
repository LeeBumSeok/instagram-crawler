package instagram

import (
	"crawler/model/instagram"
	"encoding/json"
	"fmt"
	"net/http"
)

const max_death_cnt = 3

func PageParserGenerator(hashtag string) func() (instagram.TagPage, error) {
	var (
		has_next_page = true
		end_cursor    string
	)

	return func() (page instagram.TagPage, _ error) {
		if !has_next_page {
			return page, fmt.Errorf("has_next_page: %t", has_next_page)
		}

		var death_cnt int

		for death_cnt < max_death_cnt {
			resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/explore/tags/%s/?__a=1&max_id=%s", hashtag, end_cursor))

			if err != nil {
				death_cnt++
				continue // RETRY INSTRUCTION
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTooManyRequests {
				death_cnt++
				continue // RETRY INSTRUCTION
			}

			if err = json.NewDecoder(resp.Body).Decode(&page); err != nil {
				death_cnt++
				continue // RETRY INSTRUCTION
			}

			return page, nil
		}
		return page, fmt.Errorf("death_cnt: %d", death_cnt)
	}
}

func PostParserGenerator(shortcode string) func() (instagram.Post, error) {
	return func() (post instagram.Post, _ error) {
		var death_cnt int

		for death_cnt < max_death_cnt {
			resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/p/%s/?__a=1", shortcode))

			if err != nil {
				death_cnt++
				continue // RETRY INSTRUCTION
			}

			defer resp.Body.Close()

			if resp.StatusCode == http.StatusNotFound {
				return post, fmt.Errorf("status code: %d", resp.StatusCode)
			}

			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTooManyRequests {
				death_cnt++
				continue // RETRY INSTRUCTION
			}

			if err = json.NewDecoder(resp.Body).Decode(&post); err != nil {
				death_cnt++
				continue // RETRY INSTRUCTION
			}

			return post, nil
		}
		return post, fmt.Errorf("death_cnt: %d", death_cnt)
	}
}
