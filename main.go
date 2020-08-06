package main

import (
	"os"

	"github.com/netauth/pam-helper/internal/module"
)

func main() {
	os.Exit(module.Exec())
}
