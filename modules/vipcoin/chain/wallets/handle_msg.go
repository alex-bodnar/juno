package wallets

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v2/types"
)

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *types.Tx) error {
	return nil
}
