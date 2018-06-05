package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)
var verOnly bool

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "version description",

  Run: func(cmd *cobra.Command, args []string) {
    if quiet{
      fmt.Print(VERSION)
    } else {
      fmt.Println(RootCmd.Use + " " + VERSION)
    }
  },
}

func init() {
  RootCmd.AddCommand(versionCmd)
  versionCmd.Flags().BoolVarP(&verOnly, "quiet", "q", false, "Show only the version")
}
