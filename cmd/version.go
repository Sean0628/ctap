package cmd

import (
  "fmt"
)

var Version = "0.0.1"

const ProductName = "ctap"

func GetVersionDisplay() string {
  return fmt.Sprintf("%s version %s\n", ProductName, Version)
}

