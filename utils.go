package loggy

func GetPrefix(text string, color string, labelColor string) string {
	return "\033[" + color + "m\033[1m\033[" + labelColor + "m" + text + "\033[0m"
}
