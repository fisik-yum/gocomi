package main

import "fmt"

func main() {
	var sKey string
	selec := -1
	fmt.Print("Enter Comic Strip Name:")
	fmt.Scanf("%s", &sKey)
	hList := search(sKey)
	fmt.Printf("\nFound %d results\n\n", len(hList))
	for x := range hList {
		fmt.Printf("[%d] %s\n", x, hList[x].name)
	}
	for selec < 0 || selec >= len(hList) {
		fmt.Printf("\nSelect the entry number[0-%d]:", len(hList)-1)
		fmt.Scanf("%d\n", &selec)
	}
	fmt.Println(hList[selec].id)
}
