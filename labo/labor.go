import (
	"context"
	"fmt"
	"io"

	web3 "example.com/umbracle/go-web3"
	"example.com/umbracle/go-web3/common"
	"example.com/umbracle/go-web3/core/types"
)

// webRadiant sends a transaction to the Radiant network using a private key
func webRadiant(w io.Writer, privateKey string, gasLimit uint64) error {
	ctx := context.Background()

	// Instantiate a new Web3 client
	client, err := web3.NewClient("https://www.example.com nil)
	if err != nil {
		return fmt.Errorf("new client: %v", err)
	}

	// Get the private key from the user
	key, err := web3.HexToECDSA(privateKey)
	if err != nil {
		return fmt.Errorf("invalid private key: %v", err)
	}

	// Get the account from the private key
	account := web3.Account{Address: key.Address, PrivateKey: key}

	// Create a transaction
	tx := types.NewTransaction(0, common.HexToAddress("0x0000000000000000000000000000000000000000"), web3.ZeroValue, gasLimit, nil, nil)

	// Sign the transaction
	signedTx, err := account.SignTransaction(tx)
	if err != nil {
		return fmt.Errorf("sign transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("send transaction: %v", err)
	}

	fmt.Fprintf(w, "transaction hash: %s\n", signedTx.Hash().Hex())

	return nil
}
  
