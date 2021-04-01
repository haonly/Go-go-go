package main


import (
	"fmt"
	"math/rand"
	"time"
)


func calculateCruiseControlSpeed(button string, currentSpeed int) int {
	fmt.Println("CruiseControl Button pressed: ", button, "Current Speed: ", currentSpeed)
	val := 0
	if button == "UPH"{
		val = 5
	}else if button == "UPS"{
		val = 1
	}else if button == "DNH"{
		val = -5
	}else {
		val = -1
	}
	newSpeed := currentSpeed + val
	return newSpeed
}

func main() {
	rand.Seed(time.Now().UnixNano())
	currentSpeed := 10 + rand.Intn(120-10)
	var holdType int
	typeMap := map[int]string{
		1: "UPH",
		2: "UPS",
		3: "DNH",
		4: "DNS",
	}
	
	for {
		fmt.Println("Current Speed: ", currentSpeed)
		fmt.Println("=======================\nChoose a number(0 to exit):")
		fmt.Println("1. Up Hold\n2. Up Short\n3. Down Hold\n4. Down Short ")
		fmt.Print("Input: ")
		
		fmt.Scanln(&holdType)
		if holdType == 0{
			break
		}else{
			buttonTypeVal := typeMap[holdType]
			if buttonTypeVal != ""{
				currentSpeed = calculateCruiseControlSpeed(typeMap[holdType], currentSpeed)
				fmt.Println("Current Speed: ", currentSpeed)
				continue;
			}else{
				fmt.Println("Invalid Input!!\n")
			}
		}
	}
	
}