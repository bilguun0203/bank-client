package cmd

import (
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/bilguun0203/bank-client/pkg/khan"
	"github.com/bilguun0203/bank-client/utils"
	"github.com/google/uuid"
)

func CreateBlankConfig(name, deviceId, userAgent, username, password string) error {
	if deviceId == "" {
		deviceId = uuid.NewString()
	}
	if password != "" {
		password = base64.StdEncoding.EncodeToString([]byte(password))
	}

	blank_config, err := json.MarshalIndent(khan.KhanClient{
		DeviceId:  deviceId,
		UserAgent: userAgent,
		LoginInfo: khan.LoginInfo{
			Username: username,
			Password: password,
		},
	}, "", "  ")

	if err != nil {
		return &utils.BankClientError{Message: "Failed to initialize blank config"}
	}

	err = os.WriteFile(name, blank_config, 0600)
	if err != nil {
		return &utils.BankClientError{Message: "Failed to save blank config"}
	}

	return nil
}
