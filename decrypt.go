package main

import (
    "crypto/aes"
    "crypto/cipher"
    "errors"
    "io"
)

func decrypt(ciphertext []byte, key []byte, algorithm string) ([]byte, error) {
    return nil, errors.New("decryption not implemented")
}
