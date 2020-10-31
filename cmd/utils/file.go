package cmd

import (
	"fmt"
	"strings"
)

func GetFomattedLines(lines [][3]string) (results []string) {
	for _, lineSet := range lines {
		formattedLine := strings.Join(lineSet[:], " | ")
		results = append(results, formattedLine)
	}

	return results
}

func GetMdFormattedLines(lines [][3]string) (results []string) {
	results = append(results, "| Original | Translated | Command |")
	results = append(results, "|---|---|---|")

	for _, lineSet := range lines {
		formattedLine := fmt.Sprintf("| %s |", strings.Join(lineSet[:], " | "))
		results = append(results, formattedLine)
	}

	return results
}

func GetCsvFormattedLines(lines [][3]string) (results []string) {
	results = append(results, "Original, Translated, Command")

	for _, lineSet := range lines {
		lineSet = [3]string{fmt.Sprintf("\"%s\"", lineSet[0]), fmt.Sprintf("\"%s\"", lineSet[1]), lineSet[2]}
		formattedLine := strings.Join(lineSet[:], ", ")
		results = append(results, formattedLine)
	}

	return results
}
