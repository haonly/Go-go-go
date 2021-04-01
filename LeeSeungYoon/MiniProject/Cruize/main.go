package main

import (
	"algorithm/MiniProject/Cruize/Cruize"
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	TERMINATE = "end"
)

func PrintCtrlMenu() {
	fmt.Println("==== Enter Control Command ====")
	fmt.Println(`
	0 : UPH (Up Hold)
	1 : UPS (Up Short)
	2 : DNH (Down Hold)
	3 : DNS (Down Short)
	end : driving end`)
	fmt.Print("> ")
}

func Drive() {
	log.Println("Drive start...")
	defer func() {
		log.Println("Drive Func end...")
		return
	}()

	var currentSpeed int = 0

	// keyboard input reader
	ioReader := bufio.NewReader(os.Stdin)

	for {
		PrintCtrlMenu() // print control menu

		// get keyboard input
		command, _ := ioReader.ReadString('\n')
		log.Printf("drive input command : %v\n", command)

		// 개행문자 \n 제거
		command = command[:len(command)-2]

		// 종료 조건
		if command == TERMINATE {
			log.Println("drive terminated....")
			break
		}

		// command -> speed 변경
		Cruize.CommandToSpeed(&command, &currentSpeed)
	}

}

func main() {
	Drive()
	return
}

// "crz start" 라는 입력이 들어오면, 크루즈 모드 시작
// "crz off" 입력 들어오면, 크루즈 모드 종료
