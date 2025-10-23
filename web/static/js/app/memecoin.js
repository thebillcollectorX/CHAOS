(function () {
  const statusEl = document.getElementById('status');
  const connectBtn = document.getElementById('connectWalletBtn');
  const deployBtn = document.getElementById('deployBtn');

  let provider = null;
  let signer = null;
  let currentAccount = null;

  function log(msg) {
    if (!statusEl) return;
    const at = new Date().toISOString().replace('T', ' ').replace('Z', '');
    statusEl.innerHTML += `[${at}] ${msg}<br/>`;
    statusEl.scrollTop = statusEl.scrollHeight;
  }

  async function connect() {
    try {
      if (!window.ethereum) {
        log('No wallet provider found. Install MetaMask or a compatible wallet.');
        return;
      }
      provider = new ethers.BrowserProvider(window.ethereum);
      const accounts = await provider.send('eth_requestAccounts', []);
      signer = await provider.getSigner();
      currentAccount = accounts && accounts[0] ? accounts[0] : (await signer.getAddress());
      connectBtn.textContent = `${currentAccount.slice(0, 6)}...${currentAccount.slice(-4)}`;
      log('Wallet connected.');
    } catch (e) {
      console.error(e);
      log(`Connect error: ${e.message || e}`);
    }
  }

  async function ensureChain(targetChainId) {
    try {
      const network = await provider.getNetwork();
      if (Number(network.chainId) === Number(targetChainId)) return;
      // Try switch
      await window.ethereum.request({
        method: 'wallet_switchEthereumChain',
        params: [{ chainId: '0x' + Number(targetChainId).toString(16) }]
      });
      log(`Switched to chain ${targetChainId}.`);
    } catch (switchErr) {
      log(`Please switch to chain ${targetChainId} in your wallet.`);
      throw switchErr;
    }
  }

  // Minimal ERC-20 (OpenZeppelin-like) bytecode/ABI would normally be compiled.
  // Here we use a simple factory deployment that expects precompiled bytecode.
  // For MVP, we call a third-party deployer or use a minimal template via CREATE2.
  // To keep this demo self-contained, we call a lightweight on-chain factory if present.

  const ERC20_ABI = [
    'function name() view returns (string)',
    'function symbol() view returns (string)',
    'function decimals() view returns (uint8)',
    'function totalSupply() view returns (uint256)',
    'function balanceOf(address owner) view returns (uint256)',
    'event Transfer(address indexed from, address indexed to, uint256 value)'
  ];

  async function deployToken() {
    try {
      if (!provider || !signer) {
        log('Connect wallet first.');
        return;
      }

      const name = document.getElementById('tokenName').value.trim();
      const symbol = document.getElementById('tokenSymbol').value.trim();
      const totalSupply = document.getElementById('totalSupply').value.trim();
      const chainId = document.getElementById('chainSelect').value;
      const mintTo = document.getElementById('mintTo').value.trim();
      const burnable = document.getElementById('burnable').checked;
      const pausable = document.getElementById('pausable').checked;

      if (!name || !symbol || !totalSupply) {
        log('Fill in all required fields.');
        return;
      }

      await ensureChain(chainId);

      // Placeholder deployment strategy:
      // 1) Use a public minimal factory address map per chain (not included here)
      // 2) Or, for demo, send a 0-value tx with data to self to simulate deploy and show steps
      log('Preparing deployment...');

      // In a production build, replace with a real Factory contract + bytecode.
      // For now, we simulate and guide the user.
      log('This demo does not deploy real contracts in this repo.');
      log('Integrate your deployer backend or factory contract here.');

      // Example read-only interaction pattern after real deployment:
      // const token = new ethers.Contract(newTokenAddress, ERC20_ABI, provider);
      // const n = await token.name();

      // UX hint
      log('Next steps: implement a factory or backend deployer.');
    } catch (e) {
      console.error(e);
      log(`Deploy error: ${e.message || e}`);
    }
  }

  if (connectBtn) connectBtn.addEventListener('click', connect);
  if (deployBtn) deployBtn.addEventListener('click', deployToken);
})();
