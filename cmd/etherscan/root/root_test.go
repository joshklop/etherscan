package root_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/joshklop/etherscan/cmd/etherscan/root"
	"github.com/peterbourgon/ff/v4"
	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	testcases := map[string]struct {
		args  []string
		isErr bool
	}{
		"no args": {
			args:  []string{},
			isErr: true, // Require api key.
		},
		"key only": {
			args: []string{"--key", "a"}, // URL not required.
		},
		"url only": {
			args: []string{"--url", "https://example.com"},
			isErr: true, // Require api key.
		},
		"key and invalid url": {
			args: []string{"--key", "a", "--url", "invalid"},
			isErr: true,
		},
		"key and valid url": {
			args: []string{"--key", "a", "--url", "https://example.com"},
		},
	}

	for description, test := range testcases {
		t.Run(description, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}
			rootCmd := root.New(stdout, stderr).Command

			err := rootCmd.Parse(test.args, ff.WithConfigFileFlag("config"))
			require.NoError(t, err)

			if err = rootCmd.Run(context.Background()); test.isErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Zero(t, stdout.Len())
			require.Zero(t, stderr.Len())
		})
	}
}
