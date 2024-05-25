package main

import (
	"github.com/cryptomator/cryptolib-go"
)

func generateKey(algorithm string) ([]byte, error) {
	var key []byte
	switch algorithm {
	case "AES":
		key = make([]byte, 32) // 256-bit key for AES
	case "RSA":
		key = make([]byte, 2048) // 2048-bit key for RSA
	case "ECC":
		key = make([]byte, 32) // 256-bit key for ECC
	default:
		return nil, errors.New("unsupported encryption algorithm")
	}

	// Generate random key
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func storeKey(key []byte, algorithm string, storagePath string) error {
	// Store key in Cryptomator vault
	vault, err := cryptolib.NewVault(storagePath)
	if err != nil {
		return err
	}

	err = vault.StoreKey(key, algorithm)
	if err != nil {
		return err
	}

	return nil
}

func retrieveKey(algorithm string, storagePath string) ([]byte, error) {
	// Retrieve key from Cryptomator vault
	vault, err := cryptolib.NewVault(storagePath)
	if err != nil {
		return nil, err
	}

	key, err := vault.RetrieveKey(algorithm)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func rotateKey(algorithm string, storagePath string) ([]byte, error) {
	// Generate new key
	newKey, err := generateKey(algorithm)
	if err != nil {
		return nil, err
	}

	// Store new key
	err = storeKey(newKey, algorithm, storagePath)
	if err != nil {
		return nil, err
	}

	return newKey, nil
}
