## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Charles-Leo-G/tss-lib
cd tss-lib
```

2. Install dependencies:

```bash
go mod tidy
```

## Configuration

The service uses the following default configuration:
- Total Participants: 4 (configurable in demo.go)
- Threshold: 2 (50% of participants)
- Server Port: 8080

To modify the TSS configuration, update the constants in `demo.go`:
```go
const (
    TestParticipants = 4
    TestThreshold    = TestParticipants / 2
)
```

## Running the Service

Start the server:
```bash
cd demo
go run demo.go server.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### 1. Create New Wallet
Creates a new TSS wallet with distributed key shares.

```bash
curl -X POST http://localhost:8080/wallet
```

Response:
```json
{
    "address": "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"
}
```

### 2. List All Wallets
Returns a list of all generated wallet addresses.

```bash
curl http://localhost:8080/wallets
```

Response:
```json
{
    "wallets": [
        "0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
        "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"
    ]
}
```

### 3. Sign Data
Signs data using a specified wallet's distributed key shares.

```bash
curl "http://localhost:8080/sign?data=HelloWorld&wallet=0x71C7656EC7ab88b098defB751B7401B5f6d8976F"
```

Response:
```json
{
    "signature": {
        "signature":"8uSZkBWyNpWl+S3QWFsahFrUhE+bXo3EmNFiBJayXdRE7Gso2q49EZt/PfXEGIhaWhfy1rbwzICayxLpN0v3Pg==",
        "signature_recovery":"AA==",
        "r":"8uSZkBWyNpWl+S3QWFsahFrUhE+bXo3EmNFiBJayXdQ=","s":"ROxrKNquPRGbfz31xBiIWloX8ta28MyAmssS6TdL9z4=","m":"hy5OUM6ZkNiwQTMMR8nd0Rvsa1A66ThqmdqFhOm7EsQ="
    }
}
```