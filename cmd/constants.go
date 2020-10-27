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

	validFormatTypes = []string{formatCsv, formatMd}
	formatCsv        = "csv"
	formatMd         = "markdown"

	locale     = "en"
	inputFile  = ""
	outputFile = ""
	format     = ""
	version    = false
	dayOfWeek  = false
	verbose    = false
	format24   = false
)
