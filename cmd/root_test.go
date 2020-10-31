package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Sean0628/ctap/test"
	"github.com/spf13/cobra"
)

func getRootCommand() *cobra.Command {
	return New()
}

func TestRootCmd(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "")

	if result.Error != nil {
		t.Error(result.Error)
	}

	if !strings.Contains(result.Output, "Usage:") {
		t.Error("Expected usage message to be printed out.")
		t.Error(result.Output)
	}
}

func TestRootCmd_BoolOptions(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "")

	if result.Error != nil {
		t.Error(result.Error)
	}

	boolArgs := [4][2]bool{
		{fVersion, false},
		{fDayOfWeek, false},
		{fVerbose, false},
		{fFormat24, false},
	}

	if !matchBool(boolArgs) {
		t.Error("Expected options should have been initialized defalut values.")
	}
}

func matchBool(arr [4][2]bool) bool {
	for _, setOfVal := range arr {
		if setOfVal[0] != setOfVal[1] {
			return false
		}
	}

	return true
}

func TestRootCmd_StringOptions(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "")

	if result.Error != nil {
		t.Error(result.Error)
	}

	stringArgs := [4][2]string{
		{fLocale, "en"},
		{fInputFile, ""},
		{fOutputFile, ""},
		{fFormat, ""},
	}

	if !matchString(stringArgs) {
		t.Error("Expected options should have been initialized defalut values.")
	}
}

func matchString(arr [4][2]string) bool {
	for _, setOfVal := range arr {
		if setOfVal[0] != setOfVal[1] {
			return false
		}
	}

	return true
}

func TestRootCmd_Help(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "-h")

	if result.Error != nil {
		t.Error(result.Error)
	}

	if !strings.Contains(result.Output, "ctap is a CLI crontab parser written in Go.") {
		t.Error("Expected help message to be printed out.")
	}
}

func TestRootCmd_HelpLong(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "--help")

	if result.Error != nil {
		t.Error(result.Error)
	}

	if !strings.Contains(result.Output, "ctap is a CLI crontab parser written in Go.") {
		t.Error("Expected help message to be printed out.")
	}
}

func TestRootCmd_Locale(t *testing.T) {
	localeType := "ja"

	cmd := getRootCommand()
	opts := fmt.Sprintf("-l@%s", localeType)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fLocale != localeType {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_LocaleLong(t *testing.T) {
	localeType := "ja"

	cmd := getRootCommand()
	opts := fmt.Sprintf("--locale@%s", localeType)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fLocale != localeType {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_Input(t *testing.T) {
	fileName := test.WriteTmpCrontabFile("")
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("-i@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fInputFile != fileName {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_InputLong(t *testing.T) {
	fileName := test.WriteTmpCrontabFile("")
	defer test.RemoveTmpCrontabFile(fileName)

	cmd := getRootCommand()
	opts := fmt.Sprintf("--input@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fInputFile != fileName {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_Output(t *testing.T) {
	fileName := "tmp/crontab.test.txt"

	cmd := getRootCommand()
	opts := fmt.Sprintf("-o@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fOutputFile != fileName {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_OutputLong(t *testing.T) {
	fileName := "tmp/crontab.test.txt"

	cmd := getRootCommand()
	opts := fmt.Sprintf("--output@%s", fileName)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fOutputFile != fileName {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_Format(t *testing.T) {
	formatType := "csv"

	cmd := getRootCommand()
	opts := fmt.Sprintf("-f@%s", formatType)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fFormat != formatType {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_FormatLong(t *testing.T) {
	formatType := "csv"

	cmd := getRootCommand()
	opts := fmt.Sprintf("--format@%s", formatType)

	result := test.RunCmd(cmd, opts)
	if result.Error != nil {
		t.Error(result.Error)
	}

	if fFormat != formatType {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_Version(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "-V")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fVersion {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_VersionLong(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "--version")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fVersion {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_DowStartsAtOne(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "-d")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fDayOfWeek {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_DowStartsAtOneLong(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "--dow-starts-at-one")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fDayOfWeek {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_Verbose(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "-v")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fVerbose {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_VerboseLong(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "--verbose")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fVerbose {
		t.Error("Expected to be true")
	}
}

func TestRootCmd_24hrsFormat(t *testing.T) {
	cmd := getRootCommand()
	result := test.RunCmd(cmd, "--24-hour")
	if result.Error != nil {
		t.Error(result.Error)
	}

	if !fFormat24 {
		t.Error("Expected to be true")
	}
}
