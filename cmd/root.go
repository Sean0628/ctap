package cmd

import (
  "os"
  "fmt"

  "github.com/spf13/cobra"
)

var (
  rootCmd = &cobra.Command{
    Use:   "ctap",
    Short: "CLI crontab parser",
    Long: `ctap is a CLI crontab parser written in Go.`,
    RunE: func(cmd *cobra.Command, args []string) error {
			if version {
				cmd.Print(GetVersionDisplay())
				return nil
			}
      if len(args) < 1 {
        cmd.Println(cmd.UsageString())
        return nil
      }

      return nil
    },
}

)
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
	rootCmd.Flags().StringVarP(&locale, "locale", "l", "en", "Prints description in the specified locale")
  rootCmd.Flags().BoolVarP(&version, "version", "V", false, "Prints version information")
	rootCmd.Flags().BoolVarP(&dayOfWeek, "dow-starts-at-one", "d", false, "Is day of the week starts at 1 (Monday-Sunday: 1-7)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Prints description in verbose format")
	rootCmd.Flags().BoolVar(&format24, "24-hour", false, "Prints description in 24 hour time format")
}
