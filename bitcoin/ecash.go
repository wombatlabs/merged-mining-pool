package ecash

import (
	"regexp"
)

type ECash struct{}

func (ECash) ChainName() string {
	return "ecash"
}

func (ECash) CoinbaseDigest(coinbase string) (string, error) {
	return DoubleSha256(coinbase)
}

func (ECash) HeaderDigest(header string) (string, error) {
	return DoubleSha256(header)
}

func (ECash) ShareMultiplier() float64 {
	return 1.0
}

func (ECash) ValidMainnetAddress(address string) bool {
	// Matches either the CashAddr format (with optional "ecash:" prefix, starting with q or p)
	// or the legacy Base58 format (starting with 1 or 3).
	return regexp.MustCompile(`^((ecash:)?(q|p)[0-9a-z]{41}|[13][a-km-zA-HJ-NP-Z1-9]{25,34})$`).MatchString(address)
}

func (ECash) ValidTestnetAddress(address string) bool {
	// Matches either the testnet CashAddr format (with optional "ectest:" prefix, starting with q or p)
	// or the legacy Base58 format (starting with m or n).
	return regexp.MustCompile(`^((ectest:)?(q|p)[0-9a-z]{41}|(m|n)[a-km-zA-HJ-NP-Z1-9]{25,34})$`).MatchString(address)
}

func (ECash) MinimumConfirmations() uint {
	return uint(BitcoinMinConfirmations)
}
