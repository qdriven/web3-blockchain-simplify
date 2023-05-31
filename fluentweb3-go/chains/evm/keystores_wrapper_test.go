package evm

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
)

func TestCreateKeyStore(t *testing.T) {
	type args struct {
		walletPath string
		pwd        string
	}
	tests := []struct {
		name string
		args args
		want accounts.Account
	}{
		{name: "TestCreateKeyStore", args: args{walletPath: "./testdata/test_wallet", pwd: "test"},
			want: accounts.Account{}}} // TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateKeyStore(tt.args.walletPath, tt.args.pwd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateKeyStore() = %v, want %v", got, tt.want)
			}
		})
	}
}
