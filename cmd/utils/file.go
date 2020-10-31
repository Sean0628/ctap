package cmd

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
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

type TimeSlice []time.Time

func (p TimeSlice) Len() int {
	return len(p)
}

func (p TimeSlice) Less(i, j int) bool {
	return p[i].Before(p[j])
}

func (l TimeSlice) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

const timeLayout = "15:04"

func SortByTime(lines [][3]string) (sortedLines [][3]string) {
	var validTimeExpression = regexp.MustCompile(`([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]|[0-9])`)
	var keys []time.Time
	lineMap := make(map[string][3]string)

	for _, lineSet := range lines {
		matched := validTimeExpression.FindStringSubmatch(lineSet[1])

		if matched != nil {
			t, _ := time.Parse(timeLayout, matched[0])
			keys = append(keys, t)
			lineMap[t.String()] = lineSet
		} else {
			sortedLines = append(sortedLines, lineSet)
		}
	}

	sort.Sort(TimeSlice(keys))

	for _, k := range keys {
		sortedLines = append(sortedLines, lineMap[k.String()])
	}

	return sortedLines
}
