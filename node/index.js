import express from 'express';
import TronWeb from 'tronweb';
import cors from 'cors';

// Initialize Express app
const app = express();
app.use(cors())

// Middleware to parse JSON bodies
app.use(express.json());

// POST route for TRX transaction
app.post('/api/sendTrx', async (req, res) => {
    console.log('Received request:', req.body);
    try {
        // Extract data from request body including privateKey
        console.log('Request body:', req.body);
        const { senderAddress, receiverAddress, amount, privateKey } = req.body;

        // Validate inputs
        if (!senderAddress || !receiverAddress || !amount || !privateKey) {
            return res.status(400).json({
                error: 'Missing required fields: senderAddress, receiverAddress, amount, and privateKey are required hi'
            });
        }

        // TronWeb configuration for Shasta testnet
        const HttpProvider = TronWeb.providers.HttpProvider;
        const fullNode = new HttpProvider("https://api.shasta.trongrid.io");
        const solidityNode = new HttpProvider("https://api.shasta.trongrid.io");
        const eventServer = new HttpProvider("https://api.shasta.trongrid.io");

        // Initialize TronWeb with the provided private key
        const tronWeb = new TronWeb(fullNode, solidityNode, eventServer, privateKey);

        // Validate Base58 addresses
        if (!tronWeb.isAddress(senderAddress) || !tronWeb.isAddress(receiverAddress)) {
            return res.status(400).json({
                error: 'Invalid Base58 address format'
            });
        }

        // Verify the private key matches the sender address
        const derivedAddress = tronWeb.address.fromPrivateKey(privateKey);
        if (derivedAddress !== senderAddress) {
            return res.status(400).json({
                error: 'Private key does not match sender address'
            });
        }

        // Convert amount to SUN (1 TRX = 1,000,000 SUN)
        const amountInSun = parseInt(amount * 1000000);

        // Create unsigned transaction
        const transaction = await tronWeb.transactionBuilder.sendTrx(
            receiverAddress,
            amountInSun,
            senderAddress
        );

        // Sign the transaction
        const signedTransaction = await tronWeb.trx.sign(transaction);

        // Broadcast the transaction
        const result = await tronWeb.trx.sendRawTransaction(signedTransaction);

        // Send success response
        res.status(200).json({
            success: true,
            transactionId: result.txid,
        });

    } catch (error) {
        console.error('Transaction error:', error);
        res.status(500).json({
            success: false,
            error: error.message || 'Failed to process transaction'
        });
    }
});

// Start the server
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});