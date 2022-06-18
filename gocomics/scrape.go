/*
   gocomi- bulk downloader for gocomics!
   Copyright (C) 2022  fisik_yum
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package gocomics

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gocomi/base"

	"github.com/anaskhan96/soup"
)

func search(term string) (rList []base.Result) { //search functions
	resp, err := soup.Get(fmt.Sprintf("https://www.gocomics.com/search/results?&terms=%s", term))
	check(err)
	doc := soup.HTMLParse(resp)
	rs := doc.FindAll("div", "class", "content-section-sm")
	for x := range rs {
		rList = append(rList, base.Result{
			Name: rs[x].Find("img").Attrs()["alt"],
			ID:   rs[x].Find("a").Attrs()["href"], /*[1:]*/
		})
	}
	return
}

func startDate(name string) time.Time {
	resp, err := soup.Get(fmt.Sprintf("https://www.gocomics.com/%s/%d/%d/%d", name, time.Now().Year(), time.Now().Month(), time.Now().Day()-1))
	check(err)
	doc := soup.HTMLParse(resp)
	date := doc.Find("div", "class", "gc-calendar-nav__select").Find("div").Attrs()["data-start"]
	//fmt.Println(date)
	return parsedate(date)

}

func parsedate(date string) time.Time { //for other stuff too
	fds := strings.Split(date, "/")
	var l [3]int
	for i := range l {
		l[i], _ = strconv.Atoi(fds[i])
	}
	return time.Date(l[0], time.Month(l[1]), l[2], 0, 0, 0, 0, time.Now().Location())
}

func dlroutine(data base.DLdata) {

	start, err := time.Parse(time.UnixDate, data.Start)
	check(err)
	end, err := time.Parse(time.UnixDate, data.Stop)
	check(err)

	var ctrlLink string
	_, err = os.Stat(data.ID)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(data.ID, 0755)
		check(errDir)
	}

	for start != end.AddDate(0, 0, 1) {
		fmt.Printf("\r %d/%d/%d", start.Year(), start.Month(), start.Day())
		time.Sleep(500 * time.Millisecond)
		lnk := getLink(data.ID, start)
		if ctrlLink == "" {
			ctrlLink = lnk
		} else if lnk == ctrlLink {
			break
		}
		err := base.DownloadFile(strings.Join([]string{data.ID, "/", fmt.Sprintf("%d-%d-%d", start.Year(), start.Month(), start.Day()), ".jpg"}, ""), lnk)
		if err != nil {
			log.Fatal(err)
		}
		start = start.AddDate(0, 0, 1)
	}
}

func getLink(name string, start time.Time) string {
	resp, err := soup.Get(fmt.Sprintf("https://www.gocomics.com/%s/%d/%d/%d", name, start.Year(), start.Month(), start.Day()))
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	link := doc.Find("picture", "class", "item-comic-image").Find("img")
	return link.Attrs()["src"]
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
