# Meme Coin Creator

A comprehensive meme coin creation platform built with Go and modern web technologies. Create, deploy, and manage your own meme coins on multiple blockchain networks.

## Features

### 🚀 Core Features
- **Easy Token Creation**: Simple form-based interface for creating meme coins
- **Multi-Blockchain Support**: Deploy on Ethereum, BSC, Polygon, Arbitrum, and Optimism
- **Real-time Preview**: See your token details before deployment
- **Smart Contract Generation**: Automatic ERC-20 contract generation
- **Social Integration**: Add website, Twitter, Telegram, and Discord links
- **Image Upload**: Custom token images
- **Deployment Management**: Track deployment status and transaction hashes

### 🎨 User Interface
- **Modern Design**: Clean, responsive interface with gradient themes
- **Real-time Updates**: Live preview as you type
- **Mobile Responsive**: Works perfectly on all devices
- **Interactive Elements**: Smooth animations and hover effects
- **Form Validation**: Client and server-side validation

### 🔧 Technical Features
- **RESTful API**: Complete API for all operations
- **Database Integration**: PostgreSQL/SQLite support
- **Authentication**: JWT-based authentication
- **File Upload**: Image upload support
- **Search & Filter**: Advanced search and filtering capabilities
- **Pagination**: Efficient data loading with pagination

## Quick Start

### Prerequisites
- Go 1.22.3 or later
- PostgreSQL or SQLite
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd CHAOS
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   export PORT=8080
   export SQLITE_DATABASE=chaos
   # OR for PostgreSQL:
   export POSTGRES_DATABASE=chaos
   export POSTGRES_HOST=localhost
   export POSTGRES_USER=postgres
   export POSTGRES_PASSWORD=your_password
   export POSTGRES_PORT=5432
   ```

4. **Run the application**
   ```bash
   go run cmd/meme_coin/main.go
   ```

5. **Access the application**
   - Open your browser and go to `http://localhost:8080`
   - Navigate to "Meme Coins" in the menu
   - Click "Create New" to start creating your first meme coin

## API Endpoints

### Public Endpoints
- `GET /meme-coins` - List all meme coins
- `GET /meme-coins/:id` - Get meme coin by ID
- `GET /meme-coins/symbol/:symbol` - Get meme coin by symbol
- `GET /meme-coins/search` - Search meme coins

### Protected Endpoints (Authentication Required)
- `POST /meme-coins` - Create new meme coin
- `GET /meme-coins/my` - Get user's meme coins
- `PUT /meme-coins/:id` - Update meme coin
- `DELETE /meme-coins/:id` - Delete meme coin
- `POST /meme-coins/:id/deploy` - Deploy meme coin
- `POST /meme-coins/preview-contract` - Preview smart contract

### Web Pages
- `/meme-coins` - Browse all meme coins
- `/meme-coins/create` - Create new meme coin
- `/meme-coins/my` - User dashboard
- `/meme-coins/:id` - Meme coin details

## Database Schema

### MemeCoin Entity
```go
type MemeCoin struct {
    ID              string    `json:"id"`
    Name            string    `json:"name"`
    Symbol          string    `json:"symbol"`
    Description     string    `json:"description"`
    TotalSupply     string    `json:"total_supply"`
    Decimals        uint8     `json:"decimals"`
    ImageURL        string    `json:"image_url"`
    Website         string    `json:"website"`
    Twitter         string    `json:"twitter"`
    Telegram        string    `json:"telegram"`
    Discord         string    `json:"discord"`
    ContractAddress string    `json:"contract_address"`
    Network         string    `json:"network"`
    Status          string    `json:"status"`
    DeploymentHash  string    `json:"deployment_hash"`
    DeployedAt      *time.Time `json:"deployed_at"`
    CreatorID       string    `json:"creator_id"`
    Price           float64   `json:"price"`
    IsVerified      bool      `json:"is_verified"`
}
```

## Supported Networks

| Network | Symbol | Deployment Cost | Explorer |
|---------|--------|----------------|----------|
| Ethereum | ETH | 0.05 ETH | [Etherscan](https://etherscan.io) |
| Binance Smart Chain | BNB | 0.01 BNB | [BSCScan](https://bscscan.com) |
| Polygon | MATIC | 0.01 MATIC | [PolygonScan](https://polygonscan.com) |
| Arbitrum | ETH | 0.01 ETH | [Arbiscan](https://arbiscan.io) |
| Optimism | ETH | 0.01 ETH | [Optimistic Etherscan](https://optimistic.etherscan.io) |

## Usage Examples

### Creating a Meme Coin

1. **Fill out the form**:
   - Token Name: "DogeCoin 2.0"
   - Symbol: "DOGE2"
   - Description: "The next generation of meme coins"
   - Total Supply: "1000000000"
   - Decimals: 18
   - Network: "ethereum"
   - Add social links and image

2. **Preview your token**:
   - See real-time preview of your token
   - Review all details before deployment
   - Preview the smart contract code

3. **Deploy to blockchain**:
   - Click "Create Meme Coin"
   - Pay the deployment fee
   - Wait for confirmation
   - Get your contract address

### API Usage

```bash
# Create a new meme coin
curl -X POST http://localhost:8080/meme-coins \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "My Meme Coin",
    "symbol": "MMC",
    "description": "A great meme coin",
    "total_supply": "1000000",
    "decimals": 18,
    "network": "ethereum"
  }'

# Get all meme coins
curl http://localhost:8080/meme-coins

# Search meme coins
curl "http://localhost:8080/meme-coins/search?q=doge&limit=10"
```

## Development

### Project Structure
```
├── cmd/meme_coin/          # Main application entry point
├── entities/               # Database entities
├── repositories/           # Data access layer
├── services/              # Business logic layer
├── presentation/http/     # HTTP handlers and controllers
├── web/                   # Frontend templates and static files
└── infrastructure/        # Database and external services
```

### Adding New Features

1. **Create Entity**: Define your data model in `entities/`
2. **Create Repository**: Implement data access in `repositories/`
3. **Create Service**: Add business logic in `services/`
4. **Create Handler**: Add HTTP handlers in `presentation/http/handler/`
5. **Create Controller**: Wire everything together in `presentation/http/controller/`
6. **Add Routes**: Register routes in the controller
7. **Create Templates**: Add frontend templates in `web/includes/`

### Testing

```bash
# Run tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test ./services/meme_coin
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support and questions:
- Create an issue on GitHub
- Check the documentation
- Review the API endpoints

## Roadmap

### Phase 1 (Current)
- ✅ Basic meme coin creation
- ✅ Multi-blockchain support
- ✅ Web interface
- ✅ API endpoints

### Phase 2 (Planned)
- 🔄 Real blockchain integration
- 🔄 Payment processing
- 🔄 Advanced contract features
- 🔄 Analytics dashboard

### Phase 3 (Future)
- 📋 NFT integration
- 📋 DeFi features
- 📋 Community features
- 📋 Mobile app

---

**Happy Meme Coin Creating! 🚀💰**