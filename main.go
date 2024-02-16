package loggy

import (
	"fmt"
)

func Info(v ...interface{}) {
	fmt.Print(GetPrefix(" INFO ", "44", "30"), " ")
	fmt.Println(v...)
}

func Error(v ...interface{}) {
	fmt.Print(GetPrefix(" ERRO ", "41", "30"), " ")
	fmt.Println(v...)
}

func Warn(v ...interface{}) {
	fmt.Print(GetPrefix(" WARN ", "43", "30"), " ")
	fmt.Println(v...)
}

func Success(v ...interface{}) {
	fmt.Print(GetPrefix(" DONE ", "42", "30"), " ")
	fmt.Println(v...)
}

func Box(v string) {
	border := "──"

	for i := 0; i < len(v); i++ {
		border += "─"
	}

	fmt.Println("╭" + border + "╮")
	fmt.Println("│", v, "│")
	fmt.Println("╰" + border + "╯")
}
