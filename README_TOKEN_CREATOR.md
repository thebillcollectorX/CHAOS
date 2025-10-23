# Token Creator - Meme Coin Creation Platform

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22+-blue.svg?style=flat-square" alt="Go Version">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License">
  <img src="https://img.shields.io/badge/Web3-Enabled-green.svg?style=flat-square" alt="Web3">
</p>

A comprehensive platform for creating and deploying meme tokens with zero coding knowledge. Built with Go, Gin, GORM, and Web3 integration.

## 🚀 Features

### Token Creation
- **No Coding Required**: Intuitive drag-and-drop interface
- **Multi-Chain Support**: Deploy on Ethereum, BSC, Polygon, Avalanche
- **Advanced Features**: Tax systems, anti-whale protection, burning mechanisms
- **Real-time Preview**: See your token contract before deployment

### Wallet Integration
- **MetaMask Support**: Seamless wallet connection
- **Multi-Network**: Automatic network switching
- **Gas Estimation**: Real-time deployment cost calculation
- **Secure Transactions**: All transactions signed in your wallet

### Token Management
- **Dashboard**: Comprehensive token management interface
- **Analytics**: Real-time holder count, transactions, market cap
- **Social Integration**: Add website, Twitter, Telegram links
- **Contract Verification**: Automatic contract verification on explorers

### Advanced Features
- **Tax System**: Configurable buy/sell taxes (0-25%)
- **Anti-Whale Protection**: Maximum transaction and wallet limits
- **Mint/Burn**: Optional minting and burning capabilities
- **Pausable**: Emergency pause functionality
- **Ownership**: Renounce or transfer ownership

## 🛠 Tech Stack

- **Backend**: Go 1.22+, Gin Web Framework
- **Database**: PostgreSQL/SQLite with GORM
- **Frontend**: HTML5, Bootstrap 5, JavaScript
- **Blockchain**: Web3.js, Ethereum-compatible networks
- **Authentication**: JWT tokens
- **Smart Contracts**: OpenZeppelin standards

## 📋 Prerequisites

- Go 1.22 or higher
- Node.js (for Web3 dependencies)
- PostgreSQL or SQLite
- MetaMask browser extension (for users)

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/token-creator
cd token-creator
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Set Environment Variables
```bash
# For SQLite (Development)
export SQLITE_DATABASE=token_creator
export PORT=8080
export SECRET_KEY=your-secret-key

# For PostgreSQL (Production)
export POSTGRES_DATABASE=token_creator
export POSTGRES_HOST=localhost
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=your-password
export POSTGRES_PORT=5432
export PORT=8080
export SECRET_KEY=your-secret-key
```

### 4. Run the Application
```bash
# Development
go run cmd/chaos/main.go

# Production
go build -o token-creator cmd/chaos/main.go
./token-creator
```

### 5. Access the Platform
Open your browser and navigate to `http://localhost:8080`

## 📖 Usage Guide

### Creating Your First Token

1. **Connect Wallet**: Click "Connect Wallet" and approve MetaMask connection
2. **Fill Token Details**:
   - Token Name (e.g., "Doge Coin")
   - Symbol (e.g., "DOGE")
   - Total Supply
   - Description (optional)

3. **Choose Features**:
   - Mintable: Allow creating new tokens
   - Burnable: Allow destroying tokens
   - Pausable: Emergency pause capability
   - Tax System: Buy/sell transaction fees
   - Anti-Whale: Maximum transaction limits

4. **Select Network**: Choose deployment blockchain
5. **Deploy**: Review gas costs and deploy your token

### Managing Tokens

- **Dashboard**: View all your created tokens
- **Analytics**: Track holders, transactions, market performance
- **Edit**: Update token information and social links
- **Deploy**: Deploy draft tokens to blockchain

### Supported Networks

| Network | Chain ID | Currency | Status |
|---------|----------|----------|--------|
| Ethereum | 1 | ETH | ✅ Active |
| BSC | 56 | BNB | ✅ Active |
| Polygon | 137 | MATIC | ✅ Active |
| Avalanche | 43114 | AVAX | ✅ Active |
| Goerli (Testnet) | 5 | ETH | ✅ Active |
| BSC Testnet | 97 | BNB | ✅ Active |

