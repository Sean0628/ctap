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

	fLocale     = "en"
	fInputFile  = ""
	fOutputFile = ""
	fFormat     = ""
	fVersion    = false
	fDayOfWeek  = false
	fVerbose    = false
	fSort       = false
	fFormat24   = false
)
