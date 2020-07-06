package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blang/semver/v4"

	"github.com/Luzifer/rconfig/v2"
)

const usage = `
Compare version numbers using SemVer version logic

Usage: vercmp <ver1> <ver2>

Output values:
  < 0 : if ver1 < ver2
    0 : if ver1 == ver2
  > 0 : if ver1 > ver2
`

var (
	cfg = struct {
		VersionAndExit bool `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	if err := rconfig.ParseAndValidate(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("vercmp %s\n", version)
		os.Exit(0)
	}
}

func main() {
	if len(rconfig.Args()) != 3 {
		fmt.Println(usage)
		os.Exit(2)
	}

	ver1, err := semver.Make(rconfig.Args()[1])
	if err != nil {
		log.Fatalf("Unable to parse ver1: %s", err)
	}

	ver2, err := semver.Make(rconfig.Args()[2])
	if err != nil {
		log.Fatalf("Unable to parse ver2: %s", err)
	}

	fmt.Println(ver1.Compare(ver2))
}
