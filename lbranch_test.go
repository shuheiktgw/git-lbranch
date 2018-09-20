package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestLBranch_GetLatestBranches(t *testing.T) {
	cases := []struct {
		days      int
		through   bool
		assertion func(string) error
	}{
		{
			days:    5,
			through: false,
			assertion: func(got string) error {
				if want := "test"; !strings.Contains(got, want) {
					return fmt.Errorf("want: %s, got: %s", want, got)
				}

				return nil
			}},
		{
			days:    5,
			through: true,
			assertion: func(got string) error {
				if want := "[Name] test"; !strings.Contains(got, want) {
					return fmt.Errorf("want: %s, got: %s", want, got)
				}

				return nil
			}},
		{
			days:    0,
			through: false,
			assertion: func(got string) error {
				if want := ""; !strings.Contains(got, want) {
					return fmt.Errorf("want: %s, got: %s", want, got)
				}

				return nil
			}},
		{
			days:    0,
			through: true,
			assertion: func(got string) error {
				if want := ""; !strings.Contains(got, want) {
					return fmt.Errorf("want: %s, got: %s", want, got)
				}

				return nil
			}},
	}

	for i, tc := range cases {
		// Create a test commit
		cmd := exec.Command("/bin/sh", "scripts/create_test_commit.sh")
		_, err := cmd.Output()

		if err != nil {
			t.Fatalf("#%d LBranch.GetLatestBranches unexpected error occured while creating a test commit: %s", i, err)
		}

		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		lb := &LBranch{outStream: outStream, errStream: errStream}

		lb.GetLatestBranches(tc.days, tc.through)

		if err := tc.assertion(outStream.String()); err != nil {
			t.Errorf("#%d LBranch.GetLatestBranches output assetion failed: %s", i, err)
		}

		// Clean up the test commit
		cmd = exec.Command("/bin/sh", "scripts/delete_test_commit.sh")
		_, err = cmd.Output()

		if err != nil {
			t.Errorf("LBranch.GetLatestBranches unexpected error occured while deleting a test commit: %s", err)
		}
	}
}
