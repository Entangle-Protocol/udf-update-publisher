package keystore

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// ParseKeyFromHex parses a private key from a hex string
func ParseKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	// Decode hex to bytes
	privateKeyBytes, err := hexutil.Decode(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
		return nil, err
	}

	// Load prviate key from bytes
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
		return nil, err
	}

	return privateKey, nil
}
