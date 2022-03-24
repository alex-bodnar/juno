package wallets

import (
	extratypes "git.ooo.ua/vipcoin/chain/x/types"
	walletstypes "git.ooo.ua/vipcoin/chain/x/wallets/types"
)

// toWalletDatabase - mapping func to database model
func toWalletDatabase(wallet *walletstypes.MsgCreateWallet) (CreateWalletDB, error) {
	return CreateWalletDB{
		Creator:        wallet.Creator,
		Address:        wallet.Address,
		AccountAddress: wallet.AccountAddress,
		Kind:           int32(wallet.Kind),
		State:          int32(wallet.State),
		Extras:         toExtrasDB(wallet.Extras),
	}, nil
}

// toExtrasDB - mapping func to database model
func toExtrasDB(extras []*extratypes.Extra) ExtraDB {
	result := make([]extratypes.Extra, 0, len(extras))
	for _, extra := range extras {
		result = append(result, *extra)
	}

	return ExtraDB{Extras: result}
}
