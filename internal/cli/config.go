package cli

const (
	// RequiredArguments is the number of arguments required for normal operation.
	RequiredArguments = 2
)

// Config holds the configuration for the CLI application.
type Config struct {
	// Version is the semantic version of the tool.
	Version string

	// ShowHelp indicates whether the --help flag was provided.
	ShowHelp bool

	// ShowVersion indicates whether the --version flag was provided.
	ShowVersion bool

	// Args contains positional arguments after flag parsing.
	Args []string
}

// NewConfig creates a new Config with default values.
func NewConfig() *Config {
	return &Config{
		Version: "1.0.0",
	}
}

// Validate checks if the configuration is valid for normal operation.
func (c *Config) Validate() error {
	// If showing help or version, no validation needed
	if c.ShowHelp || c.ShowVersion {
		return nil
	}

	// Normal operation requires exactly 2 arguments: date and message
	if len(c.Args) != RequiredArguments {
		return NewMissingArgumentsError(RequiredArguments, len(c.Args))
	}

	return nil
}

// GetDate returns the date argument.
func (c *Config) GetDate() string {
	if len(c.Args) >= 1 {
		return c.Args[0]
	}
	return ""
}

// GetMessage returns the commit message argument.
func (c *Config) GetMessage() string {
	if len(c.Args) >= RequiredArguments {
		return c.Args[1]
	}
	return ""
}
