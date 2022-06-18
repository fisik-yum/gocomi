package gocomics

import (
	"fmt"
	"gocomi/base"
	"time"
)

func START() {
	var sKey string //will store id later
	selec := -1
	fmt.Print("\nEnter Comic Strip Name:")
	fmt.Scanf("%s", &sKey)
	hList := search(sKey)
	fmt.Printf("\nFound %d results\n\n", len(hList))
	for x := range hList {
		fmt.Printf("[%d] %s\n", x, hList[x].Name)
	}
	for selec < 0 || selec >= len(hList) {
		fmt.Printf("\nSelect entry [0-%d]:", len(hList)-1)
		fmt.Scanf("%d\n", &selec)
	}

	sKey = hList[selec].ID[1:] //shamelessy reusing variables
	var startDate time.Time = startDate(sKey)
	var dKey string

	fmt.Printf("Enter scrape start date[yyyy/mm/dd][default=%d/%d/%d]:", startDate.Year(), startDate.Month(), startDate.Day())
	fmt.Scanf("%s", &dKey)
	if dKey != "" {
		startDate = parsedate(dKey)
	}
	var endDate time.Time
	fmt.Printf("Enter scrape end date[yyyy/mm/dd][default=%d/%d/%d]:", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Scanf("%s", &dKey)

	if dKey != "" {
		endDate = parsedate(dKey)
	} else {
		endDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	}
	/*dl := base.DLdata{
		ID:    sKey,
		Start: startDate.Format(time.UnixDate),
		Stop:  endDate.Format(time.UnixDate),
	}*/
	dlroutine(base.DLdata{
		ID:    sKey,
		Start: startDate.Format(time.UnixDate),
		Stop:  endDate.Format(time.UnixDate),
	})
	fmt.Println("Operation Complete!")
}
