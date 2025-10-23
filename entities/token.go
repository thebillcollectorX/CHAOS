package entities

type Token struct {
	DBModel
	Name        string  `json:"name" binding:"required"`
	Symbol      string  `json:"symbol" binding:"required"`
	TotalSupply string  `json:"total_supply" binding:"required"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Network     string  `json:"network" binding:"required"` // e.g., "Ethereum", "BSC", "Solana"
	UserID      uint    `json:"user_id"`
	Decimals    int     `json:"decimals"`
	Address     string  `json:"address"` // Smart contract address after deployment
	Status      string  `json:"status"`  // pending, deployed, failed
	TxHash      string  `json:"tx_hash"` // Transaction hash
}