## 🏗 Architecture

```
├── cmd/chaos/           # Application entry point
├── entities/            # Database models
├── repositories/        # Data access layer
├── services/           # Business logic
├── presentation/http/  # HTTP handlers and routes
├── internal/           # Internal utilities
├── web/               # Frontend templates and assets
└── infrastructure/    # Database and external services
```

### Database Schema

- **Users**: User accounts and authentication
- **Tokens**: Token information and metadata
- **TokenFeatures**: Advanced token features configuration
- **TokenAnalytics**: Real-time token statistics
- **Networks**: Supported blockchain networks
- **Wallets**: Connected wallet addresses
- **Transactions**: Deployment and interaction history
- **ContractTemplates**: Smart contract templates

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `SECRET_KEY` | JWT secret key | Required |
| `SQLITE_DATABASE` | SQLite database name | Optional |
| `POSTGRES_*` | PostgreSQL configuration | Optional |

### Network Configuration

Networks are automatically seeded on first run. You can add custom networks through the database or admin interface.

## 🚀 Deployment

### Docker Deployment

```bash
# Build image
docker build -t token-creator .

# Run container
docker run -d \
  -p 8080:8080 \
  -e PORT=8080 \
  -e SECRET_KEY=your-secret-key \
  -e SQLITE_DATABASE=token_creator \
  token-creator
```

### Heroku Deployment

```bash
# Create Heroku app
heroku create your-app-name

# Set environment variables
heroku config:set SECRET_KEY=your-secret-key
heroku config:set POSTGRES_SSL_MODE=require

# Deploy
git push heroku main
```

## 🔐 Security Features

- **JWT Authentication**: Secure user sessions
- **Input Validation**: Comprehensive request validation
- **SQL Injection Protection**: GORM ORM with prepared statements
- **XSS Protection**: Template escaping and CSP headers
- **Rate Limiting**: API endpoint protection
- **Wallet Security**: All transactions signed client-side

## 📊 Analytics & Monitoring

- **Real-time Metrics**: Token holders, transactions, volume
- **Performance Tracking**: Gas usage, deployment success rates
- **User Analytics**: Token creation patterns, popular features
- **Network Statistics**: Cross-chain deployment distribution

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Write comprehensive tests for new features
- Update documentation for API changes
- Ensure cross-browser compatibility for frontend changes

## 📝 API Documentation

### Authentication
```bash
POST /auth
Content-Type: application/json

{
  "username": "admin",
  "password": "admin"
}
```

### Create Token
```bash
POST /api/tokens
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "My Token",
  "symbol": "MTK",
  "total_supply": "1000000",
  "network": "ethereum",
  "is_mintable": true,
  "is_burnable": false
}
```

### Get User Tokens
```bash
GET /api/user/tokens
Authorization: Bearer <token>
```

### Deploy Token
```bash
POST /api/tokens/deploy
Authorization: Bearer <token>
Content-Type: application/json

{
  "token_id": 1,
  "network_id": 1
}
```

## 🐛 Troubleshooting

### Common Issues

**MetaMask Connection Failed**
- Ensure MetaMask is installed and unlocked
- Check if the website is allowed in MetaMask
- Try refreshing the page and reconnecting

**Transaction Failed**
- Check if you have sufficient gas fees
- Verify the network is correct
- Ensure token parameters are valid

**Database Connection Error**
- Verify database credentials
- Check if database server is running
- Ensure proper environment variables are set

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [OpenZeppelin](https://openzeppelin.com/) for secure smart contract libraries
- [Gin](https://gin-gonic.com/) for the excellent web framework
- [GORM](https://gorm.io/) for the powerful ORM
- [Web3.js](https://web3js.readthedocs.io/) for blockchain interaction
- [Bootstrap](https://getbootstrap.com/) for responsive UI components

## 📞 Support

- 📧 Email: support@tokencreator.com
- 💬 Discord: [Join our community](https://discord.gg/tokencreator)
- 📖 Documentation: [docs.tokencreator.com](https://docs.tokencreator.com)
- 🐛 Issues: [GitHub Issues](https://github.com/your-username/token-creator/issues)

---

<p align="center">
  Made with ❤️ by the Token Creator Team
</p>