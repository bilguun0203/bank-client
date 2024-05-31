package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bilguun0203/bank-client/pkg/khan"
)

func Login(kc *khan.KhanClient) {
	scanner := bufio.NewScanner(os.Stdin)

	if kc.LoginInfo.Username == "" || kc.LoginInfo.Password == "" {
		fmt.Println("Username/Password is empty...")
		fmt.Printf("Username: ")
		scanner.Scan()
		username := scanner.Text()
		fmt.Printf("Password: ")
		scanner.Scan()
		password := scanner.Text()
		kc.LoginInfo.Username = username
		kc.LoginInfo.Password = base64.StdEncoding.EncodeToString([]byte(password))
	} else {
		fmt.Printf("Logging in with '%s' ...\n", kc.LoginInfo.Username)
	}

	state, err := kc.Login(khan.LoginTypeInitial, "")
	if err != nil {
		log.Fatal(err)
	}
	switch state {
	case khan.LoginStateLoggedIn:
		fmt.Println("Successfully logged in.")
	case khan.LoginStateMFARequired:
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Verification required.")
		fmt.Println("Sending OTP to your phone...")

		_, err = kc.Login(khan.LoginTypeSecond, "")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("OTP: ")
		scanner.Scan()
		otp := scanner.Text()
		otp = base64.RawStdEncoding.EncodeToString([]byte(otp))

		st, err := kc.Login(khan.LoginTypeFinal, otp)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(st)
		if st == khan.LoginStateLoggedIn {
			fmt.Println("Successfully logged in.")
		}
	}
}
