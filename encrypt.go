package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "io"
)

func encrypt(data []byte, key []byte, algorithm string) ([]byte, error) {
    return nil, errors.New("encryption not implemented")
}
