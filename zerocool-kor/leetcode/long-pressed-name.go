package main

import (
	"log"
	"strings"
)

func isLongPressedName(name string, typed string) bool {
	log.Println("isLongPressedName(name", name, ",typed:", typed, ") started")
	a := strings.Split(name, "")
	b := strings.Split(typed, "")
	log.Println("a (size:", len(a), ") -> ", a)
	log.Println("b (size:", len(b), ") -> ", b)
	pre := ""
	retval := true

	for i, c := range b {
		log.Println("[i:", i, "] -> ", c)
		if len(a) > 0 && a[0] == c {
			pre = a[0]
			a = a[1:]
		} else if pre == c {
			/* OK */
		} else {
			retval = false
			break
		}
	}

	return retval
}

func main() {
	log.Println("result -> ", isLongPressedName("aalex", "aaleex"))
	log.Println("result -> ", isLongPressedName("saeed", "ssaaedd"))
	log.Println("result -> ", isLongPressedName("leelee", "llllleeelee"))
	log.Println("result -> ", isLongPressedName("laiden", "laiden"))
	log.Println("result -> ", isLongPressedName("abbbbcd", "aaaaaaaaaaabbbcd"))
	log.Println("result -> ", isLongPressedName("aalex", "aaleexa"))
}
