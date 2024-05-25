package main

import (
	"errors"
)

func generateKey(algorithm string) ([]byte, error) {
	var key []byte
	switch algorithm {
	case "AES":
		key = make([]byte, 32) 
	case "RSA":
		key = make([]byte, 2048) 
	case "ECC":
		key = make([]byte, 32) 
	default:
		return nil, errors.New("unsupported encryption algorithm")
	}

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func storeKey(key []byte, algorithm string, storagePath string) error {
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
	newKey, err := generateKey(algorithm)
	if err != nil {
		return nil, err
	}

	err = storeKey(newKey, algorithm, storagePath)
	if err != nil {
		return nil, err
	}

	return newKey, nil
}
