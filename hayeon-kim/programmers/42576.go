package main

import (
	"fmt"
	"sort"
)

func solution(participant []string, completion []string) string {
	sort.Strings(participant)
	sort.Strings(completion)
	
	for i:=0; i < len(completion); i++{
		if(participant[i] != completion[i]){
			return participant[i]
		}
	}
	return participant[len(participant)-1]
}

func main() {
	participant := []string{"leo", "kiki", "eden"}
	completion := []string{"eden", "kiki"}
	ret := solution(participant, completion)
	fmt.Println(ret)
}