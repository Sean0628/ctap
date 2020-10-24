package cmd

import (
  "fmt"

  "github.com/lnquy/cron"
)

func GetParsedExpressionDisplay(args []string) string {
  expr := args[0]

  loc, err := parseLocale()
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  exprDesc, err := getExprDescriptor(loc)
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  desc, err := exprDesc.ToDescription(args[0], loc)
  if err != nil {
    return fmt.Sprintln(err.Error())
  }

  return fmt.Sprintf("%s: %s\n", expr, desc)
}

func getExprDescriptor(loc cron.LocaleType) (exprDesc *cron.ExpressionDescriptor, err error) {
  exprDesc, err = cron.NewDescriptor(
      cron.Use24HourTimeFormat(format24),
      cron.DayOfWeekStartsAtOne(dayOfWeek),
      cron.Verbose(verbose),
      cron.SetLocales(loc),
  )

  if err != nil {
    return nil, fmt.Errorf("ctap: failed to initialize cron expression descriptor: %s", err)
  }

  return exprDesc, nil
}

func parseLocale() (loc cron.LocaleType, err error) {
  loc, err = cron.ParseLocale(locale)
  if err != nil {
    return loc, fmt.Errorf("ctap: failed to get locale: %w", err)
  }

  return loc, nil
}
