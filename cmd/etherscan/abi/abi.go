package abi

import (
	"context"
	"errors"
	"fmt"

	"github.com/joshklop/etherscan"
	"github.com/joshklop/etherscan/cmd/etherscan/contract"
	"github.com/peterbourgon/ff/v4"
)

type Config struct {
	Contract *contract.Config
	Command  *ff.Command
	Flags    *ff.FlagSet
}

func New(contract *contract.Config) *Config {
	cfg := &Config{}
	abiFs := ff.NewFlagSet("abi").SetParent(contract.Flags)
	cfg.Command = &ff.Command{
		Name:      "abi",
		Usage:     "etherscan [FLAGS] contract [FLAGS] abi",
		Flags:     abiFs,
		ShortHelp: "retrieve the ABI for a contract",
		Exec: func(ctx context.Context, _ []string) error {
			if contract.Root.ApiKey == "" {
				return errors.New("api key not found")
			}
			client := etherscan.New(contract.Root.ApiKey, etherscan.WithURL(contract.Root.URL))
			abi, err := client.ABI(ctx, contract.Address)
			if err != nil {
				return fmt.Errorf("get ABI for contract %s: %v", contract.Address, err)
			}
			fmt.Fprint(contract.Root.Stdout, string(abi))
			return nil
		},
	}
	contract.Command.Subcommands = append(contract.Command.Subcommands, cfg.Command)
	return cfg
}
