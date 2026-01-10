package main

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/Goku-kun/fileorg/internal/organizer"
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

	postitionalArgs := flag.Args()

	if len(postitionalArgs) == 0 {
		fmt.Fprintf(os.Stderr, "No valid directory path supplied. Exiting...")
		os.Exit(1)
	}

	isDryRun := dryRun || n
	isVerbose := verbose || v

	validByOptions := []string{"extension", "date", "size"}
	isValidBy := slices.Contains(validByOptions, by)

	if !isValidBy {
		fmt.Fprintf(os.Stderr, "Invalid value for --by. Valid options are 'extension', 'date', or 'size'.")
		os.Exit(1)
		return
	}

	if isDryRun {
		println("Dry run mode enabled. No changes will be made.")
	}

	var strategy organizer.Strategy
	switch by {
	case "extension":
		strategy = &organizer.ExtensionStrategy{}
	default:
		fmt.Fprintf(os.Stderr, "Strategy not found. Exiting...")
		os.Exit(1)
		return
	}

	cfg := organizer.Config{
		SourceDir: postitionalArgs[0],
		Strategy:  strategy,
		DryRun:    isDryRun,
		Verbose:   isVerbose,
	}

	fmt.Printf("Config: %+v\n", cfg)

	organizer.NewOrganizer(cfg)
}
