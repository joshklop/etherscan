package root

import (
	"context"
	"errors"
	"io"
	"net/url"

	"github.com/peterbourgon/ff/v4"
)

type Config struct {
	ApiKey  string
	URL     string
	Command *ff.Command
	Flags   *ff.FlagSet
	Stdout  io.Writer
	Stderr  io.Writer
}

func New(stdout, stderr io.Writer) *Config {
	cfg := &Config{
		Flags:  ff.NewFlagSet("etherscan"),
		Stdout: stdout,
		Stderr: stderr,
	}
	cfg.Flags.StringVar(&cfg.ApiKey, 0, "key", "", "api key")
	cfg.Flags.StringVar(&cfg.URL, 0, "url", "https://api.etherscan.io/api", "the api query url.")
	_ = cfg.Flags.StringLong("config", "", "config file (optional)")
	cfg.Command = &ff.Command{
		Name:  "etherscan",
		Usage: "etherscan [FLAGS] <SUBCOMMAND>",
		Flags: cfg.Flags,
		Exec: func(_ context.Context, _ []string) error {
			// TODO we need better validation for the api key.
			if cfg.ApiKey == "" {
				return errors.New("api key not found")
			}
			if _, err := url.ParseRequestURI(cfg.URL); err != nil {
				return err
			}
			return nil
		},
	}
	return cfg
}
