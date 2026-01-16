package constants

import (
	_ "embed"
	"encoding/json"

	"github.com/21strive/doku/app/models"
)

//go:embed doku_supported_banks.json
var supportedBanksJSON []byte

// GetSupportedBanks parses and returns the list of supported banks
func GetSupportedBanks() ([]models.Bank, error) {
	var banks []models.Bank
	err := json.Unmarshal(supportedBanksJSON, &banks)
	return banks, err
}

// MustGetSupportedBanks returns the list of supported banks, panics on error
func MustGetSupportedBanks() []models.Bank {
	banks, err := GetSupportedBanks()
	if err != nil {
		panic("failed to parse supported banks JSON: " + err.Error())
	}
	return banks
}
