package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	khan "github.com/bilguun0203/bank-account-checker/pkg/khan"
	"github.com/google/uuid"
)

func main() {
	client_config_path := flag.String("config", "config.json", "Config file path")
	create_new_file := flag.Bool("create", false, "Create blank config file")
	flag.Parse()

	if *create_new_file {
		blank_config, err := json.MarshalIndent(khan.KhanClient{
			DeviceId: uuid.NewString(),
			LoginInfo: khan.LoginInfo{
				Username: "username",
				Password: "base64_encoded_password",
			},
		}, "", "  ")
		if err != nil {
			log.Fatal("Failed to initialize blank config")
		}
		err = os.WriteFile(*client_config_path, blank_config, 0600)
		if err != nil {
			log.Fatal("Failed to save blank config")
		}
		log.Printf("Successfully saved blank config to %s\n", *client_config_path)
		os.Exit(0)
	}

	log.Printf("Loading %s\n", *client_config_path)
	khanClient, err := khan.LoadKhanClient(*client_config_path)
	if err != nil {
		log.Fatal(err)
	}

	state, err := khanClient.Login()
	if err != nil {
		log.Fatal(err)
	}
	clientState, err := json.MarshalIndent(khanClient, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(*client_config_path, clientState, 0600)
	if err != nil {
		log.Fatal(err)
	}

	switch state {
	case khan.LoggedIn:
		log.Println("LoggedIn")
	case khan.NotLoggedIn:
		log.Println("NotLoggedIn")
	case khan.MFARequired:
		log.Println("MFARequired")
	default:
		log.Println("No action!")
	}
}
