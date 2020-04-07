package greeting

import "strings"

func HelloWorld() string {
	return GreetMe("World")
}

func GreetMe(s string) string {
	var str strings.Builder
	str.WriteString("Hello ")
	str.WriteString(s)
	str.WriteString("!")
	return str.String()
}
