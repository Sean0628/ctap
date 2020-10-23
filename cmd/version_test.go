package cmd

import "testing"

func TestGetVersionDisplay(t *testing.T) {
  var expectedVersion = ProductName + " version " + Version
  expectedVersion = expectedVersion + "\n"
  tests := []struct {
    name string
    want string
  }{
    {
      name: "Display Version",
      want: expectedVersion,
    },
  }
  for _, tt := range tests {
    if got := GetVersionDisplay(); got != tt.want {
      t.Errorf("%q. GetVersionDisplay() = %v, want %v", tt.name, got, tt.want)
    }
  }
}
