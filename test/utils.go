package test

import (
  "bytes"
  "fmt"
  "strings"
  "testing"
  "github.com/spf13/cobra"
)

type resulter struct {
  Error error
  Output string
  Command *cobra.Command
}

func RunCmd(c *cobra.Command, input string) resulter {
  buf := new(bytes.Buffer)
  c.SetOutput(buf)
  c.SetArgs(strings.Split(input, "@"))
  err := c.Execute()
  output := buf.String()
  return resulter{err, output, c}
}

func AssertResult(t *testing.T, expectedValue interface{}, actualValue interface{}) {
  t.Helper()
  if expectedValue != actualValue {
    t.Error("Expected <", expectedValue, "> but got <", actualValue, ">", fmt.Sprintf("%T", actualValue))
  }
}

