package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/fatih/color"
)

const (
	ExitCodeOK = iota
	ExitCodeError
	ExitCodeParseFlagsError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	var (
		days    int
		through bool
		version bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprint(cli.errStream, usage)
	}

	flags.IntVar(&days, "days", 5, "")
	flags.IntVar(&days, "d", 5, "")

	flags.BoolVar(&through, "through", false, "")
	flags.BoolVar(&through, "t", false, "")

	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if version {
		fmt.Fprint(cli.outStream, OutputVersion())
		return ExitCodeOK
	}

	lb := LBranch{outStream: cli.outStream, errStream: cli.errStream}
	err := lb.GetLatestBranches(days, through)

	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Fprintln(cli.errStream, red("[error] "), err)
		return ExitCodeError
	}

	return ExitCodeOK
}

var usage = `Usage: git lbranch [options...]

lbranch is a git subcommand to show a list of recently committed branches.

OPTIONS:
  --days value, -d value  specifies the number of days branches last committed (default 5)
  --through, -t           print detailed explanation of branches, adding last commit hash and date
  --version, -v           print the current version
  --help, -h              print help

`
