package wallets

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/bdjuno/v2/database/vipcoin/chain/wallets"
	"github.com/forbole/bdjuno/v2/modules/vipcoin/chain/wallets/source"

	junomessages "github.com/forbole/juno/v2/modules/messages"

	"github.com/forbole/juno/v2/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.MessageModule            = &Module{}
)

// Module represents the x/wallets module
type Module struct {
	cdc        codec.Marshaler
	db         *database.Db
	walletRepo wallets.Repository

	messageParser junomessages.MessageAddressesParser
	keeper        source.Source
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser, keeper source.Source, cdc codec.Marshaler, db *database.Db,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		walletRepo:    *wallets.NewRepository(db.Sqlx),
		messageParser: messageParser,
		keeper:        keeper,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "vipcoin_wallets"
}