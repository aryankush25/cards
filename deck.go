package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}
