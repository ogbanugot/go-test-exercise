## Coding Exercise

- golang - [View Golang Coding Exercise](golang/README.md)


## Sections Overview

### Section 1: Language Fundamentals

1. **Variables and Types**
   - `variables_and_types.go`: Contains a function to swap the values of two variables without using a temporary variable.

2. **Slices**
   - `slices.go`: Contains a function to sum all even numbers in a slice of integers.

3. **Interfaces**
   - `interfaces.go`: Defines a `Logger` interface and implements it in `FileLogger` (logs messages to a file) and `ConsoleLogger` (logs messages to the console).

### Section 2: Concurrency

1. **Goroutines**
   - `goroutines.go`: Contains a function that uses goroutines to calculate the sum of elements in a large array concurrently.

2. **Channels**
   - `channels.go`: Implements a producer-consumer scenario using channels. The producer generates random numbers, and the consumer calculates and prints their squares.

### Section 3: Web Development

1. **HTTP Server**
   - `http_server.go`: Creates a simple HTTP server that responds with "Hello, World!" to any incoming request.

2. **Services**
   - `services.go`: Implementation for Go backend services using Flutterwave for Payments and Stellar blockchain for wallet address creation.
   ##### Test the services
   Needs `.env` with the following (see `.env.example`):
   ```txt
   FLUTTERWAVE_API_URL='https://api.flutterwave.com'
   FLUTTERWAVE_API_KEY='<secret-key>'
   ```
   Run the server (also runs everything in the sections, but server will run last)
   ```sh
   go run main.go
   ```
   Handling Payments (Direct Charge)
   ```sh
   curl -X POST 'http://localhost:8080/payment?type=debit_ng_account' \
   -H "Content-Type: application/json" \
   --data '{
   "firstname":"John",
   "lastname":"Doe",
   "account_bank": "044",
   "account_number": "0690000037",
   "amount": 7500,
   "currency": "NGN",
   "email": "twista@rove.press",
   "tx_ref": "BJUYU399fcd43"
   }'
   ```
   Create wallet address
   ```sh
   curl -X POST 'http://localhost:8080/create-wallet'
   ```

### Section 4: Testing

1. **Unit Testing**
   - `unit_test.go`: Contains unit tests for the `SumEven` function from Section 1, covering various scenarios including edge cases.

2. **HTTP Testing**
   - `http_test.go`: Contains a test for the HTTP server from Section 3, verifying that it responds correctly to different HTTP methods.

### Section 5: Code Review

1. **Code Review**
   - `golang/snippet.go`: Reviewed and corrected the provided code snippet. The `Factorial` function logic was fixed to correctly calculate the factorial of a number.


### Run everything
```sh
go run main.go
```
