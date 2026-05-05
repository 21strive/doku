package usecases

import (
	"fmt"
	"testing"
)

func Test_dokuUseCase_GetSupportedBanks(t *testing.T) {
	tests := []struct {
		name string
	}{}

	u := NewDokuUseCase("client_id", "secret_key", "private_key", false)

	banks := u.GetSupportedBanks()
	fmt.Printf("Supported Banks: %+v\n", banks)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := u.GetSupportedBanks()

			if got[0].Name != "BANK BRI" {
				t.Errorf("GetSupportedBanks() = %v, want %v", got[0].Name, "BANK BRI")
			}

			t.Logf("Supported Banks: %+v", got)
			t.Logf("Total Supported Banks: %d", len(got))
		})
	}
}
