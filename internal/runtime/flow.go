package runtime

import (
	"context"
	"encoding/json"
	"os"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access"
	"github.com/onflow/flow-go-sdk/access/http"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
)

type (
	Config struct {
		Accounts struct {
			Contracts map[string]string `json:"contracts"`
			Emulator  EmulatorAccount   `json:"emulator-account"`
		}
	}

	EmulatorAccount struct {
		Address string `json:"address"`
		Key     string `json:"key"`
	}

	Flow struct {
		config     Config
		client     access.Client
		address    flow.Address
		accountKey *flow.AccountKey
		signer     crypto.Signer
	}
)

func NewFlow() (*Flow, error) {
	f, err := os.Open("./flow.json")
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		return nil, err
	}

	client, err := http.NewClient(http.EmulatorHost)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.DecodePrivateKeyHex(crypto.ECDSA_P256, config.Accounts.Emulator.Key)
	if err != nil {
		return nil, err
	}

	address := flow.HexToAddress(config.Accounts.Emulator.Address)
	acc, err := client.GetAccount(context.Background(), address)
	if err != nil {
		return nil, err
	}

	accountKey := acc.Keys[0]
	signer, err := crypto.NewInMemorySigner(privateKey, accountKey.HashAlgo)
	if err != nil {
		return nil, err
	}

	return &Flow{
		config,
		client,
		address,
		accountKey,
		signer,
	}, nil
}

func (f *Flow) CreateAccount(ctx context.Context, publicKeys []*flow.AccountKey, contracts []templates.Contract) (*flow.Account, error) {
	createAccountTx, err := templates.CreateAccount(publicKeys, contracts, f.address)
	if err != nil {
		return nil, err
	}

	block, err := f.client.GetLatestBlock(ctx, true)
	if err != nil {
		return nil, err
	}

	createAccountTx.
		SetProposalKey(f.address, f.accountKey.Index, f.accountKey.SequenceNumber).
		SetReferenceBlockID(block.ID).
		SetPayer(f.address)

	err = createAccountTx.SignEnvelope(f.address, f.accountKey.Index, f.signer)
	if err != nil {
		return nil, err
	}

	err = f.client.SendTransaction(ctx, *createAccountTx)
	if err != nil {
		return nil, err
	}

	result := WaitForSeal(ctx, f.client, createAccountTx.ID())
	return nil
}
