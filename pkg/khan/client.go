package khan

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func NewKhanClient(deviceId string, userAgent string, loginInfo LoginInfo) KhanClient {
	kc := KhanClient{
		DeviceId:  deviceId,
		UserAgent: userAgent,
		LoginInfo: loginInfo,
	}
	kc.initHttpClient()
	return kc
}

func LoadKhanClient(name string) (KhanClient, error) {
	contents, err := os.ReadFile(name)
	var kc KhanClient
	if err != nil {
		return kc, err
	}
	err = json.Unmarshal(contents, &kc)
	if err != nil {
		return kc, err
	}
	kc.initHttpClient()
	return kc, err
}

func (kc *KhanClient) initHttpClient() {
	tlsConfig := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		},
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
	kc.HttpClient = client
}

func (kc *KhanClient) Login() (LoginState, error) {
	request_method := "POST"
	request_url := "https://e.khanbank.com/v1/cfrm/auth/token"
	loginRequest := LoginRequest{
		Username:   kc.LoginInfo.Username,
		Password:   kc.LoginInfo.Password,
		GrantType:  "password",
		ChannelID:  "I",
		LanguageID: "003",
	}
	payload, err := json.Marshal(loginRequest)
	if err != nil {
		return NotLoggedIn, err
	}

	req, err := http.NewRequest(request_method, request_url, bytes.NewBuffer(payload))

	if err != nil {
		log.Println("Failed to create request")
		return NotLoggedIn, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("device-id", kc.DeviceId)
	if kc.UserAgent != "" {
		req.Header.Set("User-Agent", kc.UserAgent)
	}

	res, err := kc.HttpClient.Do(req)
	if err != nil {
		log.Println("Failed to send request")
		return NotLoggedIn, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to read response")
		return NotLoggedIn, err
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		log.Println("Failed to parse response", string(body))
		return NotLoggedIn, err
	}

	if res.StatusCode != 200 {
		return NotLoggedIn, &ClientError{Message: loginResponse.Message}
	}

	kc.UserInfo = UserInfo{
		AccessToken:           loginResponse.AccessToken,
		AccessTokenExpiresIn:  loginResponse.AccessTokenExpiresIn,
		RefreshToken:          loginResponse.RefreshToken,
		RefreshTokenStatus:    loginResponse.RefreshTokenStatus,
		RefreshTokenExpiresIn: loginResponse.RefreshTokenExpiresIn,
		DisplayName:           loginResponse.DisplayName,
		PrimaryAccountID:      loginResponse.PrimaryAccountID,
		UniqueID:              loginResponse.UniqueID,
	}

	if kc.UserInfo.AccessToken != "" {
		return LoggedIn, nil
	}
	return MFARequired, nil
}
