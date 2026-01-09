package main

import (
	"flag"
	"fmt"
	"slices"
)

func main() {
	var dryRun, n, verbose, v bool
	var by string

	flag.StringVar(&by, "by", "extension", "Criteria for sorting files. Valid options are 'extension', 'date', or 'size'")

	flag.BoolVar(&dryRun, "dry-run", false, "Perform a trial run with no changes made")
	flag.BoolVar(&n, "n", false, "Perform a trial run with no changes made (shorthand for --dry-run)")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&v, "v", false, "Enable verbose output (shorthand for --verbose)")

	flag.Parse()

	dirPath := flag.Args()
	fmt.Println("Directory path:", dirPath)

	isDryRun := dryRun || n
	isVerbose := verbose || v

	if isVerbose {
		println("Verbose mode enabled.")
	}

	println("Sorting files by:", by)

	validByOptions := []string{"extension", "date", "size"}
	isValidBy := slices.Contains(validByOptions, by)

	if !isValidBy {
		fmt.Println("Invalid value for --by. Valid options are 'extension', 'date', or 'size'.")
		return
	}

	if isDryRun {
		println("Dry run mode enabled. No changes will be made.")
	} else {
		println("Executing with changes.")
	}
}
