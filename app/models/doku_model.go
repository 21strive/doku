package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/guregu/null/v6"
)

// FlexibleInt is a custom type that can unmarshal from both int and float JSON values
// It validates that float values are whole numbers and stores them as int64
type FlexibleInt struct {
	null.Int
}

// UnmarshalJSON implements custom unmarshalling to handle int, float, and string
func (fi *FlexibleInt) UnmarshalJSON(data []byte) error {
	// Handle null
	if string(data) == "null" {
		fi.Valid = false
		return nil
	}

	var floatVal float64

	// Try unmarshal as string first
	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		// It's a string, parse it as float
		var parseErr error
		floatVal, parseErr = parseFloat(strVal)
		if parseErr != nil {
			return fmt.Errorf("cannot parse string %s as number: %w", string(data), parseErr)
		}
	} else {
		// Not a string, try as number (int or float)
		if err := json.Unmarshal(data, &floatVal); err != nil {
			return fmt.Errorf("cannot unmarshal %s into FlexibleInt: %w", string(data), err)
		}
	}

	// Validate that it's a whole number
	intVal := int64(floatVal)
	if floatVal != float64(intVal) {
		return fmt.Errorf("amount must be a whole number, got: %f", floatVal)
	}

	// Validate non-negative
	if floatVal < 0 {
		return fmt.Errorf("amount cannot be negative: %f", floatVal)
	}

	fi.Int64 = intVal
	fi.Valid = true
	return nil
}

// parseFloat is a helper to parse string to float64
func parseFloat(s string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
}

// MarshalJSON implements custom marshalling to always output as int
func (fi FlexibleInt) MarshalJSON() ([]byte, error) {
	// Delegate to null.Int's MarshalJSON
	return fi.Int.MarshalJSON()
}

type DokuSignatureComponent struct {
	ClientId         null.String `json:"Client-Id"`
	RequestId        null.String `json:"Request-Id"`
	RequestTimestamp null.String `json:"Request-Timestamp"`
	RequestTarget    null.String `json:"Request-Target"`
	Digest           null.String `json:"Digest"`
}

type DokuOrder struct {
	InvoiceNumber null.String `json:"invoice_number"`
	Amount        FlexibleInt `json:"amount"`
	Currency      null.String `json:"currency,omitempty"` // for response object
	SessionID     null.String `json:"session_id,omitempty"`
}

type DokuVirtualAccountInfo struct {
	ExpiredTime    null.Int    `json:"expired_time"`
	ReusableStatus null.Bool   `json:"reusable_status"`
	Info1          null.String `json:"info1,omitempty"`
	Info2          null.String `json:"info2,omitempty"`
	Info3          null.String `json:"info3,omitempty"`
}

type DokuCustomer struct {
	ID    null.String `json:"id,omitempty"`
	Name  null.String `json:"name"`
	Email null.String `json:"email"`
}

// will be used as SubAccount ID.
type DokuAdditionalInfo struct {
	Account DokuAccount `json:"account"`
}

// SubAccount ID
type DokuAccount struct {
	ID null.String `json:"id"`
}

type DokuPayment struct {
	PaymentMethodTypes []string    `json:"payment_method_types,omitempty"`
	PaymentDueDate     null.Int    `json:"payment_due_date,omitempty"`
	TokenID            null.String `json:"token_id,omitempty"`
	URL                null.String `json:"url,omitempty"`
	ExpiredDate        null.String `json:"expired_date,omitempty"`
}

type DokuPaymentAdditionalInfo struct {
	Origin struct {
		Product   null.String `json:"product,omitempty"`
		System    null.String `json:"system,omitempty"`
		ApiFormat null.String `json:"apiFormat,omitempty"`
		Source    null.String `json:"source,omitempty"`
	}
}

type DokuHeader struct {
	RequestID null.String `json:"request_id"`
	Signature null.String `json:"signature"`
	Date      *time.Time  `json:"date,omitempty"`
	ClientID  null.String `json:"client_id"`
}

type DokuBalance struct {
	Pending   null.String `json:"pending"`
	Available null.String `json:"available"`
}

type DokuTransaction struct {
	Type   null.String `json:"type,omitempty"`
	Status null.String `json:"status"`
	//Date              *time.Time  `json:"date"`
	OriginalRequestID null.String `json:"original_request_id"`
}

type DokuPaymentIdentifier struct {
	Name  null.String `json:"name"`
	Value null.String `json:"value"`
}
type DokuVirtualAccountpayment struct {
	ReferenceNumber null.String `json:"reference_number"`
	//Date            *time.Time               `json:"date"`
	Identifier []*DokuPaymentIdentifier `json:"identifier"`
}

type DokuCardPayment struct {
	MaskedCardNumber null.String `json:"masked_card_number"`
	ApprovalCode     null.String `json:"approval_code"`
	ResponseCode     null.String `json:"response_code"`
	ResponseMessage  null.String `json:"response_message"`
	Issuer           null.String `json:"issuer"`
	PaymentID        null.String `json:"payment_id"`
}

type DokuSettlement struct {
	BankAccountSettlementID null.String `json:"bank_account_settlement_id"`
	Value                   FlexibleInt `json:"value"`
	Type                    null.String `json:"type"`
}
