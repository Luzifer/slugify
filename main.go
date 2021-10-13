package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Luzifer/rconfig/v2"
	"github.com/gosimple/slug"
)

var (
	cfg = struct {
		MaxLength      int  `flag:"max-length,l" default:"0" description:"Maximum length of returned slug (0 = no limit)"`
		VersionAndExit bool `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	if err := rconfig.Parse(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("git-changerelease %s\n", version)
		os.Exit(0)
	}
}

func main() {
	parts := rconfig.Args()[1:]

	var s string
	for {
		s = slug.Make(strings.Join(parts, " "))
		if cfg.MaxLength == 0 || len(s) < cfg.MaxLength {
			break
		}

		parts = parts[0 : len(parts)-1]
	}

	fmt.Println(s)
}
