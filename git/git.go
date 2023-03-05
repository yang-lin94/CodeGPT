package git

import (
	"errors"
	"log"
	"os/exec"
)

var excludeFromDiff = []string{
	"package-lock.json",
	"pnpm-lock.yaml",
	// yarn.lock, Cargo.lock, Gemfile.lock, Pipfile.lock, etc.
	"*.lock",
	"go.sum",
}

func excludeFiles() []string {
	newFileLists := []string{}
	for _, f := range excludeFromDiff {
		newFileLists = append(newFileLists, ":(exclude)"+f)
	}

	return newFileLists
}

func diffNames() *exec.Cmd {
	args := []string{
		"diff",
		// "--staged",
		"--name-only",
	}

	args = append(args, excludeFiles()...)

	return exec.Command(
		"git",
		args...,
	)
}

func diffFiles() *exec.Cmd {
	args := []string{
		"diff",
		"--staged",
		"--ignore-all-space",
		"--diff-algorithm=minimal",
		"--function-context",
	}

	args = append(args, excludeFiles()...)

	return exec.Command(
		"git",
		args...,
	)
}

func Diff() (string, error) {
	output, err := diffNames().Output()
	if err != nil {
		log.Fatal(err)
	}
	if string(output) == "" {
		return "", errors.New("nothing output")
	}

	output, err = diffFiles().Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(output), nil
}