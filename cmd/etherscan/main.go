package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/joshklop/etherscan/cmd/etherscan/root"
	"github.com/joshklop/etherscan/cmd/etherscan/contract"
	"github.com/joshklop/etherscan/cmd/etherscan/abi"
	"github.com/peterbourgon/ff/v4"
	"github.com/peterbourgon/ff/v4/ffhelp"
)

func main() {
	ctx := context.Background()
	stdout := os.Stdout
	stderr := os.Stderr
	if err := Run(ctx, stdout, stderr, os.Args[1:]); err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func Run(ctx context.Context, stdout io.Writer, stderr io.Writer, args []string) error {
	root := root.New(stdout, stderr)
	contract := contract.New(root)
	_ = abi.New(contract)

	if err := root.Command.Parse(args,
		ff.WithEnvVarPrefix("ETHERSCAN"),
		ff.WithConfigFileFlag("config"),
		ff.WithConfigFileParser(ff.PlainParser),
	); err != nil {
		fmt.Fprintf(stderr, "\n%s\n", ffhelp.Command(root.Command))
		if errors.Is(err, ff.ErrHelp) {
			return nil
		}
		return fmt.Errorf("parse: %v", err)
	}

	if err := root.Command.Run(ctx); err != nil {
		if errors.Is(err, ff.ErrNoExec) {
			return nil
		}
		return err
	}

	return nil
}
