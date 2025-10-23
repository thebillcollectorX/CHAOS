#!/bin/bash

# Meme Coin Creator Setup Script
echo "🚀 Setting up Meme Coin Creator..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.22.3 or later."
    echo "Visit: https://golang.org/dl/"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✅ Go version: $GO_VERSION"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy

# Create necessary directories
echo "📁 Creating directories..."
mkdir -p web/static/images
mkdir -p logs

# Set up environment variables
echo "⚙️  Setting up environment..."
if [ ! -f .env ]; then
    cat > .env << EOF
# Server Configuration
PORT=8080

# Database Configuration (choose one)
# SQLite (for development)
SQLITE_DATABASE=chaos

# PostgreSQL (for production)
# POSTGRES_DATABASE=chaos
# POSTGRES_HOST=localhost
# POSTGRES_USER=postgres
# POSTGRES_PASSWORD=your_password
# POSTGRES_PORT=5432
# POSTGRES_SSL_MODE=disable

# JWT Configuration
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRE_HOURS=24

# Blockchain Configuration
ETHEREUM_RPC_URL=https://mainnet.infura.io/v3/YOUR_PROJECT_ID
BSC_RPC_URL=https://bsc-dataseed.binance.org/
POLYGON_RPC_URL=https://polygon-rpc.com/
ARBITRUM_RPC_URL=https://arb1.arbitrum.io/rpc
OPTIMISM_RPC_URL=https://mainnet.optimism.io

# Payment Configuration
STRIPE_SECRET_KEY=your_stripe_secret_key
STRIPE_PUBLISHABLE_KEY=your_stripe_publishable_key
PAYPAL_CLIENT_ID=your_paypal_client_id
PAYPAL_CLIENT_SECRET=your_paypal_client_secret
EOF
    echo "✅ Created .env file. Please update with your actual values."
else
    echo "✅ .env file already exists."
fi

# Build the application
echo "🔨 Building application..."
go build -o bin/meme-coin-creator cmd/meme_coin/main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

# Create systemd service file (optional)
if [ "$1" = "--service" ]; then
    echo "🔧 Creating systemd service..."
    sudo tee /etc/systemd/system/meme-coin-creator.service > /dev/null << EOF
[Unit]
Description=Meme Coin Creator
After=network.target

[Service]
Type=simple
User=$USER
WorkingDirectory=$(pwd)
ExecStart=$(pwd)/bin/meme-coin-creator
Restart=always
RestartSec=5
Environment=PORT=8080
Environment=SQLITE_DATABASE=chaos

[Install]
WantedBy=multi-user.target
EOF

    echo "✅ Systemd service created. Run 'sudo systemctl enable meme-coin-creator' to enable it."
fi

echo ""
echo "🎉 Setup complete!"
echo ""
echo "To start the application:"
echo "  ./bin/meme-coin-creator"
echo ""
echo "Or run directly:"
echo "  go run cmd/meme_coin/main.go"
echo ""
echo "The application will be available at: http://localhost:8080"
echo ""
echo "Next steps:"
echo "1. Update the .env file with your actual configuration"
echo "2. Set up your blockchain RPC URLs"
echo "3. Configure payment providers if needed"
echo "4. Start the application and create your first meme coin!"
echo ""
echo "For more information, see MEME_COIN_README.md"