

package section3


type VirtualCardRequest struct {
	Currency         string `json:"currency"`
	Amount           int    `json:"amount"`
	DebitCurrency    string `json:"debit_currency"`
	BillingName      string `json:"billing_name"`
	BillingAddress   string `json:"billing_address"`
	BillingCity      string `json:"billing_city"`
	BillingState     string `json:"billing_state"`
	BillingPostalCode string `json:"billing_postal_code"`
	BillingCountry   string `json:"billing_country"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	DateOfBirth      string `json:"date_of_birth"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Title            string `json:"title"`
	Gender           string `json:"gender"`
	CallbackURL      string `json:"callback_url"`
}


type PaymentRequest struct {
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	AccountBank   string `json:"account_bank"`
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	Currency      string `json:"currency"`
	Email         string `json:"email"`
	TxRef         string `json:"tx_ref"`
}
