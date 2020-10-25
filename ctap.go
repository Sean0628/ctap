package main

import (
  "os"
  "fmt"

  command "github.com/Sean0628/ctap/cmd"
)

func main() {
  cmd := command.New()
  if err := cmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
