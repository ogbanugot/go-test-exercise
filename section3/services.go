package section3

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"fmt"
	"github.com/stellar/go/keypair"
)

// PaymentService handles payment processing
type PaymentService struct {
	APIURL string
	APIKey string
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(apiURL, apiKey string) *PaymentService {
	return &PaymentService{
		APIURL: apiURL,
		APIKey: apiKey,
	}
}

func (s *PaymentService) ProcessPayment(paymentType string, paymentReq PaymentRequest) ([]byte, int, error) {
	// Validate the input data
	if paymentReq.FirstName == "" || paymentReq.LastName == "" || paymentReq.AccountBank == "" ||
		paymentReq.AccountNumber == "" || paymentReq.Amount <= 0 || paymentReq.Email == "" || paymentReq.TxRef == "" {
		return nil, http.StatusBadRequest, errors.New("missing required fields")
	}

	requestBody, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("failed to marshal request body")
	}

	switch paymentType {
	case "debit_ng_account":
		return s.sendRequest(s.APIURL+"/v3/charges?type=debit_ng_account", s.APIKey, requestBody)
	default:
		return nil, http.StatusBadRequest, errors.New("unsupported payment type")
	}
}

// WalletService handles wallet address creation
type WalletService struct{}

// NewWalletService creates a new WalletService
func NewWalletService() *WalletService {
	return &WalletService{}
}

// CreateWalletAddress generates a new Stellar wallet address
func (s *WalletService) CreateWalletAddress() (*keypair.Full, error) {
	pair, err := keypair.Random()
	if err != nil {
		return nil, fmt.Errorf("error creating key pair: %v", err)
	}
	return pair, nil
}

// sendRequest sends an HTTP request with the specified parameters and returns the response body
func (s *PaymentService) sendRequest(apiURL, apiKey string, requestBody []byte) ([]byte, int, error) {
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return body, resp.StatusCode, nil
}
