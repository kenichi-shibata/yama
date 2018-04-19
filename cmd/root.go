package cmd

import (
  "fmt"
  "log"
  "os"
  "github.com/spf13/cobra"
  "github.com/hashicorp/logutils"
)

var (
  VERSION string
  cfgFile string
)

var RootCmd = &cobra.Command{
  Use:   "yama",
  Short: "Yet Another yaMl pArser",
  Long: `Allows you to parse and merge yaml files

  山 (Yama or san) is the japanese word for mountain

  富士山 (Mount Fuji)


  Kind of useful for kubernetes specs? 

  Also add a schema and try to parse?
  `,
}
  
func Execute(version string) {
  VERSION = version
  if err := RootCmd.Execute(); err != nil {
          fmt.Println(err)
          os.Exit(1)
  }
}

// Takes either a  map[string]interface{} or a map[interface{}]interface{}} recursively converts 
func Convert(input interface{}) interface{} {
  switch x := input.(type) {

    case map[interface{}]interface{}:
      mappedData := map[string]interface{}{}
      for key, value := range x {
        mappedData[key.(string)] = Convert(value)
      }
      return mappedData

    case  []interface{}:
      for input, value := range x {
        x[input] = Convert(value)
      }
  }
  return input
}

func Typeof(variable interface{}) string {
  return fmt.Sprintf("%T", variable)
}

func SetLog() {
  filter := &logutils.LevelFilter{
    Levels: []logutils.LogLevel{"DEBUG","VERBOSE", "INFO", "ERROR"},
    MinLevel: logutils.LogLevel(GetLevel()),
    Writer: os.Stderr,
  }
  log.SetOutput(filter)
}

func GetLevel() string {
  if debug {
    return "DEBUG"
  } else if verbose {
    return "VERBOSE"
  } else if quiet {
    return "ERROR"
  } else {
    return "INFO"
  }
}

func init() {
  cobra.OnInitialize()
  RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
