package wallets

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/rs/zerolog/log"

	walletstypes "git.ooo.ua/vipcoin/chain/x/wallets/types"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "wallets").Msg("parsing genesis")

	// Unmarshal the bank state
	var walletsState walletstypes.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[walletstypes.ModuleName], &walletsState); err != nil {
		return fmt.Errorf("error while unmarhshaling wallets state: %s", err)
	}

	//return m.db.SaveVipcoinWallets(walletsState.Wallets)
	return nil
}
