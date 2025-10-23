package request

type CreateTokenRequestForm struct {
	Name        string `json:"name" binding:"required"`
	Symbol      string `json:"symbol" binding:"required"`
	TotalSupply string `json:"total_supply" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Network     string `json:"network" binding:"required"`
	Decimals    int    `json:"decimals"`
}
