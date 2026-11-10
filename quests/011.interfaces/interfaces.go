package interfaces

type PaymentMethod interface {
	Process(amount float64) bool
	Provider() string
}

type CardPayment struct {
	CardNumber string
	Limit      float64
}

func (c *CardPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	if c.Limit >= amount {
		c.Limit -= amount
		return true
	}
	return false
}

func (c CardPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return "CARD"
}

type UPIPayment struct {
	VPA string
}

func (u UPIPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions

	return true
}

func (u UPIPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return "UPI"
}

type CryptoPayment struct {
	Wallet  string
	Balance float64
}

func (c *CryptoPayment) Process(amount float64) bool {
	// TODO: implement
	// Read README.md for the instructions
	if c.Balance >= amount {
		c.Balance -= amount
		return true
	}
	return false
}

func (c CryptoPayment) Provider() string {
	// TODO: implement
	// Read README.md for the instructions
	return "CRYPTO"
}

func Checkout(p PaymentMethod, amount float64) string {
	// TODO: implement
	// Read README.md for the instructions
	sucess := p.Process(amount)
	if sucess {
		return "Payment successful via " + p.Provider()
	}
	return "payment failed via " + p.Provider()
}

func DetectCrypto(p PaymentMethod) bool {
	// TODO: implement
	// Read README.md for the instructions
	provider := p.Provider()
	if provider == "CRYPTO" {
		return true
	}
	return false
}
