package main

import (
	"fmt"
	"io"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type LBranch struct {
	outStream, errStream io.Writer
}

func (lb *LBranch) GetLatestBranches(days int, through bool) error {
	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	iterator, err := r.Branches()
	if err != nil {
		return err
	}

	err = iterator.ForEach(func(reference *plumbing.Reference) error {
		co, err := r.CommitObject(reference.Hash())
		if err != nil {
			return err
		}

		if co.Author.When.After(time.Now().AddDate(0, 0, -days)) {

			if through {
				// branch's ref is in the following form: refs/heads/branchName
				fmt.Fprint(lb.outStream, "  ")
				fmt.Fprint(lb.outStream, "[Name] ")
				fmt.Fprintln(lb.outStream, reference.Name()[11:])

				fmt.Fprint(lb.outStream, "  ")
				fmt.Fprint(lb.outStream, "[Hash] ")
				fmt.Fprintln(lb.outStream, reference.Hash())

				fmt.Fprint(lb.outStream, "  ")
				fmt.Fprint(lb.outStream, "[Date] ")
				fmt.Fprintln(lb.outStream, co.Author.When)
				fmt.Fprintln(lb.outStream, "")
				return nil
			}

			// branch's ref is in the following form: refs/heads/branchName
			fmt.Fprint(lb.outStream, "  ")
			fmt.Fprintln(lb.outStream, reference.Name()[11:])
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
