package cmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "ctap",
		Short: "CLI crontab parser",
		Long:  `ctap is a CLI crontab parser written in Go.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if fVersion {
				cmd.Print(GetVersionDisplay())
				return nil
			}

			if len(args) > 0 {
				cmd.Print(GetParsedExpressionDisplay(args))
				return nil
			}

			if len(fInputFile) > 0 {
				cmd.Print(GetParsedExpressionFromFileDisplay(args))
				return nil
			}

			cmd.Println(cmd.UsageString())
			return nil
		},

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetOut(cmd.OutOrStdout())
		},
	}

	rootCmd.Flags().StringVarP(&fLocale, "locale", "l", "en", "Prints out description in the specified locale.")
	rootCmd.Flags().StringVarP(&fInputFile, "input", "i", "", "Path to a crontab file to be read from.")
	rootCmd.Flags().StringVarP(&fOutputFile, "output", "o", "", "Path to an output file. If this option is not set, the results are printed out to standard output.")
	rootCmd.Flags().StringVarP(&fFormat, "format", "f", "", "Prints out in the specified format. options: csv, markdown")
	rootCmd.Flags().BoolVarP(&fVersion, "version", "V", false, "Prints out version information.")
	rootCmd.Flags().BoolVarP(&fDayOfWeek, "dow-starts-at-one", "d", false, "Is day of the week starts at 1 (Monday-Sunday: 1-7).")
	rootCmd.Flags().BoolVarP(&fVerbose, "verbose", "v", false, "Prints out description in verbose format.")
	rootCmd.Flags().BoolVarP(&fSort, "sort", "s", false, "Prints out cron expressions in ascending order.")
	rootCmd.Flags().BoolVar(&fFormat24, "24-hour", false, "Prints out description in 24 hour time format.")

	return rootCmd
}
