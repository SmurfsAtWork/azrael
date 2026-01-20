package main

import (
	"flag"
	"os"

	"github.com/SmurfsAtWork/azrael/actions"
	"github.com/SmurfsAtWork/azrael/cfmt"
)

var (
	fs = flag.NewFlagSet("azrael", flag.ExitOnError)

	versionFlag = fs.Bool("version", false, "display version")
	helpFlag    = fs.Bool("help", false, "display this screen")
	loginFlag   = fs.String("login", "", `logs user in, username and password are provided like HTTP basic auth, i.e. username:password

example: azrael -login SomeUser:SomeTopSecretPassword`)

	configFlag = fs.String("config", "", `set:
	api-address <some papa api address>
		sets papa's api address
del:
	api-address
		resets the api's url to the default value (https://papa.smurfsatwork.org)`)

	smurfFlag = fs.String("smurf", "", `new:
	name: sets the smurf's apparent name
		you need to provide "name" before the desired name, or just omit it
	password: sets the smurf's password
		you need to provide "password" before the desired password, and it can't be omitted as you need a password to login the Smurf.

	example: azrael -smurf new name arnold password 123
del:
	id: (don't mix with nano-id)
		ID of the Smurf to be deleted
	nano-id: (don't mix with id)
		Nano ID of the Smurf to be deleted

	example: azrael -smurf del nano-id zcg3

get: fetches Smurf's active command, config, last 5 mins logs and stats.
	id: (don't mix with nano-id)
		ID of the Smurf to be fetched
	nano-id: (don't mix with id)
		Nano ID of the Smurf to be fetched

	example: azrael -smurf get nano-id zcg3

update:
	password: updates Smurf's password
		id:
			ID of the Smurf to be updated

		example: azrael -smurf update password id 5 P@s5w0RD

	command: updates Smurf's running command (find commands from -program list or -script list)
		id:
			ID of the Smurf to be updated

		example: azrael -smurf update command id 5 1 # where 1 is some program ID

list:
	No options for now, just a list of Smurfs lol`)

	usecases *actions.Actions
)

func init() {
	usecases = &actions.Actions{}
}

func exitUsage() {
	fs.Usage()
	os.Exit(2)
}

func main() {
	fs.Parse(os.Args[1:])

	var err error
	switch {
	case len(*loginFlag) != 0:
		err = handleLogin(*loginFlag)

	case len(*smurfFlag) != 0:
		switch *smurfFlag {
		case "new":
			err = handleCreateSmurf()
		case "del":
			err = handleDeleteSmurf()
		case "update":
			updated := fs.Arg(0)
			switch updated {
			case "password":
				err = handleUpdateSmurfPassword()
			case "command":
				err = handleUpdateSmurfActiveCommand()
			default:
				exitUsage()
			}
		case "list":
			err = handleListSmurfs()
		case "get":
			err = handleGetSmurf()
		default:
			exitUsage()
		}

	case *versionFlag:
		err = handleVersion()

	case *helpFlag:
		exitUsage()

	case len(*configFlag) != 0:
		switch *configFlag {
		case "set":
			err = handleSetConfig()
		case "del":
			err = handleDeleteConfig()
		default:
			exitUsage()
		}
	default:
		exitUsage()
	}

	if err != nil {
		cfmt.Red().Bold().Println(err)
	} else {
		cfmt.Green().Println("Cya~!")
	}
}
