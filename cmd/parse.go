package cmd

import (
  "fmt"
)

func GetParsedExpressionDisplay(args []string) string {
  expr := args[0]

  loc, err := GetParseLocale()
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  exprDesc, err := GetExprDescriptor(loc)
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  desc, err := exprDesc.ToDescription(args[0], loc)
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  return fmt.Sprintf("%s: %s\n", expr, desc)
}
