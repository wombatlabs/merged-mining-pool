package bitcoin

import (
	"regexp"
)

type BitcoinCash struct{}

func (BitcoinCash) ChainName() string {
	return "bitcoincash"
}

func (BitcoinCash) CoinbaseDigest(coinbase string) (string, error) {
	return DoubleSha256(coinbase)
}

func (BitcoinCash) HeaderDigest(header string) (string, error) {
	return DoubleSha256(header)
}

func (BitcoinCash) ShareMultiplier() float64 {
	return 1.0
}

func (BitcoinCash) ValidMainnetAddress(address string) bool {
	// Matches either the CashAddr format (with optional "bitcoincash:" prefix, starting with q or p)
	// or the legacy Base58 format (starting with 1 or 3).
	return regexp.MustCompile(`^((bitcoincash:)?(q|p)[0-9a-z]{41}|[13][a-km-zA-HJ-NP-Z1-9]{25,34})$`).MatchString(address)
}

func (BitcoinCash) ValidTestnetAddress(address string) bool {
	// Matches either the testnet CashAddr format (with optional "bchtest:" prefix, starting with q or p)
	// or the legacy Base58 format (starting with m or n).
	return regexp.MustCompile(`^((bchtest:)?(q|p)[0-9a-z]{41}|(m|n)[a-km-zA-HJ-NP-Z1-9]{25,34})$`).MatchString(address)
}

func (BitcoinCash) MinimumConfirmations() uint {
	return uint(BitcoinMinConfirmations)
}
