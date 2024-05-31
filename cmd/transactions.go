package cmd

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/bilguun0203/bank-client/pkg/khan"
	"github.com/bilguun0203/bank-client/utils"
)

func DownloadTransactions(kc *khan.KhanClient, savePath, accountNumber, currency, startDate, endDate string) error {
	if accountNumber == "" || startDate == "" || endDate == "" || currency == "" {
		return &utils.BankClientError{Message: "Some parameters are empty!"}
	}

	startDateT, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return &utils.BankClientError{Message: "Start date does not match the expected format YYYY-MM-DD"}
	}
	endDateT, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		return &utils.BankClientError{Message: "End date does not match the expected format YYYY-MM-DD"}
	}
	// endDateT = endDateT.Add(24 * time.Hour - 1 * time.Second)

	log.Println("Logging in...")
	state, err := kc.Login(khan.LoginTypeInitial, "")
	if err != nil {
		return err
	}
	if state != khan.LoginStateLoggedIn {
		return &utils.BankClientError{Message: "Not fully authenticated. Please login again."}
	}
	time.Sleep(100 * time.Millisecond)

	log.Printf("Getting transactions of '%s' from '%s' to '%s'...\n", accountNumber, startDateT.Format(time.DateOnly), endDateT.Format(time.DateOnly))
	transactions, err := kc.Transactions(accountNumber, currency, startDateT, endDateT)
	if err != nil {
		return err
	}
	log.Printf("%d transactions found.", len(transactions))

	if savePath != "" {
		trans, err := json.MarshalIndent(transactions, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(savePath, trans, 0644)
	} else {
		log.Println("#\tDate\tType\tAmount\tCurrency\tRemarks")
		for i, trans := range transactions {
			log.Printf("%d\t%s %s\t%s\t%d\t%s\t%s\n", i, trans.TransactionDate.Format(time.DateOnly), trans.TxnTime, trans.AmountType.CodeDescription, trans.Amount.Amount, trans.Amount.Currency, trans.TransactionRemarks)
		}
	}
	return nil
}
