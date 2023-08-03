package zepcli

import (
	"fmt"
	"github.com/getzep/zepcli/pkg/jwttools"
	"os"

	"github.com/spf13/cobra"
)

var (
	showVersion bool
	initJWT     bool
)

var rootCmd = &cobra.Command{
	Use:   "zepcli",
	Short: "zepcli is a utility for interacting with the zep service",
	Run:   func(cmd *cobra.Command, args []string) { run() },
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "print version number")
	rootCmd.PersistentFlags().
		BoolVarP(&initJWT, "init-jwt", "i", false, "generate a secret and a new JWT token")
}

// Execute executes the root cobra command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error executing command: %s\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func run() {
	if showVersion {
		fmt.Printf("zep-cli version %s\n", VersionString)
		return
	} else if initJWT {
		err := jwttools.GenerateJWT()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing JWT: %s\n", err)
			os.Exit(1)
		}
		return
	}
}
