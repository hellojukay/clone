package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var repo string

func init() {
	flag.StringVar(&repo, "repo", "", "github repository https URL")
	flag.Parse()
	if repo == "" {
		flag.PrintDefaults()
	}
	// check repo url
	if strings.HasPrefix(repo, "git@") {
		fmt.Fprintf(os.Stderr, "%s\n", "repository ssh URL not support noew, please use https URl")
		os.Exit(1)
	}
	if !strings.HasPrefix(repo, "https://github.com") {
		fmt.Fprintf(os.Stderr, "%s\n", "only support github repository now")
		os.Exit(1)
	}
	repo = strings.Replace(repo, "github.com", "hub.fastgit.org", 1)
}

func main() {
	var clone = fmt.Sprintf("git clone %s", repo)
	fmt.Printf("%s\n", clone)
	var c = exec.Command("git", "clone", repo)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "run git clone failed %s", err)
		os.Exit(1)
	}
}
