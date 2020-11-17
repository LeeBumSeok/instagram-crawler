package main

import (
	"crawler/dataservice/instagram"
	"fmt"
	"log"
)

// func main() {
// 	const tag = "fff"
// 	var pages []instagram.TagPage

// 	url := fmt.Sprintf("https://www.instagram.com/explore/tage/%s/?__a=1&max_id=%s", tag, end_cursor)

// 	for i := 0; i<10; i++ {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		defer resp.Body.Close()

// 		var page instagram.TagPage

// 		iff err = json.NewDecoder(resp.Body).Decode(&page); err != nil {
// 			log.Fatalln(err)
// 		}
// 		end_cursor = page.Graphql.Hashtag.EdgeHashtagToMedia.PageInfo.EndCursor
// 		pages = append(pages, page)
// 	}

// 	for _, page := range pages {
// 		fmt.Println(page.Shortcodes())
// 	}
// }

// func main() {
// 	const tag = "fff"
// 	parser := instagram.PageParserGenerator(tag)
// 	for i := 0; i < 10; i++ {
// 		page, err := parser()
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		for _. shortcode := range page.Shortcodes() {
// 			fmt.Println(shortcode)
// 		}
// 	}
// }

// func main() {
// 	tags := []string{"fff". "f4f", "ootd"}
// 	workers := [] func() (insta.TagPage, error) {}
// 	for _, tag := range tags {
// 		worker = instagram.PageParserGenerator(tag)
// 		workers = append(workers, worker)
// 	}

// 	for i := 0; i < 2; i++ {
// 		parser = workers[i]
// 		for j := 0; j < 10; j++ {
// 			page, err := worker()
// 			if err != nil {
// 				log.Fatalln(err)
// 			}
// 			for _, shortcode := range page.Shortcodes() {
// 				fmt.Println(shortcode)
// 			}
// 		}
// 	}
// }

func main() {
	const tag = "fff"
	parser = instagram.PageParserGenerator(tag)
	page, err := parser()

	if err != nil {
		log.Fatalln(err)
	}

	for _, shortcode := range page.Shortcodes() {
		post, err := instagram.PostParserGenerator(shortcode)()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(post.Graphql.ShortcodeMedia.Owner.Username)
	}
}
