package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SmurfsAtWork/azrael/config"
	"github.com/SmurfsAtWork/azrael/log"
)

var (
	// Version contains the application's version number,
	// It's set by ldflags on build time.
	Version = ""

	// CommitSHA contains the commit's sha that this
	// current version was built against.
	CommitSHA = ""

	versionFlag = flag.Bool("version", false, "display version")
	helpFlag    = flag.Bool("help", false, "display this screen")

	configFlag = flag.String("config", "", `set:
	api-address <some api url>
		sets the api's url
del:
	api-address
		resets the api's url to the default value (https://papa.smurfsatwork.org)`)
)

func displayVersion() {
	if len(CommitSHA) > 7 {
		CommitSHA = CommitSHA[:7]
	}
	if Version == "" {
		Version = "(built from source)"
	}
	fmt.Printf("azrael %s", Version)
	if len(CommitSHA) > 0 {
		fmt.Printf(" (%s)", CommitSHA)
	}
	fmt.Println()
	os.Exit(0)
}

func handleConfig() {
	option := flag.Arg(0)
	newValue := flag.Arg(1)
	if len(newValue) == 0 && *configFlag != "del" {
		// passive aggressive usage
		flag.Usage()
		return
	}

	var err error
	switch *configFlag {
	case "set":
		switch option {
		case "api-address":
			err = config.SetApiAddress(newValue)
		default:
			// passive aggressive usage
			flag.Usage()
			return
		}
	case "del":
		switch option {
		case "api-address":
			err = config.ResetApiAddress()
		}
	case "":
		fallthrough
	default:
		// passive aggressive usage
		flag.Usage()
	}
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(0)
}

func main() {
	flag.Parse()
	switch {
	case *versionFlag:
		displayVersion()
	case *helpFlag:
		flag.Usage()
		return
	case len(*configFlag) != 0:
		handleConfig()
	default:
		flag.Usage()
	}
}
