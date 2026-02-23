package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"sync"
)

var (
	apiKeys   = map[int]string{}
	apiKeyMu  sync.RWMutex
	nextKeyID = 1
)

func CreateAPIKey() (int, string, error) {
	newKey, err := generateAPIKey()
	if err != nil {
		return 0, "", err
	}

	apiKeyMu.Lock()
	defer apiKeyMu.Unlock()

	keyID := nextKeyID
	nextKeyID++
	apiKeys[keyID] = newKey

	return keyID, newKey, nil
}

func DeleteAPIKeyByID(keyID int) bool {
	apiKeyMu.Lock()
	defer apiKeyMu.Unlock()

	if _, exists := apiKeys[keyID]; !exists {
		return false
	}

	delete(apiKeys, keyID)
	return true
}

func HasAnyAPIKey() bool {
	apiKeyMu.RLock()
	defer apiKeyMu.RUnlock()

	return len(apiKeys) > 0
}

func IsValidAPIKey(requestAPIKey string) bool {
	apiKeyMu.RLock()
	defer apiKeyMu.RUnlock()

	for _, currentAPIKey := range apiKeys {
		if subtle.ConstantTimeCompare([]byte(requestAPIKey), []byte(currentAPIKey)) == 1 {
			return true
		}
	}

	return false
}

func IsSecretMatch(input, expected string) bool {
	if input == "" || expected == "" {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(input), []byte(expected)) == 1
}

func generateAPIKey() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	return hex.EncodeToString(buf), nil
}
