// Token creation functionality

function createToken() {
    const tokenData = {
        name: document.getElementById('token-name').value.trim(),
        symbol: document.getElementById('token-symbol').value.trim().toUpperCase(),
        total_supply: document.getElementById('total-supply').value,
        decimals: parseInt(document.getElementById('decimals').value),
        network: document.getElementById('network').value,
        description: document.getElementById('description').value.trim(),
        image_url: document.getElementById('image-url').value.trim()
    };

    // Validation
    if (!tokenData.name || !tokenData.symbol || !tokenData.total_supply || !tokenData.network) {
        Swal.fire({
            icon: 'error',
            title: 'Validation Error',
            text: 'Please fill in all required fields!',
            confirmButtonColor: '#6c5ce7'
        });
        return;
    }

    if (tokenData.symbol.length < 2 || tokenData.symbol.length > 10) {
        Swal.fire({
            icon: 'error',
            title: 'Invalid Symbol',
            text: 'Token symbol must be between 2-10 characters!',
            confirmButtonColor: '#6c5ce7'
        });
        return;
    }

    if (parseInt(tokenData.total_supply) < 1) {
        Swal.fire({
            icon: 'error',
            title: 'Invalid Supply',
            text: 'Total supply must be greater than 0!',
            confirmButtonColor: '#6c5ce7'
        });
        return;
    }

    // Show loading
    Swal.fire({
        title: 'Creating Your Token...',
        html: `
            <div class="token-creation-progress">
                <i class="fas fa-spinner fa-spin fa-3x mb-3"></i>
                <p>Deploying <strong>${tokenData.name}</strong> on ${tokenData.network}...</p>
                <p class="small text-muted">This may take a few moments</p>
            </div>
        `,
        allowOutsideClick: false,
        showConfirmButton: false
    });

    // Send request to create token
    $.ajax({
        url: '/token/create',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(tokenData),
        success: function(response) {
            Swal.fire({
                icon: 'success',
                title: 'Token Created Successfully!',
                html: `
                    <div class="success-message">
                        <p><strong>${tokenData.name} (${tokenData.symbol})</strong> has been created!</p>
                        ${response.address ? `<p class="small">Contract Address: <code>${response.address}</code></p>` : ''}
                        ${response.tx_hash ? `<p class="small">TX Hash: <code>${response.tx_hash}</code></p>` : ''}
                    </div>
                `,
                confirmButtonColor: '#00b894',
                confirmButtonText: 'Awesome!'
            }).then(() => {
                // Reset form
                document.getElementById('token-form').reset();
                // Reload page to show new token
                location.reload();
            });
        },
        error: function(xhr, status, error) {
            let errorMessage = 'An error occurred while creating your token.';
            if (xhr.responseJSON && xhr.responseJSON.error) {
                errorMessage = xhr.responseJSON.error;
            }
            
            Swal.fire({
                icon: 'error',
                title: 'Creation Failed',
                text: errorMessage,
                confirmButtonColor: '#d63031'
            });
        }
    });
}

function viewToken(address) {
    Swal.fire({
        title: 'Token Contract',
        html: `
            <p>Contract Address:</p>
            <code style="word-break: break-all;">${address}</code>
            <div class="mt-3">
                <a href="#" class="btn btn-sm btn-primary" target="_blank">View on Explorer</a>
            </div>
        `,
        confirmButtonColor: '#6c5ce7'
    });
}

// Real-time symbol formatting
document.addEventListener('DOMContentLoaded', function() {
    const symbolInput = document.getElementById('token-symbol');
    if (symbolInput) {
        symbolInput.addEventListener('input', function() {
            this.value = this.value.toUpperCase().replace(/[^A-Z0-9]/g, '');
        });
    }

    // Preview functionality
    const previewElements = {
        name: document.getElementById('token-name'),
        symbol: document.getElementById('token-symbol'),
        supply: document.getElementById('total-supply')
    };

    Object.keys(previewElements).forEach(key => {
        if (previewElements[key]) {
            previewElements[key].addEventListener('input', updatePreview);
        }
    });
});

function updatePreview() {
    // This could be used to show a live preview of the token
    console.log('Preview updated');
}

// Add animation to creation button
document.addEventListener('DOMContentLoaded', function() {
    const createBtn = document.querySelector('.create-btn');
    if (createBtn) {
        createBtn.addEventListener('mouseenter', function() {
            this.style.transform = 'scale(1.05)';
        });
        createBtn.addEventListener('mouseleave', function() {
            this.style.transform = 'scale(1)';
        });
    }
});
