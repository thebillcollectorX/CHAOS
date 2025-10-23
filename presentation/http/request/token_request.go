package request

type CreateTokenRequest struct {
	Name        string `json:"name" form:"name" binding:"required,min=1,max=100"`
	Symbol      string `json:"symbol" form:"symbol" binding:"required,min=2,max=20"`
	Description string `json:"description" form:"description" binding:"max=1000"`
	TotalSupply string `json:"total_supply" form:"total_supply" binding:"required"`
	Decimals    uint8  `json:"decimals" form:"decimals" binding:"min=0,max=18"`
	Network     string `json:"network" form:"network" binding:"required"`
	ImageURL    string `json:"image_url" form:"image_url" binding:"url"`
	Website     string `json:"website" form:"website" binding:"omitempty,url"`
	Twitter     string `json:"twitter" form:"twitter" binding:"omitempty,url"`
	Telegram    string `json:"telegram" form:"telegram" binding:"omitempty,url"`
	Discord     string `json:"discord" form:"discord" binding:"omitempty,url"`

	// Features
	IsMintable        bool   `json:"is_mintable" form:"is_mintable"`
	IsBurnable        bool   `json:"is_burnable" form:"is_burnable"`
	IsPausable        bool   `json:"is_pausable" form:"is_pausable"`
	HasMaxSupply      bool   `json:"has_max_supply" form:"has_max_supply"`
	HasTaxes          bool   `json:"has_taxes" form:"has_taxes"`
	BuyTaxPercentage  uint8  `json:"buy_tax_percentage" form:"buy_tax_percentage" binding:"min=0,max=25"`
	SellTaxPercentage uint8  `json:"sell_tax_percentage" form:"sell_tax_percentage" binding:"min=0,max=25"`
	IsAntiWhale       bool   `json:"is_anti_whale" form:"is_anti_whale"`
	MaxTxAmount       string `json:"max_tx_amount" form:"max_tx_amount"`
	MaxWalletAmount   string `json:"max_wallet_amount" form:"max_wallet_amount"`
}

type UpdateTokenRequest struct {
	Name        string `json:"name" form:"name" binding:"omitempty,min=1,max=100"`
	Description string `json:"description" form:"description" binding:"omitempty,max=1000"`
	ImageURL    string `json:"image_url" form:"image_url" binding:"omitempty,url"`
	Website     string `json:"website" form:"website" binding:"omitempty,url"`
	Twitter     string `json:"twitter" form:"twitter" binding:"omitempty,url"`
	Telegram    string `json:"telegram" form:"telegram" binding:"omitempty,url"`
	Discord     string `json:"discord" form:"discord" binding:"omitempty,url"`
}

type DeployTokenRequest struct {
	TokenID   uint `json:"token_id" form:"token_id" binding:"required"`
	NetworkID uint `json:"network_id" form:"network_id" binding:"required"`
}

type ConnectWalletRequest struct {
	Address string `json:"address" form:"address" binding:"required,len=42"`
	Type    string `json:"type" form:"type" binding:"required,oneof=metamask walletconnect"`
}