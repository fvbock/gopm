package main

import (
	gopm "github.com/fvbock/gopm/app"
)

var (
	gpm *gopm.GoPM
)

func init() {
	gpm = gopm.NewGoPMApp()
}

func main() {
	entries := gpm.ScanFile("/home/morpheus/Documents/hm.txt")
	gpm.ShowEntries(entries)

	gpm.Run()
}
