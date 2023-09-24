package contract_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/joshklop/etherscan/cmd/etherscan/contract"
	"github.com/joshklop/etherscan/cmd/etherscan/root"
	"github.com/stretchr/testify/require"
)

func TestContract(t *testing.T) {
	// Assume we already have the args required by the root.
	argsPrefix := []string{"--key", "a"}

	testcases := map[string]struct {
		args  []string
		want  string
		isErr bool
	}{
		"requires address": {
			args:  []string{},
			isErr: true,
		},
		"good 0x address": {
			args: []string{"--address", "0x1111111111111111111111111111111111111111"},
			want: "0x1111111111111111111111111111111111111111",
		},
		"good address": {
			args: []string{"--address", "1111111111111111111111111111111111111111"},
			want: "0x1111111111111111111111111111111111111111",
		},
		"good 0x address with leading zero": {
			args: []string{"--address", "0x0111111111111111111111111111111111111111"},
			want: "0x0111111111111111111111111111111111111111",
		},
		"good address with leading zero": {
			args: []string{"--address", "0111111111111111111111111111111111111111"},
			want: "0x0111111111111111111111111111111111111111",
		},
		"good 0x short address": {
			args: []string{"--address", "0x01111111111111111"},
			want: "0x0000000000000000000000001111111111111111",
		},
		"good short address": {
			args: []string{"--address", "01111111111111111"},
			want: "0x0000000000000000000000001111111111111111",
		},
		"long 0x address": {
			args:  []string{"--address", "0x1111111111111001111111111111111111100001111"},
			isErr: true,
		},
		"long address": {
			args:  []string{"--address", "01111111111111110000000011111110000001111111"},
			isErr: true,
		},
		"short 0x address": {
			args: []string{"--address", "0x"},
			want: "0x0000000000000000000000000000000000000000",
		},
		"invalid 0x hex": {
			args:  []string{"--address", "0xx"},
			isErr: true,
		},
		"invalid hex": {
			args:  []string{"--address", "x"},
			isErr: true,
		},
	}

	for description, test := range testcases {
		t.Run(description, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			contractCfg := contract.New(root.New(stdout, stderr))
			args := append(argsPrefix, test.args...)
			require.NoError(t, contractCfg.Command.Parse(args))
			if err := contractCfg.Command.Run(context.Background()); test.isErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.want, contractCfg.Address)
			}

			require.Zero(t, stdout.Len())
			require.Zero(t, stderr.Len())
		})
	}
}
