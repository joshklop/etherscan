package main_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	etherscan "github.com/joshklop/etherscan/cmd/etherscan"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		args    []string
		stdout  string
		stderr  string
		wantErr bool
	}{
		"no args": {
			args: []string{},
			wantErr: true, // Require api key.
		},
		"only api key": {
			args: []string{"--key", "a"},
		},
		"root help": {
			args:   []string{"--help"},
			stderr: rootHelp,
		},
		"contract help": {
			args:   []string{"contract", "--help"},
			stderr: contractHelp,
		},
		"abi help": {
			args:   []string{"contract", "abi", "--help"},
			stderr: abiHelp,
		},
	}

	for description, test := range testcases {
		test := test
		t.Run(description, func(t *testing.T) {
			t.Parallel()

			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			if err := etherscan.Run(context.Background(), stdout, stderr, test.args); test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			{
				want := strings.TrimSpace(test.stdout)
				got := strings.TrimSpace(stdout.String())
				require.Equal(t, want, got)
			}

			{
				want := strings.TrimSpace(test.stderr)
				got := strings.TrimSpace(stderr.String())
				require.Equal(t, want, got)
			}
		})
	}
}

const rootHelp = `
COMMAND
  etherscan

USAGE
  etherscan [FLAGS] <SUBCOMMAND>

SUBCOMMANDS
  contract   

FLAGS
  --key STRING      api key
  --url STRING      the api query url. (default: https://api.etherscan.io/api)
  --config STRING   config file (optional)
`

const contractHelp = `
COMMAND
  contract

USAGE
  etherscan [FLAGS] contract [FLAGS] <SUBCOMMAND>

SUBCOMMANDS
  abi   retrieve the ABI for a contract

FLAGS (contract)
  --address STRING   contract address encoded as hex with the leading 0x.

FLAGS (etherscan)
  --key STRING       api key
  --url STRING       the api query url. (default: https://api.etherscan.io/api)
  --config STRING    config file (optional)
`

const abiHelp = `
COMMAND
  abi -- retrieve the ABI for a contract

USAGE
  etherscan [FLAGS] contract [FLAGS] abi

FLAGS (contract)
  --address STRING   contract address encoded as hex with the leading 0x.

FLAGS (etherscan)
  --key STRING       api key
  --url STRING       the api query url. (default: https://api.etherscan.io/api)
  --config STRING    config file (optional)
`
