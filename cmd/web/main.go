package main

import (
	"fmt"
	"os"

	"github.com/iamanders/subscriptions/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
