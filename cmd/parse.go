package cmd

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func GetParsedExpressionDisplay(args []string) string {
  expr := args[0]
  outputFilePath := strings.TrimSpace(outputFile)

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

  results := fmt.Sprintf("%s: %s\n", expr, desc)

  if len(outputFilePath) > 0 {
    // TODO: extract this procedure to another function
    message := []byte(results)
    err := ioutil.WriteFile(outputFilePath, message, 0644)
    if err != nil {
      return fmt.Sprintln(err.Error())
    }

    return fmt.Sprintf("ctap: the results are printed out to %s.\n", outputFilePath)
  }

  return results
}
