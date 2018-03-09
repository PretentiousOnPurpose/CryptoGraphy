package main

import (
	"fmt"
	"strings"
)

func indexAdjust(v int32) int32 {
	if v > 122 {
		v = 96 + v - 122
	} else if v < 97 {
		v = 123 - 97 + v
	}
	return v
}

func shiftCipher(shift int32, str string) string {
	v := []rune(strings.ToLower(str))
	for i := 0; i < len(v); i++ {
		v[i] = (v[i] - 96 + shift) % 26 + 96
		v[i] = indexAdjust(v[i])
	}
	return string(v)
}

func shiftDeCipher(shift int32, str string) string {
	v := []rune(strings.ToLower(str))
	for i := 0; i < len(v); i++ {
		v[i] = (v[i] - 96 - shift) % 26 + 96
		v[i] = indexAdjust(v[i])
	}
	return string(v)

}

func main() {
	fmt.Println(shiftCipher(5, "Hello"))
}
