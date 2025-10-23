// Wallet Integration for Token Creator
class WalletManager {
    constructor() {
        this.web3 = null;
        this.account = null;
        this.chainId = null;
        this.provider = null;
        this.isConnected = false;
        
        this.supportedNetworks = {
            1: { name: 'Ethereum', currency: 'ETH', rpcUrl: 'https://mainnet.infura.io/v3/' },
            56: { name: 'BSC', currency: 'BNB', rpcUrl: 'https://bsc-dataseed.binance.org/' },
            137: { name: 'Polygon', currency: 'MATIC', rpcUrl: 'https://polygon-rpc.com/' },
            43114: { name: 'Avalanche', currency: 'AVAX', rpcUrl: 'https://api.avax.network/ext/bc/C/rpc' }
        };
        
        this.init();
    }
    
    async init() {
        // Check if wallet is already connected
        if (typeof window.ethereum !== 'undefined') {
            this.provider = window.ethereum;
            this.web3 = new Web3(window.ethereum);
            
            // Check if already connected
            const accounts = await this.web3.eth.getAccounts();
            if (accounts.length > 0) {
                this.account = accounts[0];
                this.chainId = await this.web3.eth.getChainId();
                this.isConnected = true;
                this.updateUI();
            }
            
            // Listen for account changes
            window.ethereum.on('accountsChanged', (accounts) => {
                if (accounts.length === 0) {
                    this.disconnect();
                } else {
                    this.account = accounts[0];
                    this.updateUI();
                }
            });
            
            // Listen for chain changes
            window.ethereum.on('chainChanged', (chainId) => {
                this.chainId = parseInt(chainId, 16);
                this.updateUI();
                window.location.reload(); // Reload to update network-specific data
            });
        }
    }
    
    async connectWallet() {
        try {
            if (typeof window.ethereum === 'undefined') {
                throw new Error('MetaMask is not installed. Please install MetaMask to continue.');
            }
            
            // Request account access
            const accounts = await window.ethereum.request({
                method: 'eth_requestAccounts'
            });
            
            if (accounts.length === 0) {
                throw new Error('No accounts found. Please make sure MetaMask is unlocked.');
            }
            
            this.account = accounts[0];
            this.chainId = await this.web3.eth.getChainId();
            this.isConnected = true;
            
            // Save wallet connection to backend
            await this.saveWalletConnection();
            
            this.updateUI();
            this.showNotification('Wallet connected successfully!', 'success');
            
            return {
                address: this.account,
                chainId: this.chainId,
                network: this.getNetworkName(this.chainId)
            };
            
        } catch (error) {
            console.error('Error connecting wallet:', error);
            this.showNotification(error.message || 'Failed to connect wallet', 'error');
            throw error;
        }
    }
    
    async saveWalletConnection() {
        try {
            const response = await fetch('/api/wallet/connect', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem('token')
                },
                body: JSON.stringify({
                    address: this.account,
                    type: 'metamask'
                })
            });
            
