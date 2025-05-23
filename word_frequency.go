package main

import (
	"fmt"
	"strings"
)

func word_frequency_counter(text string) map[string]uint {

	mapp := make(map[string]uint)

	words := strings.Split(text, " ")

	for _,word := range words {
		if mapp[word] != 0 {
			mapp[word] = mapp[word] + 1

		}else{
			mapp[word] = 1
		}
	}

	return mapp

}

func main() {
	fmt.Println("The final Out-put was \n", word_frequency_counter("phani is a super man, who is phani"))
}