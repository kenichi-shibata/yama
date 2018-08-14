package cmd

import (
  "fmt"
  "log"
  "io/ioutil"
  "github.com/spf13/cobra"
  "strings"
)

var (
  merger string
  mergee string
)
var mergeCmd = &cobra.Command{
  Use:   "merge",
  Short: "A brief description of your command",
  Args: cobra.MinimumNArgs(2),
  Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Run: func(cmd *cobra.Command, args []string) {
    // fmt.Println("Print: " + strings.Join(args, "--"))
    merge(args)
  },
}

func merge(f... []string) {
  mergee, merger = f[0], f[1]
  mergeeFile, err := ioutil.ReadFile(mergee)
  
  if err != nil {
    log.Fatal("[ERROR] err reading: \n", err)
  }

  mergerFile, err := ioutil.ReadFile(merger)

  if err != nil {
    log.Fatal("[ERROR] err reading: \n", err)
  }

  log.Printf("[VERBOSE] Input: \n %s \n\n\n", mergeeFile)
  log.Printf("[VERBOSE] Input: \n %s", mergerFile)

  // yamlFile := mergeFiles(mergeeFile, mergerFile)
  // #for key, value := fd


}

// func mergeFiles(interface{}, ...interface{}) interface{} {
//   return
// }

func init() {
  RootCmd.AddCommand(mergeCmd)
}
