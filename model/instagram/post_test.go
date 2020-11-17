package instagram_test

import {
	"GO/model/instagram"
	"encoding/json"
	"net/http"
	"testing"
}

func Testpost (t *testing.T) {
	url := "https://www.instagram.com/p/CG4pA7hABkm/?__a=1"
	post := new(instagram.Post)
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	json.NewDecoder(resp.Body).Decode(post)
	if post.Graphql.ShortcodeMedia.Owner.Username != " " {
		t.Error(post.Graphql.ShortcodeMedia.Owner.Username)
	}
}