package encryption

import (
	"testing"
)

func TestAesGcmEncryptor(t *testing.T) {
	secretKey := "some-long-but-not-secure-random-key"
	encryptor, err := NewAesGcmEncryptor(secretKey)
	if err != nil {
		t.Errorf("NewAesGcmEncryptor() error = %v", err)
		return
	}

	tests := []struct {
		name      string
		plaintext string
		wantErr   bool
	}{
		{"Encrypt and Decrypt valid text", "test-refresh-token", false},
		{"Encrypt and Decrypt empty text", "", false},
		{"Encrypt and Decrypt special characters", "!@#$%^&*()_+-=", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Encrypt
			ciphertext, err := encryptor.Encrypt(tt.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Test Decrypt
			decryptedText, err := encryptor.Decrypt(ciphertext)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if decryptedText != tt.plaintext {
				t.Errorf("Decrypt() got = %v, want %v", decryptedText, tt.plaintext)
			}
		})
	}
}
