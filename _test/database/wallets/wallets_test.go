package wallets

import (
	"testing"

	extratypes "git.ooo.ua/vipcoin/chain/x/types"
	kind "git.ooo.ua/vipcoin/chain/x/wallets/types"
	walletstypes "git.ooo.ua/vipcoin/chain/x/wallets/types"

	walletsdb "github.com/forbole/bdjuno/v2/database/vipcoin/chain/wallets"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestRepository_CreateWallet(t *testing.T) {
	db, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		wallets []*walletstypes.MsgCreateWallet
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				wallets: []*walletstypes.MsgCreateWallet{
					{
						Creator:        "alex",
						Address:        "vcg1qqhekng612225sgmqzujhrzmulk48mhtadxyyn9tk",
						AccountAddress: "vcg1vgpsdg7h3pnw22dcmuzsusscuxmcqqyc2gsles08",
						Kind:           kind.WALLET_KIND_REFERRER_REWARD,
						State:          kind.WALLET_STATE_ACTIVE,
						Extras: []*extratypes.Extra{
							{
								Kind: extratypes.EXTRA_KIND_COMMENT,
								Data: "text",
							},
						},
					},
				},
			},
		}, {
			name: "valid",
			args: args{
				wallets: []*walletstypes.MsgCreateWallet{
					{
						Creator:        "vova",
						Address:        "vcg1qqhekn26p5sg2m22q2ujhrzmulk48mhtadxyyn9tk",
						AccountAddress: "vcg1vgpsdg7h3pnwd22c2uzsusscuxmcqqyc2gsles08",
						Kind:           kind.WALLET_KIND_REFERRER_REWARD,
						State:          kind.WALLET_STATE_ACTIVE,
						Extras: []*extratypes.Extra{
							{
								Kind: extratypes.EXTRA_KIND_COMMENT,
								Data: "text",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := walletsdb.NewRepository(db)

			if err := r.CreateWallet(tt.args.wallets...); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
