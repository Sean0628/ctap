package cmd

import (
  "bufio"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "strings"

  "github.com/lnquy/cron"
)

func GetParsedExpressionFromFileDisplay(args []string) string {
  filePath := strings.TrimSpace(inputFile)
  outputFilePath := strings.TrimSpace(outputFile)

  loc, err := GetParseLocale()
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  exprDesc, err := GetExprDescriptor(loc)
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
  if err != nil {
    return fmt.Sprintf("ctap: failed to %s.\n", err)
  }

  results, err := stream(exprDesc, loc, bufio.NewReader(f))
  if err != nil {
    return fmt.Sprintf("ctap: unexpected error occured (%s).\n", err)
  }
  if len(results) == 0 {
    return fmt.Sprintln("ctap: noting to be printed out. possibly input file does not have any contents.")
  }

  if len(outputFilePath) > 0 {
    message := []byte(results)

    absolutePath, _ := filepath.Abs(outputFilePath)
    err := ioutil.WriteFile(absolutePath, message, 0644)
    if err != nil {
      return fmt.Sprintln(err.Error())
    }

    return fmt.Sprintf("ctap: the results are printed out to %s.\n", outputFilePath)
  }

  return results
}

func stream(exprDesc *cron.ExpressionDescriptor, localeType cron.LocaleType, reader *bufio.Reader) (results string, err error) {
  var lines []string

  for {
    line, _, err := reader.ReadLine()
    if err != nil && err == io.EOF {
      return strings.Join(lines[:], ""), nil
    }

    expr, remaining := normalize(string(line))

    if len(expr) == 0 {
      continue
    }

    desc, err := exprDesc.ToDescription(expr, localeType)
    if err != nil {
      return results, err
    }

    if len(remaining) > 0 {
      lines = append(lines, fmt.Sprintf("%s: %s | %s\n", expr, desc, remaining))
      continue
    }
    lines = append(lines, fmt.Sprintf("%s: %s\n", expr, desc))
  }
}


func normalize(line string) (expr string, remainder string) {
  if strings.HasPrefix(line, commentOutPrefix) {
    return "", line
  }

  parts := strings.Fields(line)
  if len(parts) < validCronExpressionCount {
    return "", line
  }

  if !validCronExpression.MatchString(line) {
    return strings.Join(parts[:validCronExpressionCount], " "), strings.Join(parts[validCronExpressionCount:], " ")
  }

  return line, ""
}
