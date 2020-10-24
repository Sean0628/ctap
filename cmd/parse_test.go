package cmd

import (
  "fmt"
  "strings"
  "testing"

  "github.com/Sean0628/ctap/test"
)

const cronExpression = "0 15 * * 1-5"

func TestGetParsedExpressionDisplay(t *testing.T) {
  result := test.RunCmd(rootCmd, fmt.Sprintln(cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 15 * * 1-5: At 03:00 PM, Monday through Friday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayInSpecifiedLocale(t *testing.T) {
  result := test.RunCmd(rootCmd, fmt.Sprintf("-l@ja@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 15 * * 1-5: 次において実施03:00 PM、月曜日 から 金曜日 まで"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayInSpecifiedLocaleLong(t *testing.T) {
  result := test.RunCmd(rootCmd, fmt.Sprintf("--locale@ja@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 15 * * 1-5: 次において実施03:00 PM、月曜日 から 金曜日 まで"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayIn24TimeFormat(t *testing.T) {
  result := test.RunCmd(rootCmd, fmt.Sprintf("--24-hour@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 15 * * 1-5: At 15:00, Monday through Friday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayStartingAt1(t *testing.T) {
  result := test.RunCmd(rootCmd, fmt.Sprintf("-d@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 15 * * 1-5: At 03:00 PM, Monday through Friday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}

func TestGetParsedExpressionDisplayStartingAt1Long(t *testing.T) {
  result := test.RunCmd(rootCmd, fmt.Sprintf("--dow-starts-at-one@%s", cronExpression))
  if result.Error != nil {
    t.Error(result.Error)
  }
  expectedOutput := "0 15 * * 1-5: At 03:00 PM, Sunday through Thursday"
  test.AssertResult(t, expectedOutput, strings.Trim(result.Output, "\n "))
}
