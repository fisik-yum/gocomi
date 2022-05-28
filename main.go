package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

func main() {

	var name string
	var dresp string
	fmt.Print("comic name (this should be in the format gocomics uses gocomics.com/name/):")
	fmt.Scanf("%s", &name)

	start := findStartDate(name)
	end := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	fmt.Println("Due to the nature of the gocomics website, a comic's end date cannot be determined.")

	fmt.Print("Do you want to input a custom stop date[yes/no][no]?")
	fmt.Scanf("%s", &dresp)

	if dresp == "yes" {
		end = getEndDate()
	}
	fmt.Printf("\nSet threshold date to %d/%d/%d \n", end.Year(), end.Month(), end.Day())

	fmt.Print("Do you want to save links to a file or download images[link/download][download]")
	fmt.Scanf("%s", &dresp)

	var ctrlLink string

	if dresp == "link" { //save to file. not very useful, but ill keep it as an option anyway.
		f, err := os.OpenFile(fmt.Sprintf("%s_%d-%d-%d_%d:%d.txt", name, time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute()), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		defer f.Close()
		for start != end.AddDate(0, 0, 1) {
			fmt.Printf("\r %d/%d/%d", start.Year(), start.Month(), start.Day())
			time.Sleep(1000 * time.Millisecond)
			lnk := getLink(name, start)
			if ctrlLink == "" {
				ctrlLink = lnk
			} else if lnk == ctrlLink {
				break
			}
			fmt.Fprintf(f, "%s\n", lnk)
			start = start.AddDate(0, 0, 1)
		}
	} else { //default option, download images

		_, err := os.Stat(name)
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(name, 0755)
			if errDir != nil {
				log.Fatal(err)
			}
		}
		////DO NOT EDIT (for now)
		for start != end.AddDate(0, 0, 1) {
			fmt.Printf("\r %d/%d/%d", start.Year(), start.Month(), start.Day())
			time.Sleep(1000 * time.Millisecond)
			lnk := getLink(name, start)
			if ctrlLink == "" {
				ctrlLink = lnk
			} else if lnk == ctrlLink {
				break
			}
			err := DownloadFile(strings.Join([]string{name, "/", fmt.Sprintf("%d-%d-%d", start.Year(), start.Month(), start.Day()), ".jpg"}, ""), lnk)
			if err != nil {
				log.Fatal(err)
			}
			start = start.AddDate(0, 0, 1)
		}
	}
	fmt.Println("Operation Complete!")
}

//this function is core functionality
func getLink(name string, start time.Time) string {
	resp, err := soup.Get(fmt.Sprintf("https://www.gocomics.com/%s/%d/%d/%d", name, start.Year(), start.Month(), start.Day()))
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	link := doc.Find("picture", "class", "item-comic-image").Find("img")
	return link.Attrs()["src"]
}

func findStartDate(name string) time.Time {
	resp, err := soup.Get(fmt.Sprintf("https://www.gocomics.com/%s/%d/%d/%d", name, time.Now().Year(), time.Now().Month(), time.Now().Day()))
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	date := doc.Find("div", "class", "gc-calendar-nav__select").Find("div").Attrs()["data-start"]
	fmt.Println("Start Date " + date)
	fds := strings.Split(date, "/")
	var l [3]int
	for i := range l {
		l[i], _ = strconv.Atoi(fds[i])
	}
	return time.Date(l[0], time.Month(l[1]), l[2], 0, 0, 0, 0, time.Now().Location())
}

func getEndDate() time.Time {
	var yr int
	var mt int
	var dy int

	fmt.Print("Enter year [YYYY]:")
	fmt.Scanf("%d", &yr)
	fmt.Print("Enter month [MM]:")
	fmt.Scanf("%d", &mt)
	fmt.Print("Enter day [DD]:")
	fmt.Scanf("%d", &dy)
	return time.Date(yr, time.Month(mt), dy, 0, 0, 0, 0, time.Now().Location())
}

func DownloadFile(filepath string, url string) error {
	//blatantly stolen from https://golangcode.com/download-a-file-from-a-url/
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
