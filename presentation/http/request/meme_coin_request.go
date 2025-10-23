package request

type CreateMemeCoinRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=50"`
	Symbol      string `json:"symbol" binding:"required,min=2,max=10"`
	Description string `json:"description" binding:"max=500"`
	TotalSupply string `json:"total_supply" binding:"required"`
	Decimals    uint8  `json:"decimals" binding:"required,min=0,max=18"`
	ImageURL    string `json:"image_url" binding:"omitempty,url"`
	Website     string `json:"website" binding:"omitempty,url"`
	Twitter     string `json:"twitter" binding:"omitempty"`
	Telegram    string `json:"telegram" binding:"omitempty"`
	Discord     string `json:"discord" binding:"omitempty"`
	Network     string `json:"network" binding:"required,oneof=ethereum bsc polygon arbitrum optimism"`
}

type UpdateMemeCoinRequest struct {
	Name        string `json:"name" binding:"omitempty,min=1,max=50"`
	Description string `json:"description" binding:"omitempty,max=500"`
	ImageURL    string `json:"image_url" binding:"omitempty,url"`
	Website     string `json:"website" binding:"omitempty,url"`
	Twitter     string `json:"twitter" binding:"omitempty"`
	Telegram    string `json:"telegram" binding:"omitempty"`
	Discord     string `json:"discord" binding:"omitempty"`
}

type SearchMemeCoinsRequest struct {
	Query  string `form:"q" binding:"omitempty"`
	Limit  int    `form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int    `form:"offset" binding:"omitempty,min=0"`
	Network string `form:"network" binding:"omitempty"`
	Status string `form:"status" binding:"omitempty"`
}

type DeployMemeCoinRequest struct {
	GasPrice string `json:"gas_price" binding:"omitempty"`
	GasLimit uint64 `json:"gas_limit" binding:"omitempty,min=21000"`
}