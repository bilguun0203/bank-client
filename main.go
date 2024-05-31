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
	// used for creating new config file
	create_new_file := flag.Bool("create", false, "Create blank config file")
	device_id_flag := flag.String("deviceid", "", "Client UUID (used with -create)")
	user_agent_flag := flag.String("useragent", "", "User agent for the client (used with -create)")
	username_flag := flag.String("username", "", "Login username (used with -create)")
	password_flag := flag.String("password", "", "Login password (used with -create)")
	// interactive login
	login_flag := flag.Bool("login", false, "Interactive login")
	// download and save transcations
	transactions_flag := flag.Bool("transactions", false, "Get transactions")
	account_number_flag := flag.String("account", "", "Account number (used with -transactions)")
	currency_flag := flag.String("currency", "MNT", "Transaction currency (used with -transactions)")
	start_date_flag := flag.String("start-date", "", "Start date (used with -transactions)")
	end_date_flag := flag.String("end-date", "", "End date (used with -transactions)")
	save_path := flag.String("save-path", "", "Transactions save path (.json optional) (used with -transactions)")
	flag.Parse()

	// Create new configuration file
	if *create_new_file {
		err := cmd.CreateBlankConfig(*config_path, *device_id_flag, *user_agent_flag, *username_flag, *password_flag)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Successfully saved blank config to %s\n", *config_path)
		os.Exit(0)
	}

	// Load existing configuration file
	log.Printf("Loading %s\n", *config_path)
	khanClient, err := khan.LoadKhanClient(*config_path)
	if err != nil {
		log.Fatal(err)
	}

	// Interactive Login
	if *login_flag {
		cmd.Login(&khanClient)
		err = khanClient.SaveKhanClient(*config_path)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Updated the configuration file.")
		os.Exit(0)
	}

	// Transaction downloader
	if *transactions_flag {
		err = cmd.DownloadTransactions(&khanClient, *save_path, *account_number_flag, *currency_flag, *start_date_flag, *end_date_flag)
		if err != nil {
			log.Fatal(err)
		}
		if *save_path != "" {
			log.Printf("Saved transactions to '%s'\n", *save_path)
		}
		os.Exit(0)
	}
}
