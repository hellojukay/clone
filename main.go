package main

import (
	"flag"
)

var repo string

func init() {
	flag.StringVar(&repo, "repo", "", "github repostiry https URL")
	flag.Parse()
	if repo == "" {
		flag.PrintDefaults()
	}

}
func main() {

}
