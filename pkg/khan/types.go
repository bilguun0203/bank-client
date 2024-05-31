package khan

import (
	"net/http"
	"time"
)

type LoginState int

const (
	LoginStateNotLoggedIn LoginState = iota
	LoginStateLoggedIn
	LoginStateMFARequired
)

type LoginType int

const (
	LoginTypeInitial LoginType = iota
	LoginTypeSecond
	LoginTypeFinal
)

type KhanClient struct {
	DeviceId   string       `json:"deviceId"`
	UserAgent  string       `json:"userAgent"`
	UserInfo   UserInfo     `json:"userInfo"`
	LoginInfo  LoginInfo    `json:"loginInfo"`
	HttpClient *http.Client `json:"-"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	AccessToken           string `json:"access_token,omitempty"`
	AccessTokenExpiresIn  string `json:"access_token_expires_in,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	RefreshTokenStatus    string `json:"refresh_token_status,omitempty"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in,omitempty"`
	DisplayName           string `json:"display_name,omitempty"`
	PrimaryAccountID      string `json:"primary_account_id,omitempty"`
	UniqueID              string `json:"-"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	// Code    string `json:"code,omitempty"`
	// Error   string `json:"error,omitempty"`
}

type LoginRequest struct {
	GrantType      string `json:"grant_type"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	ChannelID      string `json:"channelId"`
	LanguageID     string `json:"languageId"`
	IsPrelogin     string `json:"isPrelogin,omitempty"`
	RequestID      string `json:"requestId,omitempty"`
	SecondaryMode  string `json:"secondaryMode,omitempty"`
	RememberDevice string `json:"rememberDevice,omitempty"`
}

type LoginResponse struct {
	AccessToken                      string    `json:"access_token,omitempty"`
	AccessTokenExpiresIn             string    `json:"access_token_expires_in,omitempty"`
	RefreshToken                     string    `json:"refresh_token,omitempty"`
	RefreshTokenStatus               string    `json:"refresh_token_status,omitempty"`
	RefreshTokenExpiresIn            string    `json:"refresh_token_expires_in,omitempty"`
	OrganizationName                 string    `json:"organization_name,omitempty"`
	TokenType                        string    `json:"token_type,omitempty"`
	DefaultAccount                   string    `json:"default_account,omitempty"`
	Business                         string    `json:"business,omitempty"`
	Business2                        string    `json:"business2,omitempty"`
	Birthday                         string    `json:"birthday,omitempty"`
	RegisterNo                       string    `json:"register_no,omitempty"`
	Lang                             string    `json:"lang,omitempty"`
	SendEmail                        string    `json:"send_email,omitempty"`
	SendSms                          string    `json:"send_sms,omitempty"`
	LastIP                           string    `json:"last_ip,omitempty"`
	LoginCount                       string    `json:"login_count,omitempty"`
	LastDate                         time.Time `json:"last_date,omitempty"`
	AutoExt                          string    `json:"auto_ext,omitempty"`
	Address                          string    `json:"address,omitempty"`
	BillCustid                       string    `json:"bill_custid,omitempty"`
	ConnectedTime                    string    `json:"connected_time,omitempty"`
	OptionLoginName                  string    `json:"option_login_name,omitempty"`
	ShowAccessConfirm                string    `json:"show_access_confirm,omitempty"`
	Segment                          string    `json:"segment,omitempty"`
	DisplayName                      string    `json:"display_name,omitempty"`
	Department                       string    `json:"department,omitempty"`
	ShowSecAqAlert                   string    `json:"show_sec_aq_alert,omitempty"`
	CreateSecAq                      string    `json:"create_sec_aq,omitempty"`
	ChangePass                       string    `json:"change_pass,omitempty"`
	Mobile                           string    `json:"mobile,omitempty"`
	Email                            string    `json:"email,omitempty"`
	LastNameEn                       string    `json:"last_name_en,omitempty"`
	FirstNameEn                      string    `json:"first_name_en,omitempty"`
	LastName                         string    `json:"last_name,omitempty"`
	FirstName                        string    `json:"first_name,omitempty"`
	RoleID                           string    `json:"role_id,omitempty"`
	UserID                           string    `json:"user_id,omitempty"`
	CustomerType                     string    `json:"customer_type,omitempty"`
	EmailMasked                      string    `json:"email_masked,omitempty"`
	MobileMasked                     string    `json:"mobile_masked,omitempty"`
	LastNameMasked                   string    `json:"last_name_masked,omitempty"`
	LastNameEnMasked                 string    `json:"last_name_en_masked,omitempty"`
	Cif                              string    `json:"cif,omitempty"`
	OrgCode                          string    `json:"org_code,omitempty"`
	Feedback                         string    `json:"feedback,omitempty"`
	BankID                           string    `json:"bank_id,omitempty"`
	PrincipalID                      string    `json:"principal_id,omitempty"`
	Salutation                       string    `json:"salutation,omitempty"`
	DivisionAccessIndicator          string    `json:"division_access_indicator,omitempty"`
	LoginDate                        time.Time `json:"login_date,omitempty"`
	LogoffDate                       time.Time `json:"logoff_date,omitempty"`
	MultiCurrencyTxnAllowed          string    `json:"multi_currency_txn_allowed,omitempty"`
	CategoryCode                     string    `json:"category_code,omitempty"`
	LimitScheme                      string    `json:"limit_scheme,omitempty"`
	RangeLimitScheme                 string    `json:"range_limit_scheme,omitempty"`
	LoginAllowed                     string    `json:"login_allowed,omitempty"`
	TransactionAllowedaccountFormat  string    `json:"transaction_allowedaccount_format,omitempty"`
	AccountFormat                    string    `json:"account_format,omitempty"`
	AccessScheme                     string    `json:"access_scheme,omitempty"`
	LoginChannel                     string    `json:"login_channel,omitempty"`
	LastUnsuccessfullLoginChannel    string    `json:"last_unsuccessfull_login_channel,omitempty"`
	DateFormat                       string    `json:"date_format,omitempty"`
	TransactionAuthorizationScheme   string    `json:"transaction_authorization_scheme,omitempty"`
	CalendarType                     string    `json:"calendar_type,omitempty"`
	UnifiedLoginUser                 string    `json:"unified_login_user,omitempty"`
	ConfidentialTransactionAccess    string    `json:"confidential_transaction_access,omitempty"`
	AmountFormat                     string    `json:"amount_format,omitempty"`
	OutOfOfficePreference            string    `json:"out_of_office_preference,omitempty"`
	AccountMaskingRequired           string    `json:"account_masking_required,omitempty"`
	ProspectIndicator                string    `json:"prospect_indicator,omitempty"`
	LastUnsuccessfullLoginTime       time.Time `json:"last_unsuccessfull_login_time,omitempty"`
	PhoneNumber                      string    `json:"phone_number,omitempty"`
	ConnectionID                     string    `json:"connection_id,omitempty"`
	DeviceConfirmtype                string    `json:"device_confirmtype,omitempty"`
	SecQaConfirmed                   string    `json:"sec_qa_confirmed,omitempty"`
	DeviceConfirmTryCount            string    `json:"device_confirm_try_count,omitempty"`
	LoginOptionShowAccessConfirm     string    `json:"login_option_show_access_confirm,omitempty"`
	DeviceConfirmBlocked             string    `json:"device_confirm_blocked,omitempty"`
	DeviceIsPermanent                string    `json:"device_is_permanent,omitempty"`
	LocationIsPermanent              string    `json:"location_is_permanent,omitempty"`
	DeviceConfirmOption              string    `json:"device_confirm_option,omitempty"`
	LocationConfirmOption            string    `json:"location_confirm_option,omitempty"`
	SystemSettings                   string    `json:"system_settings,omitempty"`
	PrimaryAccountID                 string    `json:"primary_account_id,omitempty"`
	UniqueID                         string    `json:"unique_id,omitempty"`
	Bucketid                         string    `json:"bucketid,omitempty"`
	ForceSignonPasswordChangeFlag    string    `json:"force_signon_password_change_flag,omitempty"`
	ForceTxnPasswordChangeFlag       string    `json:"force_txn_password_change_flag,omitempty"`
	ForceSecurityQuestionChangeFlag  string    `json:"force_security_question_change_flag,omitempty"`
	ForceUssdPinChangeFlag           string    `json:"force_ussd_pin_change_flag,omitempty"`
	ForceModeChangeFlag              string    `json:"force_mode_change_flag,omitempty"`
	PrimaryAuthenticationMode        string    `json:"primary_authentication_mode,omitempty"`
	AuthorizationMode                string    `json:"authorization_mode,omitempty"`
	TotalLoginCount                  string    `json:"total_login_count,omitempty"`
	ActualLoginCount                 string    `json:"actual_login_count,omitempty"`
	MaximumSkipCount                 string    `json:"maximum_skip_count,omitempty"`
	SegmentCode                      string    `json:"segment_code,omitempty"`
	SecondaryAuthenticationMode      string    `json:"secondary_authentication_mode,omitempty"`
	UserApplicableSecondaryAuthModes string    `json:"user_applicable_secondary_auth_modes,omitempty"`
	UserType                         string    `json:"user_type,omitempty"`
	CorporateID                      string    `json:"corporate_id,omitempty"`
	EmailAddress                     string    `json:"email_address,omitempty"`
	MobileNumber                     string    `json:"mobile_number,omitempty"`
	CMPhoneNo                        string    `json:"c_m_phone_no,omitempty"`
}

type StaticCode struct {
	CodeType        string `json:"codeType,omitempty"`
	CmCode          string `json:"cmCode,omitempty"`
	CodeDescription string `json:"codeDescription,omitempty"`
}

type CurrencyAmount struct {
	Amount   int    `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type Transaction struct {
	TransactionDate    time.Time      `json:"transactionDate,omitempty"`
	AccountID          string         `json:"accountId,omitempty"`
	AmountType         StaticCode     `json:"amountType,omitempty"`
	Amount             CurrencyAmount `json:"amount,omitempty"`
	TransactionRemarks string         `json:"transactionRemarks,omitempty"`
	TxnTime            string         `json:"txnTime,omitempty"`
	BeginBalance       CurrencyAmount `json:"beginBalance,omitempty"`
	EndBalance         CurrencyAmount `json:"endBalance,omitempty"`
	TxnBranchID        string         `json:"txnBranchId,omitempty"`
}
