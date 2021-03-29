/*
함수 pseudo code 는 대충

CalculateCruiseControlSpeed(button ButtonType , currentSpeed int) {

return int

}


ButtonType 은
UPH (Up Hold - 위로 길게 버튼 누름)
UPS (Up Short - 위로 짧게 버튼 누름)
DNH (Down Hold - 아래로 길게 버튼 누름)
DNS (Down Short - 아래로 짧게 버튼 누름)

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type ButtonType int
const (
	UPH =  iota + 1
	UPS
	DNH
	DNS
)

func calculateCruiseControlSpeed(button ButtonType, currentSpeed int) int {
	log.Println("calculateCruiseControlSpeed button:", button, ", currentSpeed:", currentSpeed)
	delta := 0
	switch button {
		case UPH:
			log.Println("UPH detected")
			delta = 10 - (currentSpeed % 10)
		case UPS:
			log.Println("UPS detected")
			delta = 1 
		case DNH:
			log.Println("DNH detected")
			mod := currentSpeed % 10
			if(mod == 0){
				delta = -10
			}else {
				delta = -1 * (currentSpeed % 10)
			}
		case DNS:
			log.Println("DNS detected")
			delta = -1
		default :
			log.Println("default detected")		
	}
	r := currentSpeed + delta 
	if (r <= 0) {
		r = 1
	}else if (r > 300){
		r = 300
	}
	return r
}

func main(){
	currentSpeed := rand.Intn(100 - 10) + 10
	
	reader := bufio.NewReader(os.Stdin)
  	fmt.Println("---------------------")

	for {
		fmt.Println("Current Speed : ", currentSpeed)
		fmt.Println("<< Press number key >>")		
		fmt.Println("1 : Up Hold")
		fmt.Println("2 : Up Short")
		fmt.Println("3 : Down Hold")
		fmt.Println("4 : Down Short")
		fmt.Print(">> ")

		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if(text==""){
			continue;
		}else if(text=="q" || text=="exit"){
			fmt.Println("Bye")
			break
		}else{
			a, err := strconv.Atoi(text)
			if(err==nil){
				log.Println("a:", a)
				currentSpeed = calculateCruiseControlSpeed(ButtonType(a), currentSpeed)
				log.Println("currentSpeed:", currentSpeed)
				continue;
			}
			fmt.Println("Sorry. Invalid input(", text, "). try again. (Exit : q or exit)")
			fmt.Println("")
		}	
	}
}