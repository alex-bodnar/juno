package wallets

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	extratypes "git.ooo.ua/vipcoin/chain/x/types"
)

type (
	// CreateWalletDB represents a single row inside the "vipcoin_chain_wallets_create_wallet" table
	CreateWalletDB struct {
		Creator        string  `db:"creator"`
		Address        string  `db:"address"`
		AccountAddress string  `db:"account_address"`
		Kind           int32   `db:"kind"`
		State          int32   `db:"state"`
		Extras         ExtraDB `db:"extras"`
	}

	// ExtraDB helpers type
	ExtraDB struct {
		Extras []extratypes.Extra
	}
)

// Value - Make the ExtraDB struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (e ExtraDB) Value() (driver.Value, error) {
	return json.Marshal(e.Extras)
}

// Scan - Make the ExtraDB struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (e *ExtraDB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &e.Extras)
}
