package main

import "os"

func commandExit(*conf) error {
	os.Exit(0)
	return nil
}
