package cmd

import (
  "os"
  "fmt"

  "github.com/spf13/cobra"
)

var (
  version bool

  rootCmd = &cobra.Command{
    Use:   "ctap",
    Short: "CLI crontab parser",
    Long: `ctap is a CLI crontab parser written in Go.`,
    RunE: func(cmd *cobra.Command, args []string) error {
			if version {
				cmd.Print(GetVersionDisplay())
				return nil
			}
      cmd.Println(cmd.UsageString())

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
  rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Prints version information")
}
