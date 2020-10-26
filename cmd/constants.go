package cmd

import (
	"regexp"
)

const (
	validCronExpressionCount = 5
	commentOutPrefix         = "#"
)

var (
	validCronExpression = regexp.MustCompile(`^[wWlL /?,*#\-0-9]*$`)

	locale     = "en"
	inputFile  = ""
	outputFile = ""
	version    = false
	dayOfWeek  = false
	verbose    = false
	format24   = false
)
