// Section 3: Web Development

// HTTP Server
package section3

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func StartServer() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}

	apiURL := os.Getenv("FLUTTERWAVE_API_URL")
	apiKey := os.Getenv("FLUTTERWAVE_API_KEY")
	if apiURL == "" || apiKey == "" {
		log.Fatalf("API configuration error")
	}

	paymentService := NewPaymentService(apiURL, apiKey)
	walletService := NewWalletService()

	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/payment", HandlePayment(paymentService))
	http.HandleFunc("/create-wallet", HandleCreateWallet(walletService))	

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
