package utils

import (
	"strings"
)

func GetIndex(text []string, colorText string) [][]bool {
	newTxt := text
	newColorTxt := strings.Split(colorText, "\\n")
	colored2Array := make([][]bool, len(newTxt))
	for i := range newTxt {
		colored2Array[i] = make([]bool, len(newTxt[i]))
	}
	if colorText == "" {
		return colored2Array
	}
	for line := range newTxt {
		var start, end, index int
		for {
			if len(newColorTxt) == len(newTxt) {
				index = strings.Index(newTxt[line][start:], newColorTxt[line])
			} else {
				index = strings.Index(newTxt[line][start:], colorText)
			}
			if index == -1 {
				break
			}
			index += start
			if len(newColorTxt) == len(newTxt) {
				end = index + len(newColorTxt[line])
			} else {
				end = index + len(colorText)
			}
			for i := index; i < end; i++ {
				colored2Array[line][i] = true
			}
			start = end
		}
	}
	return colored2Array
}
