package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SmurfsAtWork/azrael/actions"
	"github.com/SmurfsAtWork/azrael/config"
)

var (
	// Version contains the application's version number,
	// It's set by ldflags on build time.
	Version = ""

	// CommitSHA contains the commit's sha that this
	// current version was built against.
	CommitSHA = ""
)

func handleVersion() error {
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

	return nil
}

func handleSetConfig() error {
	configName := fs.Arg(0)
	newValue := fs.Arg(1)
	if newValue == "" {
		return errors.New("idiot")
	}

	var err error
	switch configName {
	case "api-address":
		err = config.SetApiAddress(newValue)
	default:
		err = errors.New("idiot")
	}

	return err
}

func handleDeleteConfig() error {
	configName := fs.Arg(0)

	var err error
	switch configName {
	case "api-address":
		err = config.ResetApiAddress()
	default:
		err = errors.New("idiot")
	}

	return err
}

func handleLogin(usernamePasswordPair string) error {
	usernamePassword := strings.SplitN(usernamePasswordPair, ":", 2)
	if len(usernamePassword) != 2 {
		return errors.New("Username and password weren't provided!")
	}

	username := usernamePassword[0]
	password := usernamePassword[1]

	err := usecases.Login(actions.LoginParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

func handleCreateSmurf() error {
	args := fs.Args()
	smurfName := ""
	smurfPassword := ""

	for i := 0; i < len(args)-1; i++ {
		if args[i] == "name" && args[i+1] != "password" {
			smurfName = args[i+1]
		}
		if args[i] == "password" && args[i+1] != "name" {
			smurfPassword = args[i+1]
		}
	}

	if smurfPassword == "" {
		return errors.New("idiot you need password")
	}

	err := usecases.CreateSmurf(actions.CreateSmurfParams{
		Name:     smurfName,
		Password: smurfPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func handleGetSmurf() error {
	return errors.New("get smurf not implemented")
}

func handleListSmurfs() error {
	return errors.New("list smurfs not implemented")
}

func handleUpdateSmurfPassword() error {
	return errors.New("update password not implemented")
}

func handleUpdateSmurfActiveCommand() error {
	return errors.New("update smurf command not implemented")
}

func handleDeleteSmurf() error {
	return errors.New("delete smurf not implemented")
}