            if (!response.ok) {
                throw new Error('Failed to save wallet connection');
            }
        } catch (error) {
            console.error('Error saving wallet connection:', error);
        }
    }
    
    disconnect() {
        this.account = null;
        this.chainId = null;
        this.isConnected = false;
        this.updateUI();
        this.showNotification('Wallet disconnected', 'info');
    }
    
    async switchNetwork(chainId) {
        try {
            await window.ethereum.request({
                method: 'wallet_switchEthereumChain',
                params: [{ chainId: '0x' + chainId.toString(16) }],
            });
        } catch (switchError) {
            // This error code indicates that the chain has not been added to MetaMask
            if (switchError.code === 4902) {
                await this.addNetwork(chainId);
            } else {
                throw switchError;
            }
        }
    }
    
    async addNetwork(chainId) {
        const network = this.supportedNetworks[chainId];
        if (!network) {
            throw new Error('Unsupported network');
        }
        
        try {
            await window.ethereum.request({
                method: 'wallet_addEthereumChain',
                params: [{
                    chainId: '0x' + chainId.toString(16),
                    chainName: network.name,
                    nativeCurrency: {
                        name: network.currency,
                        symbol: network.currency,
                        decimals: 18
                    },
                    rpcUrls: [network.rpcUrl],
                    blockExplorerUrls: [this.getExplorerUrl(chainId)]
                }]
            });
        } catch (addError) {
            throw new Error('Failed to add network to MetaMask');
        }
    }
    
    getNetworkName(chainId) {
        return this.supportedNetworks[chainId]?.name || 'Unknown Network';
    }
    
    getExplorerUrl(chainId) {
        const explorers = {
            1: 'https://etherscan.io',
            56: 'https://bscscan.com',
            137: 'https://polygonscan.com',
            43114: 'https://snowtrace.io'
        };
        return explorers[chainId] || '';
    }
    
    async getBalance() {
        if (!this.isConnected) return '0';
        
        try {
            const balance = await this.web3.eth.getBalance(this.account);
            return this.web3.utils.fromWei(balance, 'ether');
        } catch (error) {
            console.error('Error getting balance:', error);
            return '0';
        }
    }
    
    async estimateGas(contractData, value = '0') {
        if (!this.isConnected) throw new Error('Wallet not connected');
        
        try {
            const gasEstimate = await this.web3.eth.estimateGas({
                from: this.account,
                data: contractData,
                value: this.web3.utils.toWei(value, 'ether')
            });
            
            const gasPrice = await this.web3.eth.getGasPrice();
            const gasCost = this.web3.utils.fromWei((BigInt(gasEstimate) * BigInt(gasPrice)).toString(), 'ether');
            
            return {
                gasLimit: gasEstimate,
                gasPrice: gasPrice,
                gasCost: gasCost
            };
        } catch (error) {
            console.error('Error estimating gas:', error);
            throw error;
        }
    }
    
    async deployContract(bytecode, abi, constructorParams = []) {
        if (!this.isConnected) throw new Error('Wallet not connected');
        
        try {
            const contract = new this.web3.eth.Contract(abi);
            
            const deployTransaction = contract.deploy({
                data: bytecode,
                arguments: constructorParams
            });
            
            const gasEstimate = await deployTransaction.estimateGas({ from: this.account });
            const gasPrice = await this.web3.eth.getGasPrice();
            
            const result = await deployTransaction.send({
                from: this.account,
                gas: gasEstimate,
                gasPrice: gasPrice
            });
            
            return {
                contractAddress: result.options.address,
                transactionHash: result.transactionHash,
                gasUsed: result.gasUsed
            };
            
        } catch (error) {
            console.error('Error deploying contract:', error);
            throw error;
        }
    }
    
    formatAddress(address) {
        if (!address) return '';
        return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
    }
    
    updateUI() {
        const connectBtn = document.getElementById('connectWalletBtn');
        const walletInfo = document.getElementById('walletInfo');
        const walletAddress = document.getElementById('walletAddress');
        const walletNetwork = document.getElementById('walletNetwork');
        const walletBalance = document.getElementById('walletBalance');
        
        if (this.isConnected) {
            if (connectBtn) {
                connectBtn.style.display = 'none';
            }
            
            if (walletInfo) {
                walletInfo.style.display = 'block';
            }
            
            if (walletAddress) {
                walletAddress.textContent = this.formatAddress(this.account);
                walletAddress.title = this.account;
            }
            
            if (walletNetwork) {
                walletNetwork.textContent = this.getNetworkName(this.chainId);
            }
            
            if (walletBalance) {
                this.getBalance().then(balance => {
                    walletBalance.textContent = parseFloat(balance).toFixed(4);
                });
            }
            
            // Enable deploy buttons
            document.querySelectorAll('.deploy-btn').forEach(btn => {
                btn.disabled = false;
                btn.textContent = 'Deploy Token';
            });
            
        } else {
            if (connectBtn) {
                connectBtn.style.display = 'block';
            }
            
            if (walletInfo) {
                walletInfo.style.display = 'none';
            }
            
            // Disable deploy buttons
            document.querySelectorAll('.deploy-btn').forEach(btn => {
                btn.disabled = true;
                btn.textContent = 'Connect Wallet to Deploy';
            });
        }
    }
    
    showNotification(message, type = 'info') {
        // Create notification element
        const notification = document.createElement('div');
        notification.className = `alert alert-${type === 'error' ? 'danger' : type} alert-dismissible fade show position-fixed`;
        notification.style.cssText = 'top: 20px; right: 20px; z-index: 9999; min-width: 300px;';
        notification.innerHTML = `
            ${message}
            <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
        `;
        
        document.body.appendChild(notification);
        
        // Auto remove after 5 seconds
        setTimeout(() => {
            if (notification.parentNode) {
                notification.parentNode.removeChild(notification);
            }
        }, 5000);
    }
}

// Initialize wallet manager
let walletManager;

document.addEventListener('DOMContentLoaded', function() {
    walletManager = new WalletManager();
    
    // Connect wallet button handler
    const connectBtn = document.getElementById('connectWalletBtn');
    if (connectBtn) {
        connectBtn.addEventListener('click', async () => {
            try {
                connectBtn.disabled = true;
                connectBtn.innerHTML = '<i class="fas fa-spinner fa-spin"></i> Connecting...';
                
                await walletManager.connectWallet();
                
            } catch (error) {
                console.error('Connection failed:', error);
            } finally {
                connectBtn.disabled = false;
                connectBtn.innerHTML = '<i class="fas fa-wallet"></i> Connect Wallet';
            }
        });
    }
    
    // Disconnect wallet button handler
    const disconnectBtn = document.getElementById('disconnectWalletBtn');
    if (disconnectBtn) {
        disconnectBtn.addEventListener('click', () => {
            walletManager.disconnect();
        });
    }
    
    // Network switch handlers
    document.querySelectorAll('.network-switch-btn').forEach(btn => {
        btn.addEventListener('click', async () => {
            const chainId = parseInt(btn.dataset.chainId);
            try {
                await walletManager.switchNetwork(chainId);
            } catch (error) {
                walletManager.showNotification('Failed to switch network: ' + error.message, 'error');
            }
        });
    });
});

// Export for global use
window.walletManager = walletManager;