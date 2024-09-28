package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// AesGcmEncryptor implements the Encryptor interface using AES-GCM
type AesGcmEncryptor struct {
	cipherBlock cipher.Block
}

// NewAesGcmEncryptor creates a new AES-GCM Encryptor
func NewAesGcmEncryptor(secretKey string) (*AesGcmEncryptor, error) {
	hashedKey := sha256.Sum256([]byte(secretKey))

	// AES-256 key
	block, err := aes.NewCipher(hashedKey[:])
	if err != nil {
		return nil, fmt.Errorf("error creating new cipher: %w", err)
	}

	return &AesGcmEncryptor{block}, nil
}

// Encrypt encrypts the plaintext using AES-256 GCM
func (e *AesGcmEncryptor) Encrypt(plaintext string) (string, error) {
	gcm, err := cipher.NewGCM(e.cipherBlock)
	if err != nil {
		return "", fmt.Errorf("error creating new GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("error generating nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the ciphertext using AES-256 GCM
func (e *AesGcmEncryptor) Decrypt(inputCiphertext string) (string, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(inputCiphertext)
	if err != nil {
		return "", fmt.Errorf("error decoding ciphertext: %w", err)
	}

	gcm, err := cipher.NewGCM(e.cipherBlock)
	if err != nil {
		return "", fmt.Errorf("error creating new GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(decodedCiphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertext := decodedCiphertext[:nonceSize], decodedCiphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting: %w", err)
	}

	return string(plaintext), nil
}
