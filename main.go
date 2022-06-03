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

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("gocomi version 2, fisik_yum\npress enter to select default values")
	var sKey string //will store id later
	selec := -1
	fmt.Print("\nEnter Comic Strip Name:")
	fmt.Scanf("%s", &sKey)
	hList := search(sKey)
	fmt.Printf("\nFound %d results\n\n", len(hList))
	for x := range hList {
		fmt.Printf("[%d] %s\n", x, hList[x].name)
	}
	for selec < 0 || selec >= len(hList) {
		fmt.Printf("\nSelect entry [0-%d]:", len(hList)-1)
		fmt.Scanf("%d\n", &selec)
	}

	sKey = hList[selec].id[1:] //shamelessy reusing variables
	var startDate time.Time = startDate(sKey)
	var dKey string

	fmt.Printf("Enter scrape start date[yyyy/mm/dd][default=%d/%d/%d]:", startDate.Year(), startDate.Month(), startDate.Day())
	fmt.Scanf("%s", &dKey)
	if dKey != "" {
		startDate = parsedate(dKey)
	}
	var endDate time.Time
	fmt.Printf("Enter scrape start date[yyyy/mm/dd][default=%d/%d/%d]:", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Scanf("%s", &dKey)

	if dKey != "" {
		endDate = parsedate(dKey)
	} else {
		endDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	}
	dlroutine(sKey, startDate, endDate)
	fmt.Println("Operation Complete!")
}
