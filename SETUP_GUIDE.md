# 🚀 MemeToken Creator - Setup and Usage Guide

## Quick Start

### 1. Prerequisites
Ensure you have the following installed:
- Go 1.16+ (`go version`)
- PostgreSQL or SQLite database
- Git

### 2. Environment Setup

Create a `.env` file or set environment variables:

```bash
# Server Configuration
SERVER_PORT=8080

# Database Configuration (choose one)
# For SQLite:
SQLITE_DATABASE_NAME=chaos

# For PostgreSQL:
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DATABASE=chaos
POSTGRES_SSL_MODE=disable

# JWT Secret Key
SECRET_KEY=your-secret-key-here
```

### 3. Installation

```bash
# Clone the repository
git clone <repository-url>
cd workspace

# Install dependencies
go mod download

# Build the application
go build -o chaos ./cmd/chaos/main.go

# Run the application
./chaos
```

### 4. First Time Setup

When you first run the application:
1. A default admin user will be created
2. Database tables will be automatically migrated
3. The server will start on the configured port (default: 8080)

### 5. Access the Application

Open your browser and navigate to:

- **Landing Page**: `http://localhost:8080/landing`
- **Login Page**: `http://localhost:8080/login`
- **Token Creator**: `http://localhost:8080/tokens` (requires login)

**Default Login Credentials:**
- Username: `admin`
- Password: (check console output or configuration)

## Features Overview

### 🎨 Landing Page (`/landing`)
A beautiful, modern landing page featuring:
- Hero section with animated statistics
- Feature highlights
- Step-by-step guide
- Supported blockchain networks
- Call-to-action buttons

**Technologies:**
- Gradient backgrounds
- Floating animations
- Responsive design
- Modern UI/UX

### 🪙 Token Creation Page (`/tokens`)
Comprehensive token creation interface:

**Form Fields:**
- Token Name (e.g., "DogeCoin")
- Symbol (2-10 characters, e.g., "DOGE")
- Total Supply (numeric)
- Decimals (0, 6, 9, or 18)
- Blockchain Network (Ethereum, BSC, Polygon, Solana, Base)
- Description (optional)
- Image URL (optional)

**Features:**
- Real-time form validation
- Beautiful alerts using SweetAlert2
- Deployment status tracking
- Recently created tokens table
- Contract address and transaction hash display

## API Endpoints

### Authentication
```
POST /auth
GET  /logout
```

### Token Management
```
GET  /tokens              # View token creation page (requires auth)
POST /token/create        # Create a new token (requires auth)
GET  /token/:id           # Get token by ID (requires auth)
```

### Example: Create Token

**Request:**
```bash
curl -X POST http://localhost:8080/token/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "MoonCoin",
    "symbol": "MOON",
    "total_supply": "1000000000",
    "decimals": 18,
    "network": "ethereum",
    "description": "A meme coin that goes to the moon!",
    "image_url": "https://example.com/moon.png"
  }'
```

**Response:**
```json
{
  "id": 1,
  "status": "pending",
  "message": "Token creation initiated. Deployment in progress...",
  "address": "0x1234567890abcdef...",
  "tx_hash": "0xabcdef1234567890..."
}
```

## Database Schema

### Tokens Table
```sql
CREATE TABLE v1_0_tokens (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    total_supply VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(500),
    network VARCHAR(50) NOT NULL,
    user_id INTEGER,
    decimals INTEGER,
    address VARCHAR(255),
    status VARCHAR(50),
    tx_hash VARCHAR(255)
);
```

## Development

### Project Structure
```
/workspace/
├── cmd/chaos/main.go              # Application entry point
├── entities/token.go              # Token data model
├── repositories/token/            # Data access layer
│   ├── token.go                   # Repository interface
│   └── token_repository.go        # Repository implementation
├── services/token/                # Business logic layer
│   ├── token.go                   # Service interface
│   └── token_service.go           # Service implementation
├── presentation/http/             # HTTP layer
│   ├── controller.go              # Route definitions
│   ├── handler.go                 # Request handlers
│   └── request/token_request.go   # Request models
└── web/                           # Frontend files
    ├── includes/
    │   ├── landing.html           # Landing page
    │   └── tokens.html            # Token creation page
    └── static/
        ├── css/
        │   ├── landing.css        # Landing page styles
        │   └── tokens.css         # Token page styles
        └── js/app/
            ├── landing.js         # Landing page scripts
            └── tokens.js          # Token creation scripts
```

### Adding New Features

1. **Add New Entity Field:**
   - Update `entities/token.go`
   - Run the application (auto-migration will update the database)

2. **Add New Route:**
   - Add route in `presentation/http/controller.go`
   - Add handler in `presentation/http/handler.go`

3. **Add New Blockchain Network:**
   - Update network dropdown in `web/includes/tokens.html`
   - Update service logic in `services/token/token_service.go`

## Customization

### Change Branding
Update these files:
- `web/layouts/header.html` - Navigation bar and logo
- `web/layouts/base.html` - Page title and meta tags
- `web/static/css/tokens.css` - Color scheme
- `web/static/css/landing.css` - Landing page colors

### Change Color Scheme
Edit the gradient colors in CSS files:
```css
/* Current gradient */
background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);

/* Example alternative */
background: linear-gradient(135deg, #00b894 0%, #00cec9 100%);
```

## Production Deployment

### Building for Production
```bash
# Build optimized binary
CGO_ENABLED=1 go build -ldflags="-s -w" -o chaos ./cmd/chaos/main.go

# Set production environment variables
export SERVER_PORT=80
export POSTGRES_HOST=your-prod-db-host
export SECRET_KEY=your-secure-secret-key

# Run the application
./chaos
```

### Docker Deployment
The project includes a `Dockerfile`:
```bash
docker build -t memetoken-creator .
docker run -p 8080:8080 -e POSTGRES_HOST=db memetoken-creator
```

### Heroku Deployment
The project includes `heroku.yml` and `Procfile`:
```bash
heroku create your-app-name
git push heroku main
```

## Security Considerations

⚠️ **Important for Production:**

1. **Change Default Credentials**: Update admin password
2. **Use Strong Secret Key**: Generate a secure JWT secret
3. **Enable HTTPS**: Use SSL/TLS certificates
4. **Database Security**: Use strong passwords and proper access controls
5. **Rate Limiting**: Implement rate limiting for API endpoints
6. **Input Validation**: Always validate and sanitize user input
7. **CORS Configuration**: Configure CORS appropriately

## Troubleshooting

### Build Errors
```bash
# Clean and rebuild
go clean
go mod tidy
go build ./cmd/chaos/main.go
```

### Database Connection Issues
```bash
# Check PostgreSQL is running
sudo service postgresql status

# Test connection
psql -h localhost -U your_user -d chaos
```

### Port Already in Use
```bash
# Find process using port 8080
lsof -i :8080
# or
netstat -tulpn | grep 8080

# Kill process or change SERVER_PORT
```

## Next Steps

1. **Integrate Real Blockchain**: Connect to actual blockchain networks using Web3
2. **Add Wallet Connection**: Implement MetaMask/WalletConnect integration
3. **Smart Contract Templates**: Add customizable smart contract options
4. **Token Analytics**: Add charts and statistics for created tokens
5. **Liquidity Pools**: Integrate DEX liquidity pool creation
6. **Social Features**: Add sharing and community features

## Support

For issues and questions:
- Check the main README: `TOKEN_CREATOR_README.md`
- Review the CONTRIBUTING guide: `CONTRIBUTING.md`
- Submit issues on GitHub

---

Happy Token Creating! 🚀🌙
