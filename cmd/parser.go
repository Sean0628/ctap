package cmd

import (
  "fmt"

  "github.com/lnquy/cron"
)

func GetExprDescriptor(loc cron.LocaleType) (exprDesc *cron.ExpressionDescriptor, err error) {
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

func GetParseLocale() (loc cron.LocaleType, err error) {
  loc, err = cron.ParseLocale(locale)
  if err != nil {
    return loc, fmt.Errorf("ctap: failed to get locale: %w", err)
  }

  return loc, nil
}
