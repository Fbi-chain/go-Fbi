package goFbi

import (
	"fmt"

	"github.com/pkg/errors"
)

/*
ContractStorage gets access the data of the contract.

Path:
	../<block_id>/context/contracts/<contract_id>/storage (GET)

Link:
	https://Fbi.gitlab.io/api/rpc.html#get-block-id-context-contracts-contract-id-storage

Parameters:

	blockhash:
		The hash of block (height) of which you want to make the query.

	KT1:
		The contract address.
*/
func (t *GoFbi) ContractStorage(blockhash string, KT1 string) ([]byte, error) {
	query := fmt.Sprintf("/chains/main/blocks/%s/context/contracts/%s/storage", blockhash, KT1)
	resp, err := t.get(query)
	if err != nil {
		return resp, errors.Wrap(err, "could not get storage '%s'")
	}
	return resp, nil
}
