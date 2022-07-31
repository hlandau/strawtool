package cli

import (
	"fmt"
	"github.com/hlandau/modtest"
)

func Main() {
	fmt.Printf("This is the CLI: %v\n", modtest.Value())
}
