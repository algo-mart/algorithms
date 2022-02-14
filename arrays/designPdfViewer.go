package main

import "fmt"

/*
https://www.hackerrank.com/challenges/designer-pdf-viewer/problem?h_r=internal-search&isFullScreen=true
*/
func designerPdfViewer(h []int32, word string) int32 {
	var mymap = map[rune]int{}
	char := 'a'
	max := 0
	for _,j:= range h {
		mymap[char] = int(j)
		char++
	}
	for _,j:= range word {
		if mymap[j]> max {
			max = mymap[j]
		}
	}
	return int32(max * len(word))

}

func designerPdfViewer1(h []int32, word string) int32 {
	max:= int32(0)
	for _, j := range word {
		if max < h[j-'a']{
			max = h[j-'a']
		}
	}
	return max * int32(len(word))
}

func main () {
	//mymap :=make(map[string]int)
	slice := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word := "zaba"
	fmt.Print(designerPdfViewer(slice, word))
	fmt.Print(designerPdfViewer1(slice, word))
}