package Cruize

import (
	"log"
	"strconv"
)

func MakeChannel(command *string) (chan int, error) {
	ch := make(chan int, len(*command))
	return ch, nil
}

func CommandToSpeed(command *string, currentSpeed *int) {
	// make speed queue to save speed inputs
	speedQueue, err := MakeChannel(command)
	if err != nil {
		log.Fatal(err)
	}

	// parse keyboard only [1, 2, 3, 4]
	parsedCommand := parseCommand(command)

	// insert commands into speed queue
	insertCommandQueue(speedQueue, &parsedCommand)

	close(speedQueue)

	for {
		if val, ok := <-speedQueue; ok {
			CalculateCruiseControlSpeed(ButtonType(val), currentSpeed)
		} else {
			break
		}
	}
}

func CalculateCruiseControlSpeed(button ButtonType, currentSpeed *int) {
	defer func() {
		log.Printf("button : %v, current speed : %v\n", button, *currentSpeed)
	}()
	speedUpDown := 0
	switch button {
	case UPH:
		speedUpDown = 10
	case UPS:
		speedUpDown = 1
	case DNH:
		speedUpDown = -10
	case DNS:
		speedUpDown = -1
	}
	*currentSpeed += speedUpDown
}

func parseCommand(command *string) string {
	var parsed string = ""

	for _, s := range *command {
		value, err := strconv.Atoi(string(s))

		// s가 숫자인 경우
		if err == nil {
			// value가 UPH, UPS, DNH, DNS 중 하나인 경우
			if value == UPH || value == UPS || value == DNH || value == DNS {
				parsed += strconv.Itoa(value)
			}
		}
	}
	return parsed
}

func insertCommandQueue(ch chan int, parsedCommand *string) {
	// channel에 parsed command 넣기
	for _, s := range *parsedCommand {
		v, _ := strconv.Atoi(string(s))
		ch <- v
	}
}
