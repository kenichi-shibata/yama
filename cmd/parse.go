package cmd

import (
  "fmt"
  "log"
  "os"
  "io/ioutil"
  "encoding/json"
  "gopkg.in/yaml.v2"
  "github.com/spf13/cobra"
  "github.com/hashicorp/logutils"
)

var parseCmd = &cobra.Command{
  Use:   "parse",
  Short: "Parses and checks the yaml file",
  Long: `Parsing should return errors if your yaml file is invalid.`,
  Run: func(cmd *cobra.Command, args []string) {
    setLog()
    if infile != "" {
      log.Print("[DEBUG] Parsing...")
      run()
    } else {
      log.Print("Could not find the file --infile <filename>")
    }
  },
}

var ( 
  verbose bool
  debug bool
  quiet bool
  infile string
)

func run() {
  
  data, err := ioutil.ReadFile(infile)
  
  if err != nil {
    log.Fatalf("[ERROR] err reading: \n",err)
  }

  log.Printf("[VERBOSE] Input: \n %s", data)
  var yamldata interface{}
  
  if err := yaml.Unmarshal(data, &yamldata); err != nil {
    log.Fatalf("[ERROR] err unmarshalling: %v\n", err)
    return
  }

  yamldata = convert(yamldata)

  if b, err := json.Marshal(yamldata); err != nil {
      panic(err)
    } else {
      log.Printf("[VERBOSE] Output: %s\n", b)
      log.Printf("[DEBUG] %s: %s", Typeof(yamldata), yamldata)
      fmt.Printf("%s is a valid yaml file", infile)
    }
}

// Takes either a  map[string]interface{} or a map[interface{}]interface{}} recursively converts 
func convert(input interface{}) interface{} {
  switch x := input.(type) {

    case map[interface{}]interface{}:
      mappedData := map[string]interface{}{}
      for key, value := range x {
        mappedData[key.(string)] = convert(value)
      }
      return mappedData

    case  []interface{}:
      for input, value := range x {
        x[input] = convert(value)
      }
  }
  return input
}

func init() {
  RootCmd.AddCommand(parseCmd)
  parseCmd.Flags().StringVarP(&infile, "infile", "i", "", "Yaml file to be parsed")
  parseCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "checks only without printing anything valid exit code 0 or print the error")
  parseCmd.Flags().BoolVarP(&debug, "debug", "g", false, "Debug Output")
  parseCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Verbose Output")
}

