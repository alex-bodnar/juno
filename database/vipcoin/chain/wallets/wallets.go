package wallets

import (
	"context"
	"database/sql"

	walletstypes "git.ooo.ua/vipcoin/chain/x/wallets/types"

	"github.com/jmoiron/sqlx"
)

type (
	// Repository - defines a repository for wallets repository
	Repository struct {
		db *sqlx.DB
	}
)

// NewRepository constructor.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateWallet create wallet the given wallets inside the database
func (r Repository) CreateWallet(wallets ...*walletstypes.MsgCreateWallet) error {
	if len(wallets) == 0 {
		return nil
	}

	tx, err := r.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `INSERT INTO vipcoin_chain_wallets_create_wallet 
			("creator", "address", "account_address", "kind", "state", "extras") 
		VALUES 
			(:creator, :address, :account_address, :kind, :state, :extras)`

	for _, wallet := range wallets {
		walletDB, err := toWalletDatabase(wallet)
		if err != nil {
			return err
		}

		if _, err := tx.NamedExec(query, walletDB); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
