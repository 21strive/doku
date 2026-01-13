package models

type Bank struct {
	Name                  string        `json:"name,omitempty"`
	BICode                string        `json:"bi_code,omitempty"`
	SwiftCode             string        `json:"swift_code,omitempty"`
	SupportedChannelCodes []ChannelCode `json:"supported_channel_code,omitempty"`
}

// Channel code 07 represents Bank Deposit (Online Transfer). Emoney only supports this channel code, such as gopay, dana, shopeepay.
// Channel code 11 represents BI Fast
type ChannelCode struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}
