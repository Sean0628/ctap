package cmd

import (
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
    {version,   false},
    {dayOfWeek, false},
    {version,   false},
    {format24,  false},
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

  stringArgs := [1][2]string{
    {locale, "en"},
  }

  if !matchString(stringArgs) {
    t.Error("Expected options should have been initialized defalut values.")
  }
}

func matchString(arr [1][2]string) bool {
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
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "-l@ja")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if locale != "ja" {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_LocaleLong(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "--locale@ja")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if locale != "ja" {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_Version(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "-V")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !version {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_VersionLong(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "--version")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !version {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_DowStartsAtOne(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "-d")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !dayOfWeek {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_DowStartsAtOneLong(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "--dow-starts-at-one")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !dayOfWeek {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_Verbose(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "-v")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !verbose {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_VerboseLong(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "--verbose")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !verbose {
    t.Error("Expected to be true")
  }
}

func TestRootCmd_24hrsFormat(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, "--24-hour")
  if result.Error != nil {
    t.Error(result.Error)
  }

  if !format24 {
    t.Error("Expected to be true")
  }
}