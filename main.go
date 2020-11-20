package main

import (
	"container/heap"
	"crawler/dataservice/instagram"
	insta "crawler/model/instagram"
	"fmt"
	"sync"
)

// type postHeap []insta.Post

// func (p postHeap) Len() int { return len(p) }
// func (p postHeap) Less(i, j int) bool {
// 	return p[i].Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count < p[j].Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count
// }
// func (p postHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
// func (p *postHeap) Push(x interface{}) { *p = append(*p, x.(insta.Post)) }
// func (p *postHeap) Pop() interface{} {
// 	defer func() { *p = (*p)[:len(*p)-1] }()
// 	return (*p)[len(*p)-1]
// }

type postHeap []insta.Post

func (p postHeap) Len() int { return len(p) }
func (p postHeap) Less(i, j int) bool {
	return p[i].Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count > p[j].Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count
}
func (p postHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *postHeap) Push(x interface{}) { *p = append(*p, x.(insta.Post)) }
func (p *postHeap) Pop() interface{} {
	defer func() { *p = (*p)[:len(*p)-1] }()
	return (*p)[len(*p)-1]
}

func main() {
	// const hashtag = "seoul"

	// parser := instagram.PageParserGenerator(hashtag)

	// posts := make(chan insta.Post)

	// var syncer sync.WaitGroup

	// for i := 0; i < 3; i++ {
	// 	page, err := parser()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	shortcodes := page.Shortcodes()

	// 	workers := make([]func() (insta.Post, error), len(shortcodes))

	// 	for i, shortcode := range shortcodes {
	// 		workers[i] = instagram.PostParserGenerator(shortcode)
	// 	}

	// 	for _, worker := range workers {
	// 		syncer.Add(1)
	// 		go func(worker func() (insta.Post, error)) {
	// 			defer syncer.Done()

	// 			post, err := worker()
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			} else {
	// 				posts <- post
	// 			}
	// 		}(worker)
	// 	}
	// }

	// go func() {
	// 	defer close(posts)
	// 	syncer.Wait()
	// }()

	// var postHeap postHeap

	// for post := range posts {
	// 	heap.Push(&postHeap, post)
	// }

	// for 0 < postHeap.Len() {
	// 	fmt.Println(heap.Pop(&postHeap).(insta.Post).Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count)
	// }

	const hashtag = "seoul"

	parser := instagram.PageParserGenerator(hashtag)

	posts := make(chan insta.Post)

	var syncer sync.WaitGroup

	var postHeap postHeap

	for i := 0; i < 3; i++ {
		page, err := parser()
		if err != nil {
			fmt.Println(err)
		}

		shortcodes := page.Shortcodes()

		workers := make([]func() (insta.Post, error), len(shortcodes))

		for i, shortcode := range shortcodes {
			workers[i] = instagram.PostParserGenerator(shortcode)
		}

		for _, worker := range workers {
			syncer.Add(1)
			go func(worker func() (insta.Post, error)) {
				defer syncer.Done()

				post, err := worker()
				if err != nil {
					fmt.Println(err)
				} else {
					posts <- post
				}
			}(worker)
		}
	}

	go func() {
		defer close(posts)
		syncer.Wait()
	}()

	for post := range posts {
		heap.Push(&postHeap, post)
	}

	for 0 < postHeap.Len() {
		post := heap.Pop(&postHeap).(insta.Post)
		fmt.Printf("%d : %s\n", post.Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count, post.Graphql.ShortcodeMedia.DisplayURL)
	}
}
