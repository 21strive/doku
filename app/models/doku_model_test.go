package models

import (
	"encoding/json"
	"testing"
)

func TestFlexibleInt_UnmarshalJSON_FromInt(t *testing.T) {
	jsonData := []byte(`{"amount": 100000}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !result.Amount.Valid {
		t.Fatal("Expected Amount to be valid")
	}

	if result.Amount.Int64 != 100000 {
		t.Errorf("Expected 100000, got %d", result.Amount.Int64)
	}
}

func TestFlexibleInt_UnmarshalJSON_FromFloat(t *testing.T) {
	jsonData := []byte(`{"amount": 100000.0}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !result.Amount.Valid {
		t.Fatal("Expected Amount to be valid")
	}

	if result.Amount.Int64 != 100000 {
		t.Errorf("Expected 100000, got %d", result.Amount.Int64)
	}
}

func TestFlexibleInt_UnmarshalJSON_FromString(t *testing.T) {
	jsonData := []byte(`{"amount": "154397"}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !result.Amount.Valid {
		t.Fatal("Expected Amount to be valid")
	}

	if result.Amount.Int64 != 154397 {
		t.Errorf("Expected 154397, got %d", result.Amount.Int64)
	}
}

func TestFlexibleInt_UnmarshalJSON_FromStringFloat(t *testing.T) {
	jsonData := []byte(`{"amount": "100000.0"}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !result.Amount.Valid {
		t.Fatal("Expected Amount to be valid")
	}

	if result.Amount.Int64 != 100000 {
		t.Errorf("Expected 100000, got %d", result.Amount.Int64)
	}
}

func TestFlexibleInt_UnmarshalJSON_FromStringWithDecimals_ShouldFail(t *testing.T) {
	jsonData := []byte(`{"amount": "100000.5"}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err == nil {
		t.Fatal("Expected error for non-whole number string, got nil")
	}
}

func TestFlexibleInt_UnmarshalJSON_FromFloatWithDecimals_ShouldFail(t *testing.T) {
	jsonData := []byte(`{"amount": 100000.5}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err == nil {
		t.Fatal("Expected error for non-whole number, got nil")
	}
}

func TestFlexibleInt_UnmarshalJSON_FromNegative_ShouldFail(t *testing.T) {
	jsonData := []byte(`{"amount": -100}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err == nil {
		t.Fatal("Expected error for negative number, got nil")
	}
}

func TestFlexibleInt_UnmarshalJSON_Null(t *testing.T) {
	jsonData := []byte(`{"amount": null}`)

	var result struct {
		Amount FlexibleInt `json:"amount"`
	}

	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result.Amount.Valid {
		t.Error("Expected Amount to be invalid (null)")
	}
}

func TestFlexibleInt_MarshalJSON(t *testing.T) {
	data := struct {
		Amount FlexibleInt `json:"amount"`
	}{
		Amount: FlexibleInt{},
	}
	data.Amount.Int64 = 100000
	data.Amount.Valid = true

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := `{"amount":100000}`
	if string(jsonBytes) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonBytes))
	}
}

func TestFlexibleInt_MarshalJSON_Null(t *testing.T) {
	data := struct {
		Amount FlexibleInt `json:"amount"`
	}{
		Amount: FlexibleInt{},
	}
	data.Amount.Valid = false

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := `{"amount":null}`
	if string(jsonBytes) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonBytes))
	}
}

func TestDokuOrder_UnmarshalJSON_WithFloatAmount(t *testing.T) {
	// Simulate webhook from Doku with float amount
	jsonData := []byte(`{
		"invoice_number": "INV-001",
		"amount": 100000.0,
		"currency": "IDR"
	}`)

	var order DokuOrder
	err := json.Unmarshal(jsonData, &order)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !order.Amount.Valid {
		t.Fatal("Expected Amount to be valid")
	}

	if order.Amount.Int64 != 100000 {
		t.Errorf("Expected 100000, got %d", order.Amount.Int64)
	}
}

func TestDokuSettlement_UnmarshalJSON_WithFloatValue(t *testing.T) {
	// Simulate webhook from Doku with float settlement value
	jsonData := []byte(`{
		"bank_account_settlement_id": "SETTLE-001",
		"value": 74300.0,
		"type": "NETT"
	}`)

	var settlement DokuSettlement
	err := json.Unmarshal(jsonData, &settlement)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !settlement.Value.Valid {
		t.Fatal("Expected Value to be valid")
	}

	if settlement.Value.Int64 != 74300 {
		t.Errorf("Expected 74300, got %d", settlement.Value.Int64)
	}
}
