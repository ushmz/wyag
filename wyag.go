package main

import (
	"flag"
	"fmt"
	"os"
	"wyag/cmd"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
	checkoutCmd := flag.NewFlagSet("checkout", flag.ExitOnError)
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
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

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		cmd.AddCmd(addCmd.Args())
	case "cat-file":
		catFileCmd.Parse(os.Args[2:])
		cmd.CatFileCmd()
	case "checkout":
		checkoutCmd.Parse(os.Args[2:])
		cmd.CheckoutCmd()
	case "commit":
		commitCmd.Parse(os.Args[2:])
		cmd.CommitCmd()
	case "hash-object":
		hashObjectCmd.Parse(os.Args[2:])
		cmd.HashObjectCmd()
	case "init":
		initCmd.Parse(os.Args[2:])
		cmd.InitCmd()
	case "log":
		logCmd.Parse(os.Args[2:])
		cmd.LogCmd()
	case "ls-tree":
		lsTreeCmd.Parse(os.Args[2:])
		cmd.LsTreeCmd()
	case "merge":
		mergeCmd.Parse(os.Args[2:])
		cmd.MergeCmd()
	case "rebase":
		rebaseCmd.Parse(os.Args[2:])
		cmd.RebaseCmd()
	case "rev-parse":
		revParseCmd.Parse(os.Args[2:])
		cmd.RebaseCmd()
	case "rm":
		rmCmd.Parse(os.Args[2:])
		cmd.RmCmd()
	case "show-ref":
		showRefCmd.Parse(os.Args[2:])
		cmd.ShowRefCmd()
	case "tag":
		tagCmd.Parse(os.Args[2:])
		cmd.TagCmd()
	default:
		fmt.Println("Invalid sub-commands")
		fmt.Println("[Insert help message here]")
		os.Exit(1)
	}
}
