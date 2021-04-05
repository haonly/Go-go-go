package Cruize

const (
	UPH = 1 * iota // Up Hold - 위로 길게 버튼 누름
	UPS            // Up Short - 위로 짧게 버튼 누름
	DNH            // Down Hold - 아래로 길게 버튼 누름
	DNS            // Down Short - 아래로 짧게 버튼 누름

	SPEED_BUFFER = 1000 // speed buffer
)

type ButtonType int
