(function () {
  const statusEl = document.getElementById('status');
  const resultEl = document.getElementById('result');
  const connectBtn = document.getElementById('connectBtn');
  const deployBtn = document.getElementById('deployBtn');

  const setStatus = (text) => {
    if (statusEl) statusEl.textContent = text || '';
  };
  const showResult = (text, type) => {
    if (!resultEl) return;
    resultEl.style.display = 'block';
    resultEl.className = 'alert alert-' + (type || 'info');
    resultEl.textContent = text;
  };

  function requireMetaMask() {
    if (!window.ethereum) {
      showResult('No wallet found. Install MetaMask or a compatible wallet.', 'warning');
      throw new Error('No wallet');
    }
  }

  async function connectWallet() {
    try {
      requireMetaMask();
      setStatus('Connecting wallet...');
      const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
      const account = accounts && accounts[0];
      if (!account) throw new Error('No account returned');
      setStatus('Connected: ' + account.slice(0, 6) + '...' + account.slice(-4));
      showResult('Wallet connected', 'success');
    } catch (err) {
      console.error(err);
      setStatus('');
      showResult(err.message || String(err), 'danger');
    }
  }

  function getSource(name, symbol, decimals) {
    return `// SPDX-License-Identifier: MIT\npragma solidity ^0.8.20;\n\ncontract MemeERC20 {\n    string public name;\n    string public symbol;\n    uint8 public decimals;\n    uint256 public totalSupply;\n    mapping(address => uint256) public balanceOf;\n    mapping(address => mapping(address => uint256)) public allowance;\n\n    event Transfer(address indexed from, address indexed to, uint256 value);\n    event Approval(address indexed owner, address indexed spender, uint256 value);\n\n    constructor(string memory _name, string memory _symbol, uint8 _decimals, uint256 _initialSupply) {\n        name = _name;\n        symbol = _symbol;\n        decimals = _decimals;\n        uint256 supply = _initialSupply * (10 ** uint256(_decimals));\n        totalSupply = supply;\n        balanceOf[msg.sender] = supply;\n        emit Transfer(address(0), msg.sender, supply);\n    }\n\n    function _transfer(address from, address to, uint256 value) internal {\n        require(to != address(0), 'zero address');\n        uint256 fromBal = balanceOf[from];\n        require(fromBal >= value, 'insufficient');\n        unchecked { balanceOf[from] = fromBal - value; }\n        balanceOf[to] += value;\n        emit Transfer(from, to, value);\n    }\n\n    function transfer(address to, uint256 value) external returns (bool) {\n        _transfer(msg.sender, to, value);\n        return true;\n    }\n\n    function approve(address spender, uint256 value) external returns (bool) {\n        allowance[msg.sender][spender] = value;\n        emit Approval(msg.sender, spender, value);\n        return true;\n    }\n\n    function transferFrom(address from, address to, uint256 value) external returns (bool) {\n        uint256 allowed = allowance[from][msg.sender];\n        require(allowed >= value, 'not allowed');\n        if (allowed != type(uint256).max) {\n            allowance[from][msg.sender] = allowed - value;\n        }\n        _transfer(from, to, value);\n        return true;\n    }\n}`;
  }

  function getSolcInput(source) {
    return {
      language: 'Solidity',
      sources: { 'MemeERC20.sol': { content: source } },
      settings: {
        optimizer: { enabled: true, runs: 200 },
        outputSelection: { '*': { '*': ['abi', 'evm.bytecode'] } }
      }
    };
  }

  async function compile(source) {
    if (typeof Module === 'undefined' || typeof window.solc === 'undefined') {
      throw new Error('Solidity compiler not loaded');
    }
    const input = getSolcInput(source);
    const output = JSON.parse(window.solc.compile(JSON.stringify(input)));
    if (output.errors && output.errors.some(e => e.severity === 'error')) {
      const msg = output.errors.map(e => e.formattedMessage || e.message).join('\n');
      throw new Error(msg);
    }
    const contract = output.contracts['MemeERC20.sol']['MemeERC20'];
    return { abi: contract.abi, bytecode: '0x' + contract.evm.bytecode.object };
  }

  async function deploy() {
    try {
      requireMetaMask();
      const name = document.getElementById('name').value.trim();
      const symbol = document.getElementById('symbol').value.trim();
      const supplyStr = document.getElementById('supply').value.trim();
      const decimalsStr = document.getElementById('decimals').value.trim();
      const decimals = parseInt(decimalsStr, 10);
      const initialSupply = ethers.BigNumber.from(supplyStr || '0');
      if (!name || !symbol || isNaN(decimals) || decimals < 0 || decimals > 18) {
        throw new Error('Fill all required fields correctly');
      }
      setStatus('Compiling Solidity...');
      const src = getSource(name, symbol, decimals);
      const { abi, bytecode } = await compile(src);

      setStatus('Awaiting wallet confirmation...');
      const provider = new ethers.providers.Web3Provider(window.ethereum);
      await provider.send('eth_requestAccounts', []);
      const signer = provider.getSigner();

      const factory = new ethers.ContractFactory(abi, bytecode, signer);
      setStatus('Sending deployment transaction...');
      const contract = await factory.deploy(name, symbol, decimals, initialSupply);
      showResult('Transaction sent: ' + contract.deployTransaction.hash, 'info');
      setStatus('Waiting for confirmation...');
      await contract.deployed();

      const addr = contract.address;
      setStatus('Deployed');
      showResult('Contract deployed at ' + addr, 'success');
    } catch (err) {
      console.error(err);
      setStatus('');
      showResult(err.message || String(err), 'danger');
    }
  }

  connectBtn && connectBtn.addEventListener('click', connectWallet);
  deployBtn && deployBtn.addEventListener('click', deploy);
})();
