package entities

type LunaciaResponse struct {
	Success            bool      `json:"success"`
	RoninAddress       string    `json:"clientID"`
	LastClaimTimestamp int       `json:"lastClaimedItemAt"`
	TotalSLP           int       `json:"total"`
	BlockchainRelated  Signature `json:"blockchainRelated"`
}

type Signature struct {
	Signature struct {
		Amount int `json:"amount"`
	} `json:"signature"`
}