package main

import (
	"flag"
	"log"
	"os"

	"github.com/bilguun0203/bank-client/cmd"
	"github.com/bilguun0203/bank-client/pkg/khan"
)

func main() {
	config_path := flag.String("config", "clients/config.json", "Config file path")
	device_id_flag := flag.String("deviceid", "", "Client UUID (used with -create)")
	user_agent_flag := flag.String("useragent", "", "User agent for the client (used with -create)")
	username_flag := flag.String("username", "", "Login username (used with -create)")
	password_flag := flag.String("password", "", "Login password (used with -create)")
	create_new_file := flag.Bool("create", false, "Create blank config file")
	login_flag := flag.Bool("login", false, "Interactive login")
	flag.Parse()

	if *create_new_file {
		err := cmd.CreateBlankConfig(*config_path, *device_id_flag, *user_agent_flag, *username_flag, *password_flag)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Successfully saved blank config to %s\n", *config_path)
		os.Exit(0)
	}

	log.Printf("Loading %s\n", *config_path)
	khanClient, err := khan.LoadKhanClient(*config_path)
	if err != nil {
		log.Fatal(err)
	}

	if *login_flag {
		cmd.Login(&khanClient)
		err = khanClient.SaveKhanClient(*config_path)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Updated the configuration file.")
		os.Exit(0)
	}
}
