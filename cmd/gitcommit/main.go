package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/sgaunet/gitcommit/internal/cli"
)

func main() {
	// Setup structured logging with slog
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Create configuration
	config := cli.NewConfig()

	// Parse command-line flags
	flag.BoolVar(&config.ShowHelp, "help", false, "Show help message")
	flag.BoolVar(&config.ShowHelp, "h", false, "Show help message (shorthand)")
	flag.BoolVar(&config.ShowVersion, "version", false, "Show version information")
	flag.BoolVar(&config.ShowVersion, "v", false, "Show version information (shorthand)")
	flag.Parse()

	// Collect positional arguments
	config.Args = flag.Args()

	// Handle --version flag
	if config.ShowVersion {
		fmt.Printf("gitcommit version %s\n", config.Version)
		os.Exit(cli.ExitSuccess)
	}

	// Handle --help flag
	if config.ShowHelp {
		fmt.Print(cli.HelpText())
		os.Exit(cli.ExitSuccess)
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cli.ExitError)
	}

	// Create and run application
	app := cli.NewApp(config)
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(cli.ExitError)
	}
}
