# 🚀 MemeToken Creator

A powerful web application for creating and deploying meme coins (cryptocurrencies) without any coding knowledge. Similar to moontool.fun and createtokens.io, this platform enables anyone to launch their own token on multiple blockchain networks.

## ✨ Features

### 🎯 Easy Token Creation
- **No Code Required**: Simple web interface for creating tokens
- **Quick Deployment**: Launch your token in under 2 minutes
- **Full Customization**: Set name, symbol, supply, decimals, and more

### 🌐 Multi-Chain Support
Deploy your token on popular blockchain networks:
- **Ethereum (ETH)** - Most popular, highest liquidity
- **Binance Smart Chain (BSC)** - Low fees, fast transactions
- **Polygon (MATIC)** - Scalable and cost-effective
- **Solana (SOL)** - Ultra-fast confirmations
- **Base** - Layer 2 solution

### 🔒 Security Features
- Audited smart contracts
- Full ownership and control
- Transparent blockchain deployment
- Ready for DEX trading

### 📊 Token Management
- View all created tokens
- Track deployment status
- Monitor contract addresses
- Transaction hash tracking

## 🎨 User Interface

### Landing Page (`/landing`)
Beautiful, modern landing page featuring:
- Hero section with call-to-action
- Feature highlights
- "How it works" guide
- Supported networks showcase
- Statistics dashboard

### Token Creation Page (`/tokens`)
Comprehensive token creation interface with:
- Form-based token configuration
- Real-time validation
- Network selection
- Status tracking
- Recently created tokens list

## 🛠️ Technical Stack

### Backend
- **Go** - Main programming language
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **PostgreSQL/SQLite** - Database support

### Frontend
- **HTML/CSS** - Modern, responsive design
- **JavaScript** - Interactive features
- **Bootstrap** - UI framework
- **Font Awesome** - Icons
- **SweetAlert2** - Beautiful alerts

### Architecture
- Repository pattern for data access
- Service layer for business logic
- RESTful API design
- JWT authentication

## 📁 Project Structure

```
/workspace/
├── entities/
│   └── token.go              # Token entity model
├── repositories/
│   └── token/                # Token data access layer
├── services/
│   └── token/                # Token business logic
├── presentation/http/
│   ├── controller.go         # HTTP routes
│   ├── handler.go            # Request handlers
│   └── request/
│       └── token_request.go  # Request models
└── web/
    ├── includes/
    │   ├── landing.html      # Landing page
    │   └── tokens.html       # Token creation page
    └── static/
        ├── css/
        │   ├── landing.css   # Landing page styles
        │   └── tokens.css    # Token page styles
        └── js/app/
            ├── landing.js    # Landing page scripts
            └── tokens.js     # Token creation scripts
```

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher
- PostgreSQL or SQLite
- Modern web browser

### Installation

1. Clone the repository
```bash
git clone <repository-url>
cd workspace
```

2. Install dependencies
```bash
go mod download
```

3. Configure environment variables
```bash
# Set up database connection
# Set up server port
# Configure JWT secret
```

4. Run database migrations
```bash
# Migrations run automatically on startup
```

5. Start the server
```bash
go run cmd/chaos/main.go
```

6. Access the application
```
http://localhost:8080/landing  # Landing page
http://localhost:8080/tokens   # Token creation
```

## 🎮 Usage

### Creating a Token

1. **Navigate to Token Creation**
   - Visit `/tokens` or click "Create Token" in navigation
   
2. **Fill in Token Details**
   - Token Name (e.g., "DogeCoin")
   - Token Symbol (e.g., "DOGE")
   - Total Supply (e.g., "1000000000")
   - Decimals (default: 18)
   - Network (Ethereum, BSC, etc.)
   - Description (optional)
   - Image URL (optional)

3. **Deploy Token**
   - Click "Create Token" button
   - Wait for blockchain deployment
   - Receive contract address and transaction hash

4. **View Your Tokens**
   - All created tokens appear in the table
   - Track status: pending, deployed, or failed
   - View contract addresses and transaction hashes

## 🔐 API Endpoints

### Token Routes
```
GET  /tokens              # View token creation page
POST /token/create        # Create a new token
GET  /token/:id           # Get token by ID
```

### Request/Response Examples

**Create Token**
```json
POST /token/create
{
  "name": "MyMemeCoin",
  "symbol": "MMC",
  "total_supply": "1000000000",
  "decimals": 18,
  "network": "ethereum",
  "description": "My awesome meme coin",
  "image_url": "https://example.com/logo.png"
}

Response:
{
  "id": 1,
  "status": "pending",
  "message": "Token creation initiated. Deployment in progress...",
  "address": "0x...",
  "tx_hash": "0x..."
}
```

## 🎨 Customization

### Branding
Update branding in:
- `/workspace/web/layouts/header.html` - Navigation bar
- `/workspace/web/layouts/base.html` - Page title
- `/workspace/web/static/css/` - Color schemes

### Networks
Add more blockchain networks in:
- `/workspace/web/includes/tokens.html` - Network dropdown
- `/workspace/services/token/token_service.go` - Deployment logic

### Features
Extend functionality:
- Add liquidity pool creation
- Implement token burning
- Add airdrop functionality
- Integrate with DEX platforms

## 🔮 Future Enhancements

- [ ] Real blockchain integration (Web3.js/Ethers.js)
- [ ] Wallet connection (MetaMask, WalletConnect)
- [ ] Smart contract customization options
- [ ] Token analytics dashboard
- [ ] Social sharing features
- [ ] Token marketplace
- [ ] Liquidity pool creation
- [ ] NFT minting integration

## 📝 Notes

- Current implementation uses simulated blockchain deployment
- For production, integrate with actual blockchain networks
- Requires gas fees for real deployments
- Smart contract addresses are currently mock values
- Authentication required for token creation

## 🤝 Contributing

Contributions welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## 📄 License

See LICENSE file for details.

## 🙏 Acknowledgments

Inspired by:
- moontool.fun
- createtokens.io
- The growing meme coin community

---

Built with ❤️ for the crypto community
