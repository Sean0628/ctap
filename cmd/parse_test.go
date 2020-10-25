package cmd

import (
  "fmt"
  "strings"
  "testing"

  "github.com/Sean0628/ctap/test"
  "github.com/spf13/cobra"
)

const cronExpression = "0 05 * * 1-5"


func getRootCommand() *cobra.Command {
  return New()
}

func TestGetParsedExpressionDisplay(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, fmt.Sprint(cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 05 * * 1-5: At 05:00 AM, Monday through Friday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayInSpecifiedLocale(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, fmt.Sprintf("-l@ja@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 05 * * 1-5: 次において実施05:00 AM、月曜日 から 金曜日 まで"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayInSpecifiedLocaleLong(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, fmt.Sprintf("--locale@ja@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 05 * * 1-5: 次において実施05:00 AM、月曜日 から 金曜日 まで"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayIn24TimeFormat(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, fmt.Sprintf("--24-hour@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 05 * * 1-5: At 05:00, Monday through Friday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayStartingAt1(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, fmt.Sprintf("-d@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 05 * * 1-5: At 05:00 AM, Sunday through Thursday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayStartingAt1Long(t *testing.T) {
  cmd := getRootCommand()
  result := test.RunCmd(cmd, fmt.Sprintf("--dow-starts-at-one@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 05 * * 1-5: At 05:00 AM, Sunday through Thursday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}
