package cmd

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strings"

  "github.com/lnquy/cron"
)

func GetParsedExpressionFromFileDisplay(args []string) string {
  filePath := strings.TrimSpace(file)

  loc, err := GetParseLocale()
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  exprDesc, err := GetExprDescriptor(loc)
  if len(filePath) > 0 {
    f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
    if err != nil {
      fmt.Printf("failed to open file: %s", err)
    }

    if err := stream(exprDesc, loc, bufio.NewReader(f)); err != nil {
      fmt.Printf("error: %s", err)
    }

    return ""
  }

  return ""
}

func stream(exprDesc *cron.ExpressionDescriptor, localeType cron.LocaleType, reader *bufio.Reader) error {
  for {
    line, _, err := reader.ReadLine()
    if err != nil && err == io.EOF {
      return nil
    }

    expr, remaining := normalize(string(line))

    if len(expr) == 0 {
      continue
    }

    desc, err := exprDesc.ToDescription(expr, localeType)
    if err != nil {
      fmt.Printf("error: %s\n", err)
      continue
    }

    if len(remaining) > 0 {
      fmt.Printf("%s: %s | %s\n", expr, desc, remaining)
      continue
    }
    fmt.Printf("%s: %s\n", expr, desc)
  }
}


func normalize(line string) (expr string, remainder string) {
  if strings.HasPrefix(line, "#") {
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
