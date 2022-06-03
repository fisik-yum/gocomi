package main

import (
	"fmt"
	"log"

	"github.com/anaskhan96/soup"
)

func search(term string) (rList []result) { //search functions
	resp, err := soup.Get(fmt.Sprintf("https://www.gocomics.com/search/results?&terms=%s", term))
	check(err)
	doc := soup.HTMLParse(resp)
	rs := doc.FindAll("div", "class", "content-section-sm")
	for x := range rs {
		rList = append(rList, result{
			name: rs[x].Find("img").Attrs()["alt"],
			id:   rs[x].Find("a").Attrs()["href"], /*[1:]*/
		})
	}
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type result struct {
	name string
	id   string
}
