// Section 3: Web Development

package section3

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandlePayment handles the HTTP request for direct charge processing
// default is Account:debit_account_ng!
func HandlePayment(paymentService *PaymentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		paymentType := queryParams.Get("type")

		var paymentReq PaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&paymentReq); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Process the payment
		body, statusCode, err := paymentService.ProcessPayment(paymentType, paymentReq)
		if err != nil {
			log.Printf("Error processing payment: %s", err.Error())
			http.Error(w, err.Error(), statusCode)
			return
		}

		// Handle the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(body)
	}
}

// HandleCreateWallet handles the HTTP request for creating a new wallet address
func HandleCreateWallet(walletService *WalletService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pair, err := walletService.CreateWalletAddress()
		if err != nil {
			http.Error(w, "Failed to create wallet address", http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"public_key": pair.Address(),
			"secret_key": pair.Seed(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
