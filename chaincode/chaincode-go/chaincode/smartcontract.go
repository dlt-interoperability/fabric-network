package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// KeyModification describes the kv write that forms part of the key history
type KeyModification struct {
	TxID      string               `json:"txId"`
	Timestamp *timestamp.Timestamp `json:"timestamp"`
	Value     string               `json:"value"`
	IsDelete  bool                 `json:"isDelete"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	key := "key0"
	value := "value0"
	err := ctx.GetStub().PutState(key, []byte(value))
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, key string, value string) error {

	return ctx.GetStub().PutState(key, []byte(value))
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	value, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if value == nil {
		return "", fmt.Errorf("the asset %s does not exist", key)
	}

	return string(value), nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, key string, value string) error {
	exists, err := s.AssetExists(ctx, key)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", key)
	}

	return ctx.GetStub().PutState(key, []byte(value))
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, key string) error {
	exists, err := s.AssetExists(ctx, key)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", key)
	}

	return ctx.GetStub().DelState(key)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, key string) (bool, error) {
	value, err := ctx.GetStub().GetState(key)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return value != nil, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]string, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var results []string
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		result := string(queryResponse.Value)
		results = append(results, result)
	}

	return results, nil
}

// GetHistoryForKey returns the version history of a state for a key.
func (s *SmartContract) GetHistoryForKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return "", err
	}
	defer resultsIterator.Close()

	var results []KeyModification
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return "", err
		}

		record := KeyModification{
			TxID:      response.TxId,
			Timestamp: response.Timestamp,
			IsDelete:  response.IsDelete,
			Value:     string(response.Value),
		}
		results = append(results, record)
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return "", err
	}
	return string(resultsJSON), nil
}
