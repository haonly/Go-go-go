package main

import (
	"./preprocess"
)

func main() {
	var brandNames = []string{"audi", "bmw", "ford", "hyundi", "merc", "skoda", "toyota", "vauxhall", "vw"}
	preprocess.Preprocess(brandNames)
}
