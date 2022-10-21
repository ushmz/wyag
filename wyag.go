package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	initCmd.Usage = printInitUsage

	logCmd := flag.NewFlagSet("log", flag.ExitOnError)
	lsTreeCmd := flag.NewFlagSet("ls-tree", flag.ExitOnError)
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	rebaseCmd := flag.NewFlagSet("rebase", flag.ExitOnError)
	revParseCmd := flag.NewFlagSet("rev-parse", flag.ExitOnError)
	rmCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	showRefCmd := flag.NewFlagSet("show-ref", flag.ExitOnError)
	tagCmd := flag.NewFlagSet("tag", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Expected sub-commands")
		fmt.Println("[Insert help message here]")
		os.Exit(1)
	}

	flag.Parse()

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		AddCmd(addCmd.Args())
	case "cat-file":
		catFileCmd.Parse(os.Args[2:])
		CatFileCmd()
	case "checkout":
		checkoutCmd.Parse(os.Args[2:])
		CheckoutCmd()
	case "commit":
		commitCmd.Parse(os.Args[2:])
		CommitCmd()
	case "hash-object":
		hashObjectCmd.Parse(os.Args[2:])
		HashObjectCmd()
	case "init":
		initCmd.Parse(os.Args[2:])
		InitCmd(flag.Arg(len(flag.Args())))
	case "log":
		logCmd.Parse(os.Args[2:])
		LogCmd()
	case "ls-tree":
		lsTreeCmd.Parse(os.Args[2:])
		LsTreeCmd()
	case "merge":
		mergeCmd.Parse(os.Args[2:])
		MergeCmd()
	case "rebase":
		rebaseCmd.Parse(os.Args[2:])
		RebaseCmd()
	case "rev-parse":
		revParseCmd.Parse(os.Args[2:])
		RebaseCmd()
	case "rm":
		rmCmd.Parse(os.Args[2:])
		RmCmd()
	case "show-ref":
		showRefCmd.Parse(os.Args[2:])
		ShowRefCmd()
	case "tag":
		tagCmd.Parse(os.Args[2:])
		TagCmd()
	default:
		fmt.Println("Invalid sub-commands")
		fmt.Println("[Insert help message here]")
		os.Exit(1)
	}
}
