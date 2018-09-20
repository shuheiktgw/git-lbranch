package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	cases := []struct {
		command           string
		expectedOutput    string
		expectedErrorCode int
	}{
		{command: "git-lbranch", expectedOutput: "test", expectedErrorCode: ExitCodeOK},
		{command: "git-lbranch -days 10", expectedOutput: "test", expectedErrorCode: ExitCodeOK},
		{command: "git-lbranch -d 10", expectedOutput: "test", expectedErrorCode: ExitCodeOK},
		{command: "git-lbranch -through", expectedOutput: "[Name] test", expectedErrorCode: ExitCodeOK},
		{command: "git-lbranch -t", expectedOutput: "[Name] test", expectedErrorCode: ExitCodeOK},
	}

	for i, tc := range cases {
		// Create a test commit
		cmd := exec.Command("/bin/sh", "scripts/create_test_commit.sh")
		_, err := cmd.Output()

		if err != nil {
			t.Fatalf("#%d CLI.Run unexpected error occured while creating a test commit: %s", i, err)
		}

		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}

		args := strings.Split(tc.command, " ")
		if got, want := cli.Run(args), tc.expectedErrorCode; got != want {
			t.Errorf("#%d CLI.Run %q exits with %d, want %d", i, tc.command, got, want)
		}

		if got, want := outStream.String(), tc.expectedOutput; !strings.Contains(got, want) {
			t.Errorf("#%d CLI.Run output assetion failed: %s", i, err)
		}

		// Clean up the test commit
		cmd = exec.Command("/bin/sh", "scripts/delete_test_commit.sh")
		_, err = cmd.Output()

		if err != nil {
			t.Errorf("#%d CLI.Run unexpected error occured while deleting a test commit: %s", i, err)
		}
	}
}
