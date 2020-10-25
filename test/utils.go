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

  splitFn := func(c rune) bool {
    return c == '@'
  }
  args := strings.FieldsFunc(input, splitFn)

  c.SetArgs(args)
  err := c.Execute()
  output := buf.String()
  return resulter{err, output, c}
}

func AssertResult(t *testing.T, expectedValue interface{}, actualValue interface{}) {
  t.Helper()
  if expectedValue != actualValue {
    t.Error("\nExpected: ", expectedValue, "\n     Got: ", actualValue, fmt.Sprintf("\n    Type:  %T", actualValue))
  }
}
