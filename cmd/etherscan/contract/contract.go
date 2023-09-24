package contract

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/joshklop/etherscan/cmd/etherscan/root"
	"github.com/peterbourgon/ff/v4"
)

const (
	numAddressNibbles = 2 * 20
)

var errTooLong = errors.New("too long")

type Config struct {
	Root    *root.Config
	Address string
	Command *ff.Command
	Flags   *ff.FlagSet
}

func New(root *root.Config) *Config {
	cfg := &Config{
		Root:  root,
		Flags: ff.NewFlagSet("contract").SetParent(root.Flags),
	}
	cfg.Flags.StringVar(&cfg.Address, 0, "address", "", "contract address encoded as hex with the leading 0x.")
	cfg.Command = &ff.Command{
		Name:  "contract",
		Usage: "etherscan [FLAGS] contract [FLAGS] <SUBCOMMAND>",
		Flags: cfg.Flags,
		Exec: func(_ context.Context, _ []string) error {
			address, err := formatAddress(cfg.Address)
			if err != nil {
				return fmt.Errorf("invalid address: %v", err)
			}
			cfg.Address = address
			return nil
		},
	}
	root.Command.Subcommands = append(root.Command.Subcommands, cfg.Command)
	return cfg
}

// formatAddress ensures address is valid and returns the Etherscan-formatted address.
//
// A valid address is:
//   - Not empty.
//   - Shorter than 42 characters if it includes the 0x prefix.
//   - Shorter than 40 characters if it omits the 0x prefix.
//   - Valid hex.
//
// An Etherscan-formatted address is:
//   - Valid.
//   - Zero-padded.
//   - Prefixed by 0x.
func formatAddress(address string) (string, error) {
	if address == "" {
		return "", errors.New("not found")
	}

	if len(address) > 1 {
		if address[:2] == "0x" {
			if len(address) > 2+numAddressNibbles {
				return "", errTooLong
			}
			address = address[2:]
		} else if len(address) > 40 {
			return "", errTooLong
		}
	}
	// hex.DecodeString expects a string of even length.
	if len(address)%2 == 1 {
		address = "0" + address
	}
	if _, err := hex.DecodeString(address); err != nil {
		return "", fmt.Errorf("decode hex: %v", err)
	}

	// Etherscan formatting.
	for len(address) < numAddressNibbles {
		address = "0" + address
	}
	return "0x" + address, nil
}
