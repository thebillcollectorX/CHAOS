package seeds

import (
	"github.com/tiagorlampert/CHAOS/entities"
	"gorm.io/gorm"
)

func SeedNetworks(db *gorm.DB) error {
	networks := []entities.Network{
		{
			Name:        "ethereum",
			DisplayName: "Ethereum Mainnet",
			ChainID:     1,
			RpcURL:      "https://mainnet.infura.io/v3/",
			ExplorerURL: "https://etherscan.io",
			Currency:    "ETH",
			IsActive:    true,
			GasPrice:    "20000000000", // 20 Gwei
		},
		{
			Name:        "bsc",
			DisplayName: "Binance Smart Chain",
			ChainID:     56,
			RpcURL:      "https://bsc-dataseed.binance.org/",
			ExplorerURL: "https://bscscan.com",
			Currency:    "BNB",
			IsActive:    true,
			GasPrice:    "5000000000", // 5 Gwei
		},
		{
			Name:        "polygon",
			DisplayName: "Polygon Mainnet",
			ChainID:     137,
			RpcURL:      "https://polygon-rpc.com/",
			ExplorerURL: "https://polygonscan.com",
			Currency:    "MATIC",
			IsActive:    true,
			GasPrice:    "30000000000", // 30 Gwei
		},
		{
			Name:        "avalanche",
			DisplayName: "Avalanche C-Chain",
			ChainID:     43114,
			RpcURL:      "https://api.avax.network/ext/bc/C/rpc",
			ExplorerURL: "https://snowtrace.io",
			Currency:    "AVAX",
			IsActive:    true,
			GasPrice:    "25000000000", // 25 Gwei
		},
		// Testnets
		{
			Name:        "goerli",
			DisplayName: "Ethereum Goerli Testnet",
			ChainID:     5,
			RpcURL:      "https://goerli.infura.io/v3/",
			ExplorerURL: "https://goerli.etherscan.io",
			Currency:    "ETH",
			IsActive:    true,
			GasPrice:    "20000000000",
		},
		{
			Name:        "bsc-testnet",
			DisplayName: "BSC Testnet",
			ChainID:     97,
			RpcURL:      "https://data-seed-prebsc-1-s1.binance.org:8545/",
			ExplorerURL: "https://testnet.bscscan.com",
			Currency:    "BNB",
			IsActive:    true,
			GasPrice:    "10000000000",
		},
	}

	for _, network := range networks {
		var existingNetwork entities.Network
		if err := db.Where("chain_id = ?", network.ChainID).First(&existingNetwork).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&network).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}

func SeedContractTemplates(db *gorm.DB) error {
	templates := []entities.ContractTemplate{
		{
			Name:        "Basic ERC-20",
			Description: "Standard ERC-20 token with basic functionality",
			Version:     "1.0.0",
			Solidity: `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract {{.Name}} is ERC20, Ownable {
    constructor() ERC20("{{.TokenName}}", "{{.Symbol}}") {
        _mint(msg.sender, {{.TotalSupply}} * 10**decimals());
    }
}`,
			IsActive: true,
		},
		{
			Name:        "Advanced ERC-20",
			Description: "ERC-20 token with advanced features like minting, burning, and pausing",
			Version:     "1.0.0",
			Solidity: `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract {{.Name}} is ERC20, ERC20Burnable, Pausable, Ownable {
    constructor() ERC20("{{.TokenName}}", "{{.Symbol}}") {
        _mint(msg.sender, {{.TotalSupply}} * 10**decimals());
    }

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    function _beforeTokenTransfer(address from, address to, uint256 amount)
        internal
        whenNotPaused
        override
    {
        super._beforeTokenTransfer(from, to, amount);
    }
}`,
			IsActive: true,
		},
		{
			Name:        "Tax Token",
			Description: "ERC-20 token with configurable buy/sell taxes",
			Version:     "1.0.0",
			Solidity: `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract {{.Name}} is ERC20, Ownable {
    uint256 public buyTaxPercentage = {{.BuyTax}};
    uint256 public sellTaxPercentage = {{.SellTax}};
    address public taxWallet;

    constructor() ERC20("{{.TokenName}}", "{{.Symbol}}") {
        _mint(msg.sender, {{.TotalSupply}} * 10**decimals());
        taxWallet = msg.sender;
    }

    function setBuyTax(uint256 _buyTax) external onlyOwner {
        require(_buyTax <= 25, "Tax cannot exceed 25%");
        buyTaxPercentage = _buyTax;
    }

    function setSellTax(uint256 _sellTax) external onlyOwner {
        require(_sellTax <= 25, "Tax cannot exceed 25%");
        sellTaxPercentage = _sellTax;
    }

    function setTaxWallet(address _taxWallet) external onlyOwner {
        taxWallet = _taxWallet;
    }

    function _transfer(address from, address to, uint256 amount) internal override {
        if (from == owner() || to == owner()) {
            super._transfer(from, to, amount);
            return;
        }

        uint256 taxAmount = 0;
        
        // Apply tax logic here based on DEX pairs
        // This is a simplified version
        if (buyTaxPercentage > 0 || sellTaxPercentage > 0) {
            taxAmount = (amount * buyTaxPercentage) / 100;
            if (taxAmount > 0) {
                super._transfer(from, taxWallet, taxAmount);
            }
        }

        super._transfer(from, to, amount - taxAmount);
    }
}`,
			IsActive: true,
		},
	}

	for _, template := range templates {
		var existingTemplate entities.ContractTemplate
		if err := db.Where("name = ?", template.Name).First(&existingTemplate).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&template).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}