package entities

type LunaciaResponse struct {
	Success            bool      `json:"success"`
	RoninAddress       string    `json:"client_id"`
	LastClaimTimestamp int       `json:"last_claimed_item_at"`
	TotalSLP           int       `json:"total"`
	BlockchainRelated  Signature `json:"blockchain_related"`
}

type Signature struct {
	Signature struct {
		Amount int `json:"amount"`
	} `json:"signature"`
}