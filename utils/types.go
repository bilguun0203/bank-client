package utils

type BankClientError struct {
	Message string
}

func (m *BankClientError) Error() string {
	return m.Message
}
